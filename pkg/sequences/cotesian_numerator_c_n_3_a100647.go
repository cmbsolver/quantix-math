package sequences

import (
	"fmt"
	"math/big"
)

// A100647: Numerator of Cotesian number C(n,3).
// URL: https://oeis.org/A100647
// Description: Numerators of Cotesian numbers C(n,3), where n starts at 3.

// GetCotesianNumeratorCN3A100647Sequence returns the A100647 sequence.
func GetCotesianNumeratorCN3A100647Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetCotesianNumeratorCN3A100647AtPosition(maxNumber)
	}

	return GenerateCotesianNumeratorCN3A100647Sequence(maxNumber)
}

// GenerateCotesianNumeratorCN3A100647Sequence generates the first maxNumber terms of A100647.
func GenerateCotesianNumeratorCN3A100647Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	if !maxNumber.IsInt64() {
		return nil, fmt.Errorf("max number too large for current implementation")
	}

	termCount := int(maxNumber.Int64())
	sequence := make([]*big.Int, termCount)

	for i := 0; i < termCount; i++ {
		n := i + 3
		sequence[i] = CalculateCotesianNumeratorA100640(n, 3)
	}

	result := big.NewInt(0)
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Numerator of Cotesian number C(n,3) (A100647)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetCotesianNumeratorCN3A100647AtPosition returns the term at zero-based position n in A100647.
func GetCotesianNumeratorCN3A100647AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	if !n.IsInt64() {
		return nil, fmt.Errorf("position too large for current implementation")
	}

	index := int(n.Int64())
	value := CalculateCotesianNumeratorA100640(index+3, 3)

	return &NumericSequence{
		Name:     "Numerator of Cotesian number C(n,3) (A100647)",
		Number:   n,
		Sequence: []*big.Int{value},
		Result:   value,
	}, nil
}
