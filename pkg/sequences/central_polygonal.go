package sequences

import (
	"fmt"
	"math/big"
)

// Central polygonal numbers (the Lazy Caterer's sequence) (A000124)
// URL: https://oeis.org/A000124
// Description: a(n) = n(n+1)/2 + 1; or, maximal number of pieces formed when slicing a pancake with n cuts.

// GetCentralPolygonalNumbersSequence returns the sequence of central polygonal numbers.
func GetCentralPolygonalNumbersSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetCentralPolygonalAtPosition(maxNumber)
	}
	return GenerateCentralPolygonalSequence(maxNumber)
}

// GenerateCentralPolygonalSequence generates central polygonal numbers up to maxNumber.
func GenerateCentralPolygonalSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Central polygonal numbers (A000124)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   nil,
		}, nil
	}

	var sequence []*big.Int
	current := big.NewInt(1)
	for i := int64(0); ; i++ {
		// a(n) = n*(n+1)/2 + 1
		// Alternatively, a(0)=1, a(n) = a(n-1) + n
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
		Name:     "Central polygonal numbers (A000124)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetCentralPolygonalAtPosition returns the n-th central polygonal number.
func GetCentralPolygonalAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be a non-negative integer")
	}

	// a(n) = n*(n+1)/2 + 1
	nPlus1 := new(big.Int).Add(n, big.NewInt(1))
	result := new(big.Int).Mul(n, nPlus1)
	result.Div(result, big.NewInt(2))
	result.Add(result, big.NewInt(1))

	return &NumericSequence{
		Name:     "Central polygonal numbers (A000124)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
