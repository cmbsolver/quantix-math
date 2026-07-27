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

// GetDictionaryWordsByParamHandler handles the dictionary words search request
// @Summary Search for dictionary words
// @Description Returns a list of dictionary words based on the provided filters
// @Tags Dictionary
// @Accept  json
// @Produce  json
// @Param   request  body      DictionaryWordsByParamRequest  true  "Search Request"
// @Success 200      {object}  DictionaryWordsByParamResponse
// @Failure 400      {string}  string "Invalid request body or filters"
// @Failure 500      {string}  string "Internal server error"
// @Router /api/dictionary/words [post]
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

// DownloadDictionaryWordsExcelHandler handles the request to download dictionary words as Excel
// @Summary Download dictionary words as Excel
// @Description Generates and returns an Excel file containing dictionary words based on filters
// @Tags Dictionary
// @Produce  application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Param   field  query     []string  true  "Filter fields (e.g., dict_word_length, gem_sum)"
// @Param   param  query     []int     true  "Filter values"
// @Success 200    {file}    binary
// @Failure 400    {string}  string "Invalid or mismatched parameters"
// @Failure 500    {string}  string "Internal server error"
// @Router /api/dictionary/words/download [get]
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

type AnagramRequest struct {
	Word string `json:"word"`
	Type string `json:"type"` // latin, runeglish, rune
}

type AnagramResponse struct {
	Anagrams []string `json:"anagrams"`
	Count    int      `json:"count"`
}

// GetAnagramsHandler handles the anagram search request
// @Summary Find anagrams for a word
// @Description Returns a list of anagrams for the given word and type
// @Tags Dictionary
// @Accept  json
// @Produce  json
// @Param   request  body      AnagramRequest  true  "Anagram Request"
// @Success 200      {object}  AnagramResponse
// @Failure 400      {string}  string "Invalid request body or word"
// @Failure 500      {string}  string "Internal server error"
// @Router /api/dictionary/anagrams [post]
func GetAnagramsHandler(c *fiber.Ctx) error {
	var req AnagramRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	if req.Word == "" {
		return c.Status(fiber.StatusBadRequest).SendString("word is required")
	}

	conn, err := db.InitConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to connect to database")
	}
	defer func() { _ = db.CloseConnection(conn) }()

	anagramWords := tables.GetAnagrams(conn, req.Word, req.Type)
	var anagrams []string
	for _, w := range anagramWords {
		switch req.Type {
		case "runeglish":
			anagrams = append(anagrams, w.RuneglishWordText)
		case "rune":
			anagrams = append(anagrams, w.RuneWordText)
		default:
			anagrams = append(anagrams, w.DictionaryWordText)
		}
	}

	return c.JSON(AnagramResponse{
		Anagrams: anagrams,
		Count:    len(anagrams),
	})
}
