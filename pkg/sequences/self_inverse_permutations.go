package sequences

import (
	"fmt"
	"math/big"
)

// Self-inverse permutations (involutions) (OEIS A000085)
// URL: https://oeis.org/A000085
// a(n) is the number of self-inverse permutations on n letters (involutions).
// Also the number of standard Young tableaux with n cells.
// Recurrence: a(n) = a(n-1) + (n-1)*a(n-2) for n > 1; a(0) = 1, a(1) = 1.

// GetSelfInversePermutationsSequence returns the self-inverse permutations sequence (OEIS A000085).
func GetSelfInversePermutationsSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetSelfInversePermutationsAtPosition(maxNumber)
	}
	return GenerateSelfInversePermutationsSequence(maxNumber)
}

// GenerateSelfInversePermutationsSequence generates the A000085 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateSelfInversePermutationsSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number must be at least 0")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		sequence[i] = calculateA000085(int64(i))
	}

	return &NumericSequence{
		Name:     "Self-inverse permutations (A000085)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetSelfInversePermutationsAtPosition returns the n-th term of A000085 (n >= 0).
func GetSelfInversePermutationsAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	result := calculateA000085(n.Int64())

	return &NumericSequence{
		Name:     "Self-inverse permutations (A000085)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateA000085 calculates the n-th term of OEIS A000085 using recurrence.
// a(n) = a(n-1) + (n-1)*a(n-2) for n > 1; a(0) = 1, a(1) = 1.
func calculateA000085(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	if n == 1 {
		return big.NewInt(1)
	}

	a0 := big.NewInt(1)
	a1 := big.NewInt(1)
	an := new(big.Int)

	for i := int64(2); i <= n; i++ {
		// an = a1 + (i-1)*a0
		term2 := new(big.Int).Mul(big.NewInt(i-1), a0)
		an = new(big.Int).Add(a1, term2)

		a0.Set(a1)
		a1.Set(an)
	}

	return a1
}
