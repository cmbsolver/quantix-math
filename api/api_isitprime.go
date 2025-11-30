package api

import (
	"fmt"
	"math/big"
	"quantix-math/pkg/sequences"

	"github.com/gofiber/fiber/v2"
)

type IsItPrimeRequest struct {
	Number string `json:"number"`
}

type IsItPrimeResponse struct {
	Number          string `json:"number"`
	IsPrime         bool   `json:"isPrime"`
	IsEmirp         bool   `json:"isEmirp"`
	IsCircularPrime bool   `json:"isCircularPrime"`
	IsSemiprime     bool   `json:"isSemiprime"`
}

func GetIsItPrimeHandler(c *fiber.Ctx) error {
	var req IsItPrimeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	if req.Number == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Number is required")
	}

	number := new(big.Int)
	number, ok := number.SetString(req.Number, 10)
	if !ok {
		fmt.Printf("Invalid max number: %s\n", req.Number)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid number")
	}

	isPrime := sequences.IsPrime(number)
	isEmirp := sequences.IsEmirp(number)
	isCircularPrime := sequences.IsCircularPrime(number)
	isSemiPrime := sequences.IsSemiPrime(number)

	response := &IsItPrimeResponse{
		Number:          req.Number,
		IsPrime:         isPrime,
		IsEmirp:         isEmirp,
		IsCircularPrime: isCircularPrime,
		IsSemiprime:     isSemiPrime,
	}

	return c.JSON(response)
}
