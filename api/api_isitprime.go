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

// GetIsItPrimeHandler handles the prime check request
// @Summary Check if a number is prime, emirp, circular prime, or semiprime
// @Description Returns the prime status of a given number
// @Tags Math
// @Accept  json
// @Produce  json
// @Param   request  body      IsItPrimeRequest  true  "Prime Check Request"
// @Success 200      {object}  IsItPrimeResponse
// @Failure 400      {string}  string "Invalid request body or number"
// @Router /api/prime [post]
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
