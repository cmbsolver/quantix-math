package sequences

import (
	"fmt"
	"math/big"
)

// GetPowersOf4Sequence returns the powers of 4 sequence (A000302) up to a max number or at a specific position.
// URL: https://oeis.org/A000302
// Powers of 4: a(n) = 4^n.
func GetPowersOf4Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetPowerOf4AtPosition(maxNumber)
	}
	return GeneratePowersOf4Sequence(maxNumber)
}

// GeneratePowersOf4Sequence generates the powers of 4 sequence up to the given max number.
func GeneratePowersOf4Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	var sequence []*big.Int
	four := big.NewInt(4)
	current := big.NewInt(1)

	for current.Cmp(maxNumber) <= 0 {
		sequence = append(sequence, new(big.Int).Set(current))
		current = new(big.Int).Mul(current, four)
	}

	return &NumericSequence{
		Name:     "Powers of 4 (A000302)",
		Number:   maxNumber,
		Sequence: sequence,
	}, nil
}

// GetPowerOf4AtPosition returns the power of 4 at the nth position (4^n).
func GetPowerOf4AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	four := big.NewInt(4)
	result := new(big.Int).Exp(four, n, nil)

	return &NumericSequence{
		Name:     "Powers of 4 (A000302)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
