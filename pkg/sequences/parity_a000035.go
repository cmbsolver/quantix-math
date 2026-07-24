package sequences

import (
	"fmt"
	"math/big"
)

// Period 2: repeat [0, 1].
// URL: https://oeis.org/A000035
// Description: a(n) = n mod 2.

// GetParityA000035Sequence returns the A000035 sequence.
func GetParityA000035Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetParityA000035AtPosition(maxNumber)
	}
	return GenerateParityA000035Sequence(maxNumber)
}

// GenerateParityA000035Sequence generates the A000035 sequence up to maxNumber.
func GenerateParityA000035Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Parity (A000035)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}

	limit := maxNumber.Int64()
	if maxNumber.IsInt64() && limit >= 0 {
		var sequence []*big.Int
		for n := int64(0); n < limit; n++ {
			val := n % 2
			sequence = append(sequence, big.NewInt(val))
		}
		var result *big.Int
		if len(sequence) > 0 {
			result = sequence[len(sequence)-1]
		} else {
			result = big.NewInt(0)
		}
		return &NumericSequence{
			Name:     "Parity (A000035)",
			Number:   maxNumber,
			Sequence: sequence,
			Result:   result,
		}, nil
	}

	return nil, fmt.Errorf("maxNumber too large for current implementation")
}

// GetParityA000035AtPosition returns the n-th term of A000035.
func GetParityA000035AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be non-negative")
	}

	val := new(big.Int).Mod(n, big.NewInt(2))
	return &NumericSequence{
		Name:     "Parity (A000035)",
		Number:   n,
		Sequence: []*big.Int{val},
		Result:   val,
	}, nil
}
