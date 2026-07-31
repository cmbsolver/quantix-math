package sequences

import (
	"fmt"
	"math/big"
)

// A100646: Denominator of Cotesian number C(n,2).
// URL: https://oeis.org/A100646
// Description: Denominators of Cotesian numbers C(n,2), where n starts at 2.

// GetCotesianDenominatorCN2A100646Sequence returns the A100646 sequence.
func GetCotesianDenominatorCN2A100646Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetCotesianDenominatorCN2A100646AtPosition(maxNumber)
	}

	return GenerateCotesianDenominatorCN2A100646Sequence(maxNumber)
}

// GenerateCotesianDenominatorCN2A100646Sequence generates the first maxNumber terms of A100646.
func GenerateCotesianDenominatorCN2A100646Sequence(maxNumber *big.Int) (*NumericSequence, error) {
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
		sequence[i] = CalculateCotesianDenominatorA100641(n, 2)
	}

	result := big.NewInt(0)
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Denominator of Cotesian number C(n,2) (A100646)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetCotesianDenominatorCN2A100646AtPosition returns the term at zero-based position n in A100646.
func GetCotesianDenominatorCN2A100646AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	if !n.IsInt64() {
		return nil, fmt.Errorf("position too large for current implementation")
	}

	index := int(n.Int64())
	value := CalculateCotesianDenominatorA100641(index+2, 2)

	return &NumericSequence{
		Name:     "Denominator of Cotesian number C(n,2) (A100646)",
		Number:   n,
		Sequence: []*big.Int{value},
		Result:   value,
	}, nil
}
