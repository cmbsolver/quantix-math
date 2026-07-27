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

type MobiusErrorResponse struct {
	Error string `json:"error"`
}

type MobiusDirectResponse struct {
	N      int    `json:"n"`
	Mu     int8   `json:"mu"`
	Reason string `json:"reason"`
}

// MobiusDirectHandler handles the direct Mobius function calculation request
// @Summary Calculate Mobius function for numbers
// @Description Returns the Mobius function value (mu) for a list of numbers
// @Tags Math
// @Accept  json
// @Produce  json
// @Param   request  body      MobiusDirectRequest  true  "Mobius Direct Request"
// @Success 200      {array}   MobiusDirectResponse
// @Failure 400      {object}  MobiusErrorResponse
// @Router /api/mobius/direct [post]
func MobiusDirectHandler(c *fiber.Ctx) error {
	var req MobiusDirectRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(MobiusErrorResponse{Error: "Invalid request"})
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

	results := make([]MobiusDirectResponse, len(nums))
	for i, n := range nums {
		mu := mobiusCalc.GetMu(n)
		reason := ""
		if mu == 0 && n > 0 {
			reason = "non-square-free"
		}
		results[i] = MobiusDirectResponse{
			N:      n,
			Mu:     mu,
			Reason: reason,
		}
	}

	return c.JSON(results)
}

// MobiusMaskHandler handles the sequential index masking request
// @Summary Apply Mobius masking to a data sequence
// @Description Applies the Mobius sequential index masking to the provided data
// @Tags Math
// @Accept  json
// @Produce  json
// @Param   request  body      MobiusMaskRequest  true  "Mobius Mask Request"
// @Success 200      {array}   float64
// @Failure 400      {object}  MobiusErrorResponse
// @Router /api/mobius/mask [post]
func MobiusMaskHandler(c *fiber.Ctx) error {
	var req MobiusMaskRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(MobiusErrorResponse{Error: "Invalid request"})
	}

	output := mobiusCalc.SequentialIndexMasking(req.Data)
	return c.JSON(output)
}

type InversionResult struct {
	D    int     `json:"d"`
	MuNd int8    `json:"mu_n_d"`
	GD   float64 `json:"g_d"`
	Term float64 `json:"term"`
}

type MobiusDivisorResponse struct {
	N                int               `json:"n"`
	SumMuD           int               `json:"sum_mu_d"`
	Divisors         []int             `json:"divisors"`
	InversionDetails []InversionResult `json:"inversion_details"`
	TotalInversion   float64           `json:"total_inversion"`
}

// MobiusDivisorHandler handles the Mobius divisor summation and inversion request
// @Summary Calculate Mobius divisor sum and inversion
// @Description Returns the divisor sum and Mobius inversion details for a number N
// @Tags Math
// @Accept  json
// @Produce  json
// @Param   request  body      MobiusDivisorRequest  true  "Mobius Divisor Request"
// @Success 200      {object}  MobiusDivisorResponse
// @Failure 400      {object}  MobiusErrorResponse
// @Router /api/mobius/divisor [post]
func MobiusDivisorHandler(c *fiber.Ctx) error {
	var req MobiusDivisorRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(MobiusErrorResponse{Error: "Invalid request"})
	}

	if req.N < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(MobiusErrorResponse{Error: "n must be >= 1"})
	}

	sumMu := mobiusCalc.DivisorSummation(req.N)
	divisors := mobiusCalc.GetDivisors(req.N)

	// For inversion step, we need a g(d). Let's use g(d) = d as a default if not provided,
	// or maybe the user wants to provide an array.
	// For now, let's just return the inversion with g(d) = d as an example,
	// or allow the user to provide an array in the request.

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

	return c.JSON(MobiusDivisorResponse{
		N:                req.N,
		SumMuD:           sumMu,
		Divisors:         divisors,
		InversionDetails: inversionDetails,
		TotalInversion:   totalInversion,
	})
}
