package sequences

import (
	"fmt"
	"math/big"
)

// Lucas numbers beginning at 2: L(n) = L(n-1) + L(n-2), L(0) = 2, L(1) = 1 (OEIS A000032).
// URL: https://oeis.org/A000032

// GetA000032Sequence returns the A000032 sequence.
func GetA000032Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000032AtPosition(maxNumber)
	}

	return GenerateA000032Sequence(maxNumber)
}

// GenerateA000032Sequence generates A000032 from a(0) through a(maxNumber).
func GenerateA000032Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := int(maxNumber.Int64())
	sequence := make([]*big.Int, limit+1)

	sequence[0] = big.NewInt(2)
	if limit >= 1 {
		sequence[1] = big.NewInt(1)
	}

	for i := 2; i <= limit; i++ {
		sequence[i] = new(big.Int).Add(sequence[i-1], sequence[i-2])
	}

	return &NumericSequence{
		Name:     "Lucas numbers (A000032)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[limit],
	}, nil
}

// GetA000032AtPosition returns the n-th term of A000032.
func GetA000032AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	result := calculateA000032(n.Int64())

	return &NumericSequence{
		Name:     "Lucas numbers (A000032)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateA000032 computes a(n) for OEIS A000032.
func calculateA000032(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(2)
	}
	if n == 1 {
		return big.NewInt(1)
	}

	prev2 := big.NewInt(2)
	prev1 := big.NewInt(1)

	for i := int64(2); i <= n; i++ {
		current := new(big.Int).Add(prev1, prev2)
		prev2 = prev1
		prev1 = current
	}

	return prev1
}
