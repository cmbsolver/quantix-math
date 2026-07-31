package api

import (
	"quantix-math/pkg/cryptanalysis"

	"github.com/gofiber/fiber/v2"
)

type CipherAnalyzeRequest struct {
	Text     string `json:"text"`
	Alphabet string `json:"alphabet"`
}

type MatsuiAnalyzeRequest struct {
	Matsui cryptanalysis.MatsuiInput `json:"matsui"`
}

type CipherAnalyzeError struct {
	Error string `json:"error"`
}

// CipherAnalyzeHandler handles cipher-statistics and ranking requests.
func CipherAnalyzeHandler(c *fiber.Ctx) error {
	var req CipherAnalyzeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(CipherAnalyzeError{Error: "Invalid request body"})
	}

	if req.Text == "" {
		return c.Status(fiber.StatusBadRequest).JSON(CipherAnalyzeError{Error: "text is required"})
	}

	result, err := cryptanalysis.Analyze(req.Text, req.Alphabet)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(CipherAnalyzeError{Error: err.Error()})
	}

	return c.JSON(result)
}

// MatsuiAnalyzeHandler handles Matsui linear cryptanalysis requests.
func MatsuiAnalyzeHandler(c *fiber.Ctx) error {
	var req MatsuiAnalyzeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(CipherAnalyzeError{Error: "Invalid request body"})
	}

	result := cryptanalysis.AnalyzeMatsui(req.Matsui)
	return c.JSON(result)
}
