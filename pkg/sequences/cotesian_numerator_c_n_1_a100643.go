package sequences

import (
	"fmt"
	"math/big"
)

// A100643: Numerator of Cotesian number C(n,1).
// URL: https://oeis.org/A100643
// Description: Numerators of Cotesian numbers C(n,1), where n starts at 1.

// GetCotesianNumeratorCN1A100643Sequence returns the A100643 sequence.
func GetCotesianNumeratorCN1A100643Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetCotesianNumeratorCN1A100643AtPosition(maxNumber)
	}

	return GenerateCotesianNumeratorCN1A100643Sequence(maxNumber)
}

// GenerateCotesianNumeratorCN1A100643Sequence generates the first maxNumber terms of A100643.
func GenerateCotesianNumeratorCN1A100643Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	if !maxNumber.IsInt64() {
		return nil, fmt.Errorf("max number too large for current implementation")
	}

	termCount := int(maxNumber.Int64())
	sequence := make([]*big.Int, termCount)

	for i := 0; i < termCount; i++ {
		n := i + 1
		sequence[i] = CalculateCotesianNumeratorA100640(n, 1)
	}

	result := big.NewInt(0)
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Numerator of Cotesian number C(n,1) (A100643)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetCotesianNumeratorCN1A100643AtPosition returns the term at zero-based position n in A100643.
func GetCotesianNumeratorCN1A100643AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	if !n.IsInt64() {
		return nil, fmt.Errorf("position too large for current implementation")
	}

	index := int(n.Int64())
	value := CalculateCotesianNumeratorA100640(index+1, 1)

	return &NumericSequence{
		Name:     "Numerator of Cotesian number C(n,1) (A100643)",
		Number:   n,
		Sequence: []*big.Int{value},
		Result:   value,
	}, nil
}
