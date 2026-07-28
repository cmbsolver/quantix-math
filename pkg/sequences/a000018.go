package sequences

import (
	"fmt"
	"math/big"
)

// A000018: Number of positive integers <= 2^n of form x^2 + 16*y^2.
// URL: https://oeis.org/A000018
// Description: Number of positive integers <= 2^n of form x^2 + 16*y^2.
// Example: n=2, 2^2=4. Positive integers <= 4 of form x^2 + 16*y^2 are 1 (1^2+16*0^2) and 4 (2^2+16*0^2). So A000018(2) = 2.

// GetA000018Sequence returns the A000018 sequence up to maxNumber terms or the n-th term.
func GetA000018Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000018AtPosition(maxNumber)
	}
	return GenerateA000018Sequence(maxNumber)
}

// GenerateA000018Sequence generates the A000018 sequence up to maxNumber terms.
func GenerateA000018Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n)

	for i := 0; i < n; i++ {
		val := CalculateA000018(i)
		sequence[i] = big.NewInt(int64(val))
	}

	result := big.NewInt(0)
	if n > 0 {
		result = sequence[n-1]
	}

	return &NumericSequence{
		Name:     "Form x^2 + 16y^2 (A000018)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetA000018AtPosition returns the n-th term of A000018.
func GetA000018AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	val := CalculateA000018(int(n.Int64()))

	return &NumericSequence{
		Name:     "Form x^2 + 16y^2 (A000018)",
		Number:   n,
		Sequence: []*big.Int{big.NewInt(int64(val))},
		Result:   big.NewInt(int64(val)),
	}, nil
}

// CalculateA000018 calculates the n-th term of A000018.
// A000018(n) is the number of positive integers <= 2^n of form x^2 + 16*y^2.
func CalculateA000018(n int) int {
	if n < 0 {
		return 0
	}

	limit := int64(1) << uint(n)
	count := 0
	seen := make(map[int64]bool)

	// m = x^2 + 16*y^2
	// x^2 <= limit => x <= sqrt(limit)
	// 16*y^2 <= limit => y^2 <= limit/16 => y <= sqrt(limit/16)

	for y := int64(0); ; y++ {
		y2_16 := 16 * y * y
		if y2_16 > limit {
			if y > 0 {
				break
			}
		}

		for x := int64(0); ; x++ {
			m := x*x + y2_16
			if m > limit {
				break
			}
			if m > 0 && !seen[m] {
				seen[m] = true
				count++
			}
		}

		if y2_16 > limit {
			break
		}
	}

	return count
}
