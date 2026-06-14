package sequences

import (
	"fmt"
	"math/big"
)

// Powers of 3 (OEIS A000244)
// URL: https://oeis.org/A000244
// a(n) = 3^n.

// GetPowersOf3Sequence returns the powers of 3 sequence (OEIS A000244).
func GetPowersOf3Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetPowersOf3AtPosition(maxNumber)
	}
	return GeneratePowersOf3Sequence(maxNumber)
}

// GeneratePowersOf3Sequence generates the A000244 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(n) such that a(n) <= maxNumber.
func GeneratePowersOf3Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	var sequence []*big.Int
	three := big.NewInt(3)
	current := big.NewInt(1) // 3^0

	for current.Cmp(maxNumber) <= 0 {
		sequence = append(sequence, new(big.Int).Set(current))
		current.Mul(current, three)
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	} else {
		result = big.NewInt(0)
	}

	return &NumericSequence{
		Name:     "Powers of 3 (A000244)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetPowersOf3AtPosition returns the n-th term of A000244 (n >= 0).
func GetPowersOf3AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	three := big.NewInt(3)
	result := new(big.Int).Exp(three, n, nil)

	return &NumericSequence{
		Name:     "Powers of 3 (A000244)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
