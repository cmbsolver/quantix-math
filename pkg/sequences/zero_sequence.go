package sequences

import (
	"fmt"
	"math/big"
)

// Zero Sequence (OEIS A000004)
// URL: https://oeis.org/A000004
// a(n) = 0.

// GetZeroSequence returns the zero sequence (OEIS A000004).
func GetZeroSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetZeroSequenceAtPosition(maxNumber)
	}
	return GenerateZeroSequence(maxNumber)
}

// GenerateZeroSequence generates the A000004 sequence up to maxNumber (a(0), a(1), ..., a(maxNumber)).
func GenerateZeroSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		sequence[i] = big.NewInt(0)
	}

	result := big.NewInt(0)

	return &NumericSequence{
		Name:     "Zero Sequence (A000004)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetZeroSequenceAtPosition returns the n-th term of A000004 (n >= 0).
func GetZeroSequenceAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	result := big.NewInt(0)

	return &NumericSequence{
		Name:     "Zero Sequence (A000004)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
