package sequences

import (
	"fmt"
	"math/big"
)

// Square pyramidal numbers (A000330)
// URL: https://oeis.org/A000330
// Description: a(n) = 0^2 + 1^2 + 2^2 + ... + n^2 = n(n+1)(2n+1)/6.

// GetSquarePyramidalSequence returns the sequence of square pyramidal numbers.
func GetSquarePyramidalSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetSquarePyramidalAtPosition(maxNumber)
	}
	return GenerateSquarePyramidalSequence(maxNumber)
}

// GenerateSquarePyramidalSequence generates square pyramidal numbers up to maxNumber.
func GenerateSquarePyramidalSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Square pyramidal numbers (A000330)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   nil,
		}, nil
	}

	var sequence []*big.Int
	sum := big.NewInt(0)
	for i := int64(0); ; i++ {
		term := new(big.Int).Mul(big.NewInt(i), big.NewInt(i))
		newSum := new(big.Int).Add(sum, term)
		if newSum.Cmp(maxNumber) > 0 {
			break
		}
		sum = newSum
		sequence = append(sequence, new(big.Int).Set(sum))
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Square pyramidal numbers (A000330)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetSquarePyramidalAtPosition returns the n-th square pyramidal number.
func GetSquarePyramidalAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be a non-negative integer")
	}

	// a(n) = n(n+1)(2n+1)/6
	nPlus1 := new(big.Int).Add(n, big.NewInt(1))
	twoNPlus1 := new(big.Int).Mul(big.NewInt(2), n)
	twoNPlus1.Add(twoNPlus1, big.NewInt(1))

	result := new(big.Int).Mul(n, nPlus1)
	result.Mul(result, twoNPlus1)
	result.Div(result, big.NewInt(6))

	return &NumericSequence{
		Name:     "Square pyramidal numbers (A000330)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
