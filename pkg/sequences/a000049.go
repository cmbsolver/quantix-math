package sequences

import (
	"fmt"
	"math/big"
)

// Number of positive integers <= 2^n of the form 3*x^2 + 4*y^2 (OEIS A000049).
// URL: https://oeis.org/A000049

// GetA000049Sequence returns A000049 either up to n or at position n.
func GetA000049Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000049AtPosition(maxNumber)
	}

	return GenerateA000049Sequence(maxNumber)
}

// GenerateA000049Sequence generates A000049 from a(0) through a(maxNumber).
func GenerateA000049Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := maxNumber.Int64()
	if limit > 25 {
		return nil, fmt.Errorf("max number too large for direct computation: %d", limit)
	}

	sequence := generateA000049Terms(limit)

	return &NumericSequence{
		Name:     "Number of positive integers <= 2^n of the form 3*x^2 + 4*y^2 (A000049)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[limit],
	}, nil
}

// GetA000049AtPosition returns the n-th term of A000049.
func GetA000049AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	limit := n.Int64()
	if limit > 25 {
		return nil, fmt.Errorf("position too large for direct computation: %d", limit)
	}

	sequence := generateA000049Terms(limit)
	result := sequence[limit]

	return &NumericSequence{
		Name:     "Number of positive integers <= 2^n of the form 3*x^2 + 4*y^2 (A000049)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// generateA000049Terms computes terms a(0) through a(limit) for OEIS A000049.
func generateA000049Terms(limit int64) []*big.Int {
	upperBound := int64(1) << uint(limit)
	representable := make([]bool, upperBound+1)

	maxX := int64Sqrt(upperBound / 3)
	maxY := int64Sqrt(upperBound / 4)

	for x := int64(0); x <= maxX; x++ {
		xPart := 3 * x * x
		for y := int64(0); y <= maxY; y++ {
			value := xPart + 4*y*y
			if value > upperBound {
				break
			}
			if value > 0 {
				representable[value] = true
			}
		}
	}

	sequence := make([]*big.Int, limit+1)
	count := int64(0)
	nextPowerIndex := int64(0)
	nextPowerValue := int64(1)

	for m := int64(1); m <= upperBound; m++ {
		if representable[m] {
			count++
		}
		if m == nextPowerValue {
			sequence[nextPowerIndex] = big.NewInt(count)
			nextPowerIndex++
			if nextPowerIndex <= limit {
				nextPowerValue <<= 1
			}
		}
	}

	return sequence
}

// int64Sqrt returns floor(sqrt(n)) for n >= 0.
func int64Sqrt(n int64) int64 {
	if n <= 0 {
		return 0
	}
	x := n
	y := (x + 1) / 2
	for y < x {
		x = y
		y = (x + n/x) / 2
	}
	return x
}
