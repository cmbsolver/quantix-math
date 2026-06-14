package sequences

import (
	"fmt"
	"math/big"
)

// Pentagonal numbers (A000326)
// URL: https://oeis.org/A000326
// Description: a(n) = n(3n-1)/2.

// GetPentagonalSequence returns the sequence of pentagonal numbers.
func GetPentagonalSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetPentagonalAtPosition(maxNumber)
	}
	return GeneratePentagonalSequence(maxNumber)
}

// GeneratePentagonalSequence generates pentagonal numbers up to maxNumber.
func GeneratePentagonalSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Pentagonal numbers (A000326)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   nil,
		}, nil
	}

	var sequence []*big.Int
	for n := int64(0); ; n++ {
		// a(n) = n(3n-1)/2
		val := calculatePentagonal(big.NewInt(n))
		if val.Cmp(maxNumber) > 0 {
			break
		}
		sequence = append(sequence, val)
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Pentagonal numbers (A000326)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetPentagonalAtPosition returns the n-th pentagonal number.
func GetPentagonalAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be a non-negative integer")
	}

	result := calculatePentagonal(n)

	return &NumericSequence{
		Name:     "Pentagonal numbers (A000326)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

func calculatePentagonal(n *big.Int) *big.Int {
	// a(n) = n(3n-1)/2
	threeN := new(big.Int).Mul(big.NewInt(3), n)
	threeNMinus1 := new(big.Int).Sub(threeN, big.NewInt(1))
	result := new(big.Int).Mul(n, threeNMinus1)
	result.Div(result, big.NewInt(2))
	return result
}
