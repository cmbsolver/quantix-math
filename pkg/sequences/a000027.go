package sequences

import (
	"fmt"
	"math/big"
)

// A000027: The positive integers.
// URL: https://oeis.org/A000027
// a(n) = n.

// GetA000027Sequence returns the positive integers sequence (OEIS A000027).
func GetA000027Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000027AtPosition(maxNumber)
	}
	return GenerateA000027Sequence(maxNumber)
}

// GenerateA000027Sequence generates the A000027 sequence up to maxNumber.
func GenerateA000027Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("max number must be at least 1")
	}

	limit := int(maxNumber.Int64())
	// Use a reasonable limit for calculation if it exceeds Int64 or is too large
	if !maxNumber.IsInt64() || limit > 10000 {
		limit = 10000
	}

	sequence := make([]*big.Int, limit)
	for i := 1; i <= limit; i++ {
		sequence[i-1] = big.NewInt(int64(i))
	}

	return &NumericSequence{
		Name:     "Positive integers (A000027)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[len(sequence)-1],
	}, nil
}

// GetA000027AtPosition returns the n-th term of A000027.
func GetA000027AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("position must be at least 1")
	}

	result := new(big.Int).Set(n)

	return &NumericSequence{
		Name:     "Positive integers (A000027)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
