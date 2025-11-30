package sequences

import (
	"fmt"
	"math/big"
)

// GenerateLucas generates the Lucas number sequence up to the given max number or at a specific position.
func GenerateLucas(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetLucasNumberAtPosition(maxNumber)
	} else {
		return GenerateLucasSequence(maxNumber)
	}
}

// GenerateLucasSequence generates the Lucas number sequence up to the given max number.
func GenerateLucasSequence(maxNumber *big.Int) (*NumericSequence, error) {
	var sequence []*big.Int
	a, b := big.NewInt(2), big.NewInt(1)

	for a.Cmp(maxNumber) <= 0 {
		sequence = append(sequence, new(big.Int).Set(a))
		a, b = b, new(big.Int).Add(a, b)
	}

	return &NumericSequence{
		Name:     "Lucas",
		Number:   maxNumber,
		Sequence: sequence,
	}, nil
}

// GetLucasNumberAtPosition returns the Lucas number at the nth position.
func GetLucasNumberAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	a, b := big.NewInt(2), big.NewInt(1)
	if n.Cmp(big.NewInt(0)) == 0 {
		return &NumericSequence{
			Name:     "Lucas",
			Number:   n,
			Sequence: []*big.Int{a},
			Result:   a,
		}, nil
	}
	if n.Cmp(big.NewInt(1)) == 0 {
		return &NumericSequence{
			Name:     "Lucas",
			Number:   n,
			Sequence: []*big.Int{b},
			Result:   b,
		}, nil
	}

	for i := big.NewInt(2); i.Cmp(n) <= 0; i.Add(i, big.NewInt(1)) {
		a, b = b, new(big.Int).Add(a, b)
	}

	return &NumericSequence{
		Name:     "Lucas",
		Number:   n,
		Sequence: []*big.Int{b},
		Result:   b,
	}, nil
}
