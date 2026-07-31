package sequences

import (
	"fmt"
	"math/big"
)

// A100648: Denominator of Cotesian number C(n,3).
// URL: https://oeis.org/A100648
// Description: Denominators of Cotesian numbers C(n,3), where n starts at 3.

// GetCotesianDenominatorCN3A100648Sequence returns the A100648 sequence.
func GetCotesianDenominatorCN3A100648Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetCotesianDenominatorCN3A100648AtPosition(maxNumber)
	}

	return GenerateCotesianDenominatorCN3A100648Sequence(maxNumber)
}

// GenerateCotesianDenominatorCN3A100648Sequence generates the first maxNumber terms of A100648.
func GenerateCotesianDenominatorCN3A100648Sequence(maxNumber *big.Int) (*NumericSequence, error) {
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
		sequence[i] = CalculateCotesianDenominatorA100641(n, 3)
	}

	result := big.NewInt(0)
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Denominator of Cotesian number C(n,3) (A100648)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetCotesianDenominatorCN3A100648AtPosition returns the term at zero-based position n in A100648.
func GetCotesianDenominatorCN3A100648AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	if !n.IsInt64() {
		return nil, fmt.Errorf("position too large for current implementation")
	}

	index := int(n.Int64())
	value := CalculateCotesianDenominatorA100641(index+3, 3)

	return &NumericSequence{
		Name:     "Denominator of Cotesian number C(n,3) (A100648)",
		Number:   n,
		Sequence: []*big.Int{value},
		Result:   value,
	}, nil
}
