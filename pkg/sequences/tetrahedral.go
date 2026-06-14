package sequences

import (
	"fmt"
	"math/big"
)

// Tetrahedral (or triangular pyramidal) numbers (A000292)
// URL: https://oeis.org/A000292
// Description: a(n) = C(n+2,3) = n(n+1)(n+2)/6.

// GetTetrahedralSequence returns the sequence of tetrahedral numbers.
func GetTetrahedralSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetTetrahedralAtPosition(maxNumber)
	}
	return GenerateTetrahedralSequence(maxNumber)
}

// GenerateTetrahedralSequence generates tetrahedral numbers up to maxNumber.
func GenerateTetrahedralSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Tetrahedral numbers (A000292)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   nil,
		}, nil
	}

	var sequence []*big.Int
	for n := int64(0); ; n++ {
		term := calculateTetrahedral(big.NewInt(n))
		if term.Cmp(maxNumber) > 0 {
			break
		}
		sequence = append(sequence, term)
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Tetrahedral numbers (A000292)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetTetrahedralAtPosition returns the n-th tetrahedral number.
func GetTetrahedralAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be a non-negative integer")
	}

	result := calculateTetrahedral(n)

	return &NumericSequence{
		Name:     "Tetrahedral numbers (A000292)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateTetrahedral calculates the n-th tetrahedral number using the formula a(n) = n(n+1)(n+2)/6.
func calculateTetrahedral(n *big.Int) *big.Int {
	// a(n) = n(n+1)(n+2)/6
	nPlus1 := new(big.Int).Add(n, big.NewInt(1))
	nPlus2 := new(big.Int).Add(n, big.NewInt(2))

	result := new(big.Int).Mul(n, nPlus1)
	result.Mul(result, nPlus2)
	result.Div(result, big.NewInt(6))

	return result
}
