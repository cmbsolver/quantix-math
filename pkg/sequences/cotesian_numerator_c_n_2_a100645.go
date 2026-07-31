package sequences

import (
	"fmt"
	"math/big"
)

// A100645: Numerator of Cotesian number C(n,2).
// URL: https://oeis.org/A100645
// Description: Numerators of Cotesian numbers C(n,2), where n starts at 2.

// GetCotesianNumeratorCN2A100645Sequence returns the A100645 sequence.
func GetCotesianNumeratorCN2A100645Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetCotesianNumeratorCN2A100645AtPosition(maxNumber)
	}

	return GenerateCotesianNumeratorCN2A100645Sequence(maxNumber)
}

// GenerateCotesianNumeratorCN2A100645Sequence generates the first maxNumber terms of A100645.
func GenerateCotesianNumeratorCN2A100645Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	if !maxNumber.IsInt64() {
		return nil, fmt.Errorf("max number too large for current implementation")
	}

	termCount := int(maxNumber.Int64())
	sequence := make([]*big.Int, termCount)

	for i := 0; i < termCount; i++ {
		n := i + 2
		sequence[i] = CalculateCotesianNumeratorA100640(n, 2)
	}

	result := big.NewInt(0)
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Numerator of Cotesian number C(n,2) (A100645)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetCotesianNumeratorCN2A100645AtPosition returns the term at zero-based position n in A100645.
func GetCotesianNumeratorCN2A100645AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	if !n.IsInt64() {
		return nil, fmt.Errorf("position too large for current implementation")
	}

	index := int(n.Int64())
	value := CalculateCotesianNumeratorA100640(index+2, 2)

	return &NumericSequence{
		Name:     "Numerator of Cotesian number C(n,2) (A100645)",
		Number:   n,
		Sequence: []*big.Int{value},
		Result:   value,
	}, nil
}
