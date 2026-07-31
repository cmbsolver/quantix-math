package sequences

import (
	"fmt"
	"math/big"
)

// Jacobsthal numbers: a(n) = a(n-1) + 2*a(n-2), with a(0) = 0, a(1) = 1 (OEIS A001045).
// URL: https://oeis.org/A001045

// GetA001045Sequence returns the A001045 sequence.
func GetA001045Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA001045AtPosition(maxNumber)
	}

	return GenerateA001045Sequence(maxNumber)
}

// GenerateA001045Sequence generates A001045 from a(0) through a(maxNumber).
func GenerateA001045Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := maxNumber.Int64()
	sequence := generateA001045Terms(limit)

	return &NumericSequence{
		Name:     "Jacobsthal numbers (A001045)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[limit],
	}, nil
}

// GetA001045AtPosition returns the n-th term of A001045.
func GetA001045AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	sequence := generateA001045Terms(n.Int64())
	result := sequence[n.Int64()]

	return &NumericSequence{
		Name:     "Jacobsthal numbers (A001045)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// generateA001045Terms computes terms a(0) through a(limit) for OEIS A001045.
func generateA001045Terms(limit int64) []*big.Int {
	sequence := make([]*big.Int, limit+1)
	sequence[0] = big.NewInt(0)

	if limit == 0 {
		return sequence
	}

	sequence[1] = big.NewInt(1)
	for i := int64(2); i <= limit; i++ {
		twoPrev2 := new(big.Int).Mul(big.NewInt(2), sequence[i-2])
		sequence[i] = new(big.Int).Add(sequence[i-1], twoPrev2)
	}

	return sequence
}
