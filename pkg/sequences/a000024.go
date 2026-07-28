package sequences

import (
	"fmt"
	"math/big"
)

// A000024: Number of positive integers <= 2^n of form x^2 + 10*y^2.
// URL: https://oeis.org/A000024
// Description: Number of positive integers <= 2^n of form x^2 + 10*y^2.
// Formerly M0368 N0139

// GetA000024Sequence returns the A000024 sequence up to maxNumber terms or the n-th term.
func GetA000024Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000024AtPosition(maxNumber)
	}
	return GenerateA000024Sequence(maxNumber)
}

// GenerateA000024Sequence generates the A000024 sequence up to maxNumber terms.
func GenerateA000024Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n)

	for i := 0; i < n; i++ {
		val := CalculateA000024(i)
		sequence[i] = new(big.Int).SetUint64(val)
	}

	var result *big.Int
	if n > 0 {
		result = sequence[n-1]
	} else {
		result = big.NewInt(0)
	}

	return &NumericSequence{
		Name:     "Form x^2 + 10y^2 (A000024)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetA000024AtPosition returns the n-th term of A000024.
func GetA000024AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	val := CalculateA000024(int(n.Int64()))

	return &NumericSequence{
		Name:     "Form x^2 + 10y^2 (A000024)",
		Number:   n,
		Sequence: []*big.Int{new(big.Int).SetUint64(val)},
		Result:   new(big.Int).SetUint64(val),
	}, nil
}

// CalculateA000024 calculates the n-th term of A000024.
// A000024(n) is the number of positive integers <= 2^n of form x^2 + 10*y^2.
func CalculateA000024(n int) uint64 {
	if n < 0 {
		return 0
	}

	limit := int64(1) << uint(n)
	count := uint64(0)
	seen := make(map[int64]bool)

	// m = x^2 + 10*y^2
	// x^2 <= limit => x <= sqrt(limit)
	// 10*y^2 <= limit => y^2 <= limit/10 => y <= sqrt(limit/10)

	for y := int64(0); ; y++ {
		y2_10 := 10 * y * y
		if y2_10 > limit {
			break
		}

		for x := int64(0); ; x++ {
			m := x*x + y2_10
			if m > limit {
				break
			}
			if m > 0 && !seen[m] {
				seen[m] = true
				count++
			}
		}
	}

	return count
}
