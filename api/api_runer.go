package api

import (
	"quantix-math/pkg/utility/runelib"
	"quantix-math/pkg/utility/runer"

	"github.com/gofiber/fiber/v2"
)

type CalculateLinesRequest struct {
	Input string `json:"input"`
}

// CalculateLinesHandler handles the rune line calculation request
// @Summary Calculate rune lines from input
// @Description Returns the calculated rune lines for the given input string
// @Tags 3301 Tools
// @Accept  json
// @Produce  json
// @Param   request  body      CalculateLinesRequest  true  "Calculate Lines Request"
// @Success 200      {array}   string
// @Failure 400      {string}  string "Invalid request body"
// @Router /api/runer/calculate-lines [post]
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
