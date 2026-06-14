package sequences

import (
	"fmt"
	"math/big"
)

// n^n (A000312)
// URL: https://oeis.org/A000312
// Description: a(n) = n^n. Number of labeled mappings from n points to themselves (endofunctions).

// GetNtoNSequence returns the sequence of n^n numbers.
func GetNtoNSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetNtoNAtPosition(maxNumber)
	}
	return GenerateNtoNSequence(maxNumber)
}

// GenerateNtoNSequence generates n^n numbers up to maxNumber.
func GenerateNtoNSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "n^n (A000312)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   nil,
		}, nil
	}

	var sequence []*big.Int
	for n := int64(0); ; n++ {
		// a(n) = n^n
		val := calculateNtoN(big.NewInt(n))
		if val.Cmp(maxNumber) > 0 {
			break
		}
		sequence = append(sequence, val)
		if n > 1000 { // Safety break for extremely fast growing sequences
			break
		}
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "n^n (A000312)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetNtoNAtPosition returns the n-th n^n number.
func GetNtoNAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be a non-negative integer")
	}

	result := calculateNtoN(n)

	return &NumericSequence{
		Name:     "n^n (A000312)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

func calculateNtoN(n *big.Int) *big.Int {
	// a(n) = n^n
	// Special case 0^0 = 1
	if n.Sign() == 0 {
		return big.NewInt(1)
	}

	res := new(big.Int).Exp(n, n, nil)
	return res
}
