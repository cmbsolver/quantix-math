package sequences

import (
	"fmt"
	"math/big"
)

// Number of integers <= 2^n of form x^2 - 2y^2 (OEIS A000047).
// URL: https://oeis.org/A000047

// GetA000047Sequence returns A000047 either up to n or at position n.
func GetA000047Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000047AtPosition(maxNumber)
	}

	return GenerateA000047Sequence(maxNumber)
}

// GenerateA000047Sequence generates A000047 from a(0) through a(maxNumber).
func GenerateA000047Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := maxNumber.Int64()
	sequence := generateA000047Terms(limit)

	return &NumericSequence{
		Name:     "Number of integers <= 2^n of form x^2 - 2y^2 (A000047)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[limit],
	}, nil
}

// GetA000047AtPosition returns the n-th term of A000047.
func GetA000047AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	sequence := generateA000047Terms(n.Int64())
	result := sequence[n.Int64()]

	return &NumericSequence{
		Name:     "Number of integers <= 2^n of form x^2 - 2y^2 (A000047)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// generateA000047Terms computes terms a(0) through a(limit) for OEIS A000047.
func generateA000047Terms(limit int64) []*big.Int {
	if limit >= 62 {
		return nil
	}

	upperBound := int64(1) << uint(limit)
	spf := make([]int64, upperBound+1)
	for i := int64(2); i <= upperBound; i++ {
		if spf[i] == 0 {
			spf[i] = i
			if i <= upperBound/i {
				for j := i * i; j <= upperBound; j += i {
					if spf[j] == 0 {
						spf[j] = i
					}
				}
			}
		}
	}

	sequence := make([]*big.Int, limit+1)
	count := int64(0)
	nextPowerIndex := int64(0)
	nextPowerValue := int64(1)

	for m := int64(1); m <= upperBound; m++ {
		if isA000047Representable(m, spf) {
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

// isA000047Representable checks representability criterion using prime exponents modulo 8.
func isA000047Representable(m int64, spf []int64) bool {
	n := m
	for n > 1 {
		p := spf[n]
		exponent := int64(0)
		for n%p == 0 {
			n /= p
			exponent++
		}
		if (p%8 == 3 || p%8 == 5) && exponent%2 == 1 {
			return false
		}
	}
	return true
}
