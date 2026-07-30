package sequences

import (
	"fmt"
	"math/big"
)

// A000038 (OEIS): Twice A000007.
// URL: https://oeis.org/A000038

// GetA000038Sequence returns sequence A000038 values.
// If isPositional is true, maxNumber is treated as a position n with offset 0.
// If isPositional is false, maxNumber is treated as the number of terms to generate from n=0.
func GetA000038Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000038AtPosition(maxNumber)
	}
	return GenerateA000038Sequence(maxNumber)
}

// GenerateA000038Sequence generates the first maxNumber terms of A000038.
func GenerateA000038Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Twice A000007 (A000038)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}
	if !maxNumber.IsInt64() {
		return nil, fmt.Errorf("maxNumber too large for current implementation")
	}

	limit := maxNumber.Int64()
	sequence := make([]*big.Int, 0, limit)
	for n := int64(0); n < limit; n++ {
		sequence = append(sequence, a000038Term(n))
	}

	result := big.NewInt(0)
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Twice A000007 (A000038)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetA000038AtPosition returns the term a(n) of A000038 for n >= 0.
func GetA000038AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be non-negative")
	}
	if !n.IsInt64() {
		return nil, fmt.Errorf("position too large for current implementation")
	}

	term := a000038Term(n.Int64())
	return &NumericSequence{
		Name:     "Twice A000007 (A000038)",
		Number:   n,
		Sequence: []*big.Int{term},
		Result:   term,
	}, nil
}

// a000038Term computes a(n) where a(0)=2 and a(n)=0 for n>0.
func a000038Term(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(2)
	}
	return big.NewInt(0)
}
