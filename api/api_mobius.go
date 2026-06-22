package api

import (
	"quantix-math/pkg/sequences"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var mobiusCalc = sequences.NewMobiusCalculator(1000000)

type MobiusDirectRequest struct {
	Numbers string `json:"numbers"` // Comma separated or single number
}

type MobiusMaskRequest struct {
	Data []float64 `json:"data"`
}

type MobiusDivisorRequest struct {
	N int `json:"n"`
}

func MobiusDirectHandler(c *fiber.Ctx) error {
	var req MobiusDirectRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	parts := strings.Split(req.Numbers, ",")
	var nums []int
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		n, err := strconv.Atoi(p)
		if err == nil {
			nums = append(nums, n)
		}
	}

	results := make([]fiber.Map, len(nums))
	for i, n := range nums {
		mu := mobiusCalc.GetMu(n)
		reason := ""
		if mu == 0 && n > 0 {
			reason = "non-square-free"
		}
		results[i] = fiber.Map{
			"n":      n,
			"mu":     mu,
			"reason": reason,
		}
	}

	return c.JSON(results)
}

func MobiusMaskHandler(c *fiber.Ctx) error {
	var req MobiusMaskRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	output := mobiusCalc.SequentialIndexMasking(req.Data)
	return c.JSON(output)
}

func MobiusDivisorHandler(c *fiber.Ctx) error {
	var req MobiusDivisorRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if req.N < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "n must be >= 1"})
	}

	sumMu := mobiusCalc.DivisorSummation(req.N)
	divisors := mobiusCalc.GetDivisors(req.N)

	// For inversion step, we need a g(d). Let's use g(d) = d as a default if not provided,
	// or maybe the user wants to provide an array.
	// For now, let's just return the inversion with g(d) = d as an example,
	// or allow the user to provide an array in the request.

	type InversionResult struct {
		D    int     `json:"d"`
		MuNd int8    `json:"mu_n_d"`
		GD   float64 `json:"g_d"`
		Term float64 `json:"term"`
	}

	var inversionDetails []InversionResult
	totalInversion := 0.0
	for _, d := range divisors {
		muNd := mobiusCalc.GetMu(req.N / d)
		gd := float64(d) // Default g(d) = d
		term := float64(muNd) * gd
		inversionDetails = append(inversionDetails, InversionResult{
			D:    d,
			MuNd: muNd,
			GD:   gd,
			Term: term,
		})
		totalInversion += term
	}

	return c.JSON(fiber.Map{
		"n":                 req.N,
		"sum_mu_d":          sumMu,
		"divisors":          divisors,
		"inversion_details": inversionDetails,
		"total_inversion":   totalInversion,
	})
}
