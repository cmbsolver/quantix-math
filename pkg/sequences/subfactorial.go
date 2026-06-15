package sequences

import (
	"fmt"
	"math/big"
)

// Subfactorial or derangements (OEIS A000166)
// URL: https://oeis.org/A000166
// a(n) is the number of permutations of n elements with no fixed points.
// Formula: a(n) = (n-1)*(a(n-1) + a(n-2)) with a(0)=1, a(1)=0.

// GetSubfactorialSequence returns the subfactorial sequence up to maxNumber (OEIS A000166).
func GetSubfactorialSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetSubfactorialAtPosition(maxNumber)
	}
	return GenerateSubfactorialSequence(maxNumber)
}

// GenerateSubfactorialSequence generates the A000166 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateSubfactorialSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number must be at least 0")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		sequence[i] = calculateSubfactorial(int64(i))
	}

	return &NumericSequence{
		Name:     "Subfactorial (A000166)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetSubfactorialAtPosition returns the n-th term of A000166 (n >= 0).
func GetSubfactorialAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	result := calculateSubfactorial(n.Int64())

	return &NumericSequence{
		Name:     "Subfactorial (A000166)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateSubfactorial calculates the n-th subfactorial number.
// Uses the recursive formula: a(n) = (n-1)*(a(n-1) + a(n-2))
func calculateSubfactorial(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	if n == 1 {
		return big.NewInt(0)
	}

	a0 := big.NewInt(1)
	a1 := big.NewInt(0)
	res := big.NewInt(0)

	for i := int64(2); i <= n; i++ {
		// res = (i-1) * (a1 + a0)
		res = new(big.Int).Add(a1, a0)
		res.Mul(res, big.NewInt(i-1))

		a0.Set(a1)
		a1.Set(res)
	}

	return a1
}
