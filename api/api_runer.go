package api

import (
	"quantix-math/pkg/utility/runelib"
	"quantix-math/pkg/utility/runer"

	"github.com/gofiber/fiber/v2"
)

type CalculateLinesRequest struct {
	Input string `json:"input"`
}

func CalculateLinesHandler(c *fiber.Ctx) error {
	var req CalculateLinesRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	charRepo := runelib.NewCharacterRepo()
	lc := runer.NewLineCalculator(charRepo)
	results := lc.CalculateLines(req.Input)

	return c.JSON(results)
}
