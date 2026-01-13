package api

import (
	"fmt"
	"io"

	"quantix-math/pkg/db"
	"quantix-math/pkg/db/tables"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

type DictionarySearchRequest struct {
	Filters []struct {
		Field string `json:"field"`
		Param int    `json:"param"`
	} `json:"filters"`
}

type DictionaryWordsByParamRequest struct {
	Filters []struct {
		Field string `json:"field"`
		Param int    `json:"param"`
	} `json:"filters"`
}

type DictionaryWordsByParamResponse struct {
	Words []tables.DictionaryWord `json:"words"`
	Count int                     `json:"count"`
}

func GetDictionaryWordsByParamHandler(c *fiber.Ctx) error {
	var req DictionaryWordsByParamRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	if len(req.Filters) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("at least one filter is required")
	}

	// IMPORTANT: whitelist DB columns to avoid SQL injection via `field`
	allowedFields := map[string]bool{
		"dict_rune_length":      true,
		"dict_runeglish_length": true,
		"dict_word_length":      true,
		"gem_sum":               true,
	}

	filterMap := make(map[string]int)
	for _, f := range req.Filters {
		if !allowedFields[f.Field] {
			return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("invalid field: %s", f.Field))
		}
		if f.Param <= 0 {
			return c.Status(fiber.StatusBadRequest).SendString("params must be positive integers")
		}
		filterMap[f.Field] = f.Param
	}

	conn, err := db.InitConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to connect to database")
	}
	defer func() { _ = db.CloseConnection(conn) }()

	words := tables.GetDictionaryWordsByFilters(conn, filterMap)

	return c.JSON(DictionaryWordsByParamResponse{
		Words: words,
		Count: len(words),
	})
}

func DownloadDictionaryWordsExcelHandler(c *fiber.Ctx) error {
	queryArgs := c.Context().QueryArgs()
	fieldArgs := queryArgs.PeekMulti("field")
	paramArgs := queryArgs.PeekMulti("param")

	if len(fieldArgs) == 0 || len(fieldArgs) != len(paramArgs) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid or mismatched parameters")
	}

	allowedFields := map[string]bool{
		"dict_rune_length":      true,
		"dict_runeglish_length": true,
		"dict_word_length":      true,
		"gem_sum":               true,
	}

	filterMap := make(map[string]int)
	for i := range fieldArgs {
		field := string(fieldArgs[i])
		if !allowedFields[field] {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid field: " + field)
		}
		var p int
		if _, err := fmt.Sscanf(string(paramArgs[i]), "%d", &p); err != nil || p <= 0 {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid param value")
		}
		filterMap[field] = p
	}

	conn, err := db.InitConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Database connection failed")
	}
	defer func() { _ = db.CloseConnection(conn) }()

	words := tables.GetDictionaryWordsByFilters(conn, filterMap)

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
	c.Set("Content-Disposition", "attachment; filename=dictionary_words_search.xlsx")

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
