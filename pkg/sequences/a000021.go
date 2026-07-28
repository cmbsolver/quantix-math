package sequences

import (
	"fmt"
	"math/big"
)

// A000021: Number of positive integers <= 2^n of form x^2 + 12*y^2.
// URL: https://oeis.org/A000021
// Description: Number of positive integers <= 2^n of form x^2 + 12*y^2.
// Formerly M0357 N0134

// GetA000021Sequence returns the A000021 sequence up to maxNumber terms or the n-th term.
func GetA000021Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000021AtPosition(maxNumber)
	}
	return GenerateA000021Sequence(maxNumber)
}

// GenerateA000021Sequence generates the A000021 sequence up to maxNumber terms.
func GenerateA000021Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n)

	for i := 0; i < n; i++ {
		val := CalculateA000021(i)
		sequence[i] = big.NewInt(int64(val))
	}

	result := big.NewInt(0)
	if n > 0 {
		result = sequence[n-1]
	}

	return &NumericSequence{
		Name:     "Form x^2 + 12y^2 (A000021)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetA000021AtPosition returns the n-th term of A000021.
func GetA000021AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	val := CalculateA000021(int(n.Int64()))

	return &NumericSequence{
		Name:     "Form x^2 + 12y^2 (A000021)",
		Number:   n,
		Sequence: []*big.Int{big.NewInt(int64(val))},
		Result:   big.NewInt(int64(val)),
	}, nil
}

// CalculateA000021 calculates the n-th term of A000021.
// A000021(n) is the number of positive integers <= 2^n of form x^2 + 12*y^2.
func CalculateA000021(n int) int {
	if n < 0 {
		return 0
	}

	limit := int64(1) << uint(n)
	count := 0
	seen := make(map[int64]bool)

	// m = x^2 + 12*y^2
	// x^2 <= limit => x <= sqrt(limit)
	// 12*y^2 <= limit => y^2 <= limit/12 => y <= sqrt(limit/12)

	for y := int64(0); ; y++ {
		y2_12 := 12 * y * y
		if y2_12 > limit {
			if y > 0 {
				break
			}
		}

		for x := int64(0); ; x++ {
			m := x*x + y2_12
			if m > limit {
				break
			}
			if m > 0 && !seen[m] {
				seen[m] = true
				count++
			}
		}

		if y2_12 > limit {
			break
		}
	}

	return count
}
