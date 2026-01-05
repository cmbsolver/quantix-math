package api

import (
	"fmt"
	"io"

	"quantix-math/pkg/db"
	"quantix-math/pkg/db/tables"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

type DictionaryWordsByParamRequest struct {
	Field string `json:"field"`
	Param int    `json:"param"`
}

type DictionaryWordsByParamResponse struct {
	Field string                  `json:"field"`
	Param int                     `json:"param"`
	Words []tables.DictionaryWord `json:"words"`
	Count int                     `json:"count"`
}

func GetDictionaryWordsByParamHandler(c *fiber.Ctx) error {
	var req DictionaryWordsByParamRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	if req.Param <= 0 {
		return c.Status(fiber.StatusBadRequest).SendString("param must be a positive integer")
	}

	// IMPORTANT: whitelist DB columns to avoid SQL injection via `field`
	allowedFields := map[string]string{
		"dict_rune_length":      "dict_rune_length",
		"dict_runeglish_length": "dict_runeglish_length",
		"dict_word_length":      "dict_word_length",
		"gem_sum":               "gem_sum",
	}

	dbField, ok := allowedFields[req.Field]
	if !ok {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf(
			"invalid field; allowed: %v",
			[]string{"dict_rune_length", "dict_runeglish_length", "dict_word_length", "gem_sum"},
		))
	}

	conn, err := db.InitConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to connect to database")
	}
	defer func() { _ = db.CloseConnection(conn) }()

	words := tables.GetDictionaryWordsByParam(conn, dbField, req.Param)

	return c.JSON(DictionaryWordsByParamResponse{
		Field: req.Field,
		Param: req.Param,
		Words: words,
		Count: len(words),
	})
}

func DownloadDictionaryWordsExcelHandler(c *fiber.Ctx) error {
	field := c.Query("field")
	paramStr := c.Query("param")
	var param int
	fmt.Sscanf(paramStr, "%d", &param)

	if param <= 0 || field == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid parameters")
	}

	allowedFields := map[string]string{
		"dict_rune_length":      "dict_rune_length",
		"dict_runeglish_length": "dict_runeglish_length",
		"dict_word_length":      "dict_word_length",
		"gem_sum":               "gem_sum",
	}

	dbField, ok := allowedFields[field]
	if !ok {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid field")
	}

	conn, err := db.InitConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Database connection failed")
	}
	defer func() { _ = db.CloseConnection(conn) }()

	words := tables.GetDictionaryWordsByParam(conn, dbField, param)

	f := excelize.NewFile()
	defer func() { _ = f.Close() }()

	sheet := "Words"
	index, _ := f.NewSheet(sheet)
	f.DeleteSheet("Sheet1")

	// Set Headers
	headers := []string{"Word", "Runeglish", "Rune", "Gem Sum", "Word Len", "Runeglish Len", "Rune Len", "Language"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// Set Data
	for i, w := range words {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), w.DictionaryWordText)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), w.RuneglishWordText)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), w.RuneWordText)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), w.GemSum)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), w.DictionaryWordLength)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), w.RuneglishWordLength)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), w.RuneWordLength)
		f.SetCellValue(sheet, fmt.Sprintf("H%d", row), w.Language)
	}

	f.SetActiveSheet(index)

	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=dictionary_words_%s_%d.xlsx", field, param))

	// Stream file to response
	reader, writer := io.Pipe()
	go func() {
		defer writer.Close()
		if err := f.Write(writer); err != nil {
			fmt.Printf("Excel write error: %v\n", err)
		}
	}()

	return c.SendStream(reader)
}
