package sequences

import (
	"fmt"
	"math/big"
)

// Fourth powers: a(n) = n^4 (A000583)
// URL: https://oeis.org/A000583
// Description: a(n) = n^4.

// GetFourthPowersSequence returns the fourth powers sequence (OEIS A000583).
func GetFourthPowersSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetFourthPowerAtPosition(maxNumber)
	}
	return GenerateFourthPowersSequence(maxNumber)
}

// GenerateFourthPowersSequence generates the A000583 sequence up to maxNumber.
func GenerateFourthPowersSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Fourth Powers (A000583)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}

	var sequence []*big.Int
	n := big.NewInt(0)
	for {
		fourthPower := new(big.Int).Exp(n, big.NewInt(4), nil)
		if fourthPower.Cmp(maxNumber) > 0 {
			break
		}
		sequence = append(sequence, fourthPower)
		n = new(big.Int).Add(n, big.NewInt(1))
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Fourth Powers (A000583)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetFourthPowerAtPosition returns the n-th term of A000583.
func GetFourthPowerAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	val := new(big.Int).Exp(n, big.NewInt(4), nil)

	return &NumericSequence{
		Name:     "Fourth Powers (A000583)",
		Number:   n,
		Sequence: []*big.Int{val},
		Result:   val,
	}, nil
}
