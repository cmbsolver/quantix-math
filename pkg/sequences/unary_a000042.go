package sequences

import (
	"fmt"
	"math/big"
	"strings"
)

// Unary representation of natural numbers.
// URL: https://oeis.org/A000042
// Description: a(n) = n 1's.

// GetUnaryA000042Sequence returns the A000042 sequence.
func GetUnaryA000042Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetUnaryA000042AtPosition(maxNumber)
	}
	return GenerateUnaryA000042Sequence(maxNumber)
}

// GenerateUnaryA000042Sequence generates the A000042 sequence up to maxNumber.
func GenerateUnaryA000042Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Unary representation (A000042)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}

	limit := maxNumber.Int64()
	if maxNumber.IsInt64() && limit >= 0 {
		var sequence []*big.Int
		for n := int64(1); n <= limit; n++ {
			val, _ := getUnary(n)
			sequence = append(sequence, val)
		}
		var result *big.Int
		if len(sequence) > 0 {
			result = sequence[len(sequence)-1]
		} else {
			result = big.NewInt(0)
		}
		return &NumericSequence{
			Name:     "Unary representation (A000042)",
			Number:   maxNumber,
			Sequence: sequence,
			Result:   result,
		}, nil
	}

	return nil, fmt.Errorf("maxNumber too large for current implementation")
}

// GetUnaryA000042AtPosition returns the n-th term of A000042.
func GetUnaryA000042AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() <= 0 {
		return nil, fmt.Errorf("position must be positive")
	}

	val, err := getUnary(n.Int64())
	if err != nil {
		return nil, err
	}
	return &NumericSequence{
		Name:     "Unary representation (A000042)",
		Number:   n,
		Sequence: []*big.Int{val},
		Result:   val,
	}, nil
}

func getUnary(n int64) (*big.Int, error) {
	if n <= 0 {
		return big.NewInt(0), nil
	}
	s := strings.Repeat("1", int(n))
	val := new(big.Int)
	_, ok := val.SetString(s, 10)
	if !ok {
		return nil, fmt.Errorf("failed to create unary number for n=%d", n)
	}
	return val, nil
}
