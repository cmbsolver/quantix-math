package sequences

import (
	"fmt"
	"math/big"
)

// A000023: Expansion of e.g.f. exp(-2*x)/(1-x).
// URL: https://oeis.org/A000023
// Description: Expansion of e.g.f. exp(-2*x)/(1-x).
// Recurrence: a(n) = n*a(n-1) + (-2)^n, with a(0) = 1.

// GetA000023Sequence returns the A000023 sequence up to maxNumber terms or the n-th term.
func GetA000023Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000023AtPosition(maxNumber)
	}
	return GenerateA000023Sequence(maxNumber)
}

// GenerateA000023Sequence generates the A000023 sequence up to maxNumber terms.
func GenerateA000023Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n)

	for i := 0; i < n; i++ {
		sequence[i] = CalculateA000023(i)
	}

	result := big.NewInt(0)
	if n > 0 {
		result = sequence[n-1]
	}

	return &NumericSequence{
		Name:     "A000023",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetA000023AtPosition returns the n-th term of A000023.
func GetA000023AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	val := CalculateA000023(int(n.Int64()))

	return &NumericSequence{
		Name:     "A000023",
		Number:   n,
		Sequence: []*big.Int{val},
		Result:   val,
	}, nil
}

// CalculateA000023 calculates the n-th term of A000023.
// Recurrence: a(n) = n*a(n-1) + (-2)^n, with a(0) = 1.
func CalculateA000023(n int) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	if n == 0 {
		return big.NewInt(1)
	}

	// a(n) = n*a(n-1) + (-2)^n
	a := big.NewInt(1)
	for i := 1; i <= n; i++ {
		// next = i * a + (-2)^i
		term1 := new(big.Int).Mul(big.NewInt(int64(i)), a)

		// Correct way to handle (-2)^i:
		twoPowerI := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(i)), nil)
		if i%2 != 0 {
			twoPowerI.Neg(twoPowerI)
		}

		a = new(big.Int).Add(term1, twoPowerI)
	}

	return a
}

// IsA000023 checks if the given number is in the A000023 sequence.
func IsA000023(n *big.Int) (bool, string, error) {
	if n.Sign() < 0 {
		return false, "", nil
	}

	// Since the sequence grows very fast, we can just calculate terms until we pass n.
	for i := 0; ; i++ {
		term := CalculateA000023(i)
		cmp := term.Cmp(n)
		if cmp == 0 {
			return true, fmt.Sprintf("%d", i), nil
		}
		if cmp > 0 && term.Sign() > 0 {
			// Sequence values are: 1, -1, 2, -2, 8, 8, 112, 656, ...
			// For n > 8, they seem to be monotonically increasing.
			if i > 5 {
				break
			}
		}
		// A000023 has some negative values and small values at the beginning.
		// 1, -1, 2, -2, 8, 8, 112, 656, 5504, 49024...
		// Let's just limit the search.
		if i > 100 {
			break
		}
	}

	return false, "", nil
}
