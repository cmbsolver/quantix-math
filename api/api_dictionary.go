package api

import (
	"fmt"

	"quantix-math/pkg/db"
	"quantix-math/pkg/db/tables"

	"github.com/gofiber/fiber/v2"
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
