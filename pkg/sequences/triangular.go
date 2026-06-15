package sequences

import (
	"fmt"
	"math/big"
)

// Triangular numbers (A000217)
// URL: https://oeis.org/A000217
// Description: a(n) = n*(n+1)/2 = 0 + 1 + 2 + ... + n.

// GetTriangularSequence returns the sequence of triangular numbers.
func GetTriangularSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetTriangularAtPosition(maxNumber)
	}
	return GenerateTriangularSequence(maxNumber)
}

// GenerateTriangularSequence generates triangular numbers up to maxNumber.
func GenerateTriangularSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Triangular numbers (A000217)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   nil,
		}, nil
	}

	var sequence []*big.Int
	current := big.NewInt(0)
	for i := int64(0); ; i++ {
		// a(n) = n*(n+1)/2
		// Alternatively, a(n) = a(n-1) + n
		if i > 0 {
			current = new(big.Int).Add(current, big.NewInt(i))
		}

		if current.Cmp(maxNumber) > 0 {
			break
		}
		sequence = append(sequence, new(big.Int).Set(current))
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Triangular numbers (A000217)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetTriangularAtPosition returns the n-th triangular number.
func GetTriangularAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be a non-negative integer")
	}

	// a(n) = n*(n+1)/2
	nPlus1 := new(big.Int).Add(n, big.NewInt(1))
	result := new(big.Int).Mul(n, nPlus1)
	result.Div(result, big.NewInt(2))

	return &NumericSequence{
		Name:     "Triangular numbers (A000217)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
