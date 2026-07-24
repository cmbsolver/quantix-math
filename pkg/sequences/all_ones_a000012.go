package sequences

import (
	"fmt"
	"math/big"
)

// The simplest sequence of positive numbers: the all 1's sequence.
// URL: https://oeis.org/A000012
// Description: a(n) = 1.

// GetAllOnesA000012Sequence returns the A000012 sequence.
func GetAllOnesA000012Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetAllOnesA000012AtPosition(maxNumber)
	}
	return GenerateAllOnesA000012Sequence(maxNumber)
}

// GenerateAllOnesA000012Sequence generates the A000012 sequence up to maxNumber (n terms).
func GenerateAllOnesA000012Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "All 1's (A000012)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}

	limit := maxNumber.Int64()
	if maxNumber.IsInt64() && limit >= 0 {
		var sequence []*big.Int
		one := big.NewInt(1)
		for i := int64(0); i < limit; i++ {
			sequence = append(sequence, one)
		}
		var result *big.Int
		if len(sequence) > 0 {
			result = one
		} else {
			result = big.NewInt(0)
		}
		return &NumericSequence{
			Name:     "All 1's (A000012)",
			Number:   maxNumber,
			Sequence: sequence,
			Result:   result,
		}, nil
	}

	return nil, fmt.Errorf("maxNumber too large for current implementation")
}

// GetAllOnesA000012AtPosition returns the n-th term of A000012.
func GetAllOnesA000012AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be non-negative")
	}

	val := big.NewInt(1)
	return &NumericSequence{
		Name:     "All 1's (A000012)",
		Number:   n,
		Sequence: []*big.Int{val},
		Result:   val,
	}, nil
}
