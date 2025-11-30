package api

import (
	"quantix-math/pkg/sequences"

	"github.com/gofiber/fiber/v2"
)

type SequenceRequest struct {
	MaxNumber    string `json:"maxNumber"`
	SequenceType string `json:"sequenceType"`
	Positional   bool   `json:"positional"`
}

func GetSequenceHandler(c *fiber.Ctx) error {
	var req SequenceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	if req.MaxNumber == "" || req.SequenceType == "" {
		return c.Status(fiber.StatusBadRequest).SendString("maxNumber and sequenceType are required")
	}

	seq, err := sequences.GetSequence(req.MaxNumber, req.SequenceType, req.Positional)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(seq)
}
