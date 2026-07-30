package sequences

import (
	"fmt"
	"math/big"
)

// Fibonacci numbers: F(n) = F(n-1) + F(n-2) with F(0) = 0 and F(1) = 1 (OEIS A000045).
// URL: https://oeis.org/A000045

// GetA000045Sequence returns the A000045 sequence.
func GetA000045Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000045AtPosition(maxNumber)
	}

	return GenerateA000045Sequence(maxNumber)
}

// GenerateA000045Sequence generates A000045 from a(0) through a(maxNumber).
func GenerateA000045Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := maxNumber.Int64()
	sequence := generateA000045Terms(limit)

	return &NumericSequence{
		Name:     "Fibonacci numbers (A000045)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[limit],
	}, nil
}

// GetA000045AtPosition returns the n-th term of A000045.
func GetA000045AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	sequence := generateA000045Terms(n.Int64())
	result := sequence[n.Int64()]

	return &NumericSequence{
		Name:     "Fibonacci numbers (A000045)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// generateA000045Terms computes terms a(0) through a(limit) for OEIS A000045.
func generateA000045Terms(limit int64) []*big.Int {
	sequence := make([]*big.Int, limit+1)
	sequence[0] = big.NewInt(0)

	if limit == 0 {
		return sequence
	}

	sequence[1] = big.NewInt(1)
	for i := int64(2); i <= limit; i++ {
		sequence[i] = new(big.Int).Add(sequence[i-1], sequence[i-2])
	}

	return sequence
}
