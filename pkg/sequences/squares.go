package sequences

import (
	"fmt"
	"math/big"
)

// The squares: a(n) = n^2.
// URL: https://oeis.org/A000290
// Description: a(n) = n^2.

// GetSquaresSequence returns the squares sequence (OEIS A000290).
func GetSquaresSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetSquareAtPosition(maxNumber)
	}
	return GenerateSquaresSequence(maxNumber)
}

// GenerateSquaresSequence generates the A000290 sequence up to maxNumber.
func GenerateSquaresSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Squares (A000290)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}

	var sequence []*big.Int
	n := big.NewInt(0)
	for {
		square := new(big.Int).Mul(n, n)
		if square.Cmp(maxNumber) > 0 {
			break
		}
		sequence = append(sequence, square)
		n = new(big.Int).Add(n, big.NewInt(1))
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Squares (A000290)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetSquareAtPosition returns the n-th term of A000290.
func GetSquareAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	val := new(big.Int).Mul(n, n)

	return &NumericSequence{
		Name:     "Squares (A000290)",
		Number:   n,
		Sequence: []*big.Int{val},
		Result:   val,
	}, nil
}
