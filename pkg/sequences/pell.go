package sequences

import (
	"fmt"
	"math/big"
)

// Pell numbers (A000129)
// URL: https://oeis.org/A000129
// Description: a(0) = 0, a(1) = 1; for n > 1, a(n) = 2*a(n-1) + a(n-2).

// GetPellSequence returns the Pell sequence.
func GetPellSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetPellAtPosition(maxNumber)
	}
	return GeneratePellSequence(maxNumber)
}

// GeneratePellSequence generates Pell numbers up to maxNumber.
func GeneratePellSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Pell numbers (A000129)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   nil,
		}, nil
	}

	var sequence []*big.Int
	a := big.NewInt(0)
	b := big.NewInt(1)

	if a.Cmp(maxNumber) <= 0 {
		sequence = append(sequence, new(big.Int).Set(a))
	}
	if b.Cmp(maxNumber) <= 0 {
		sequence = append(sequence, new(big.Int).Set(b))
	}

	for {
		// next = 2*b + a
		twoB := new(big.Int).Mul(big.NewInt(2), b)
		next := new(big.Int).Add(twoB, a)

		if next.Cmp(maxNumber) > 0 {
			break
		}
		sequence = append(sequence, new(big.Int).Set(next))
		a.Set(b)
		b.Set(next)
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Pell numbers (A000129)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetPellAtPosition returns the n-th Pell number.
func GetPellAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be a non-negative integer")
	}

	if n.Cmp(big.NewInt(0)) == 0 {
		result := big.NewInt(0)
		return &NumericSequence{
			Name:     "Pell numbers (A000129)",
			Number:   n,
			Sequence: []*big.Int{result},
			Result:   result,
		}, nil
	}

	if n.Cmp(big.NewInt(1)) == 0 {
		result := big.NewInt(1)
		return &NumericSequence{
			Name:     "Pell numbers (A000129)",
			Number:   n,
			Sequence: []*big.Int{result},
			Result:   result,
		}, nil
	}

	a := big.NewInt(0)
	b := big.NewInt(1)
	result := big.NewInt(0)

	for i := big.NewInt(2); i.Cmp(n) <= 0; i.Add(i, big.NewInt(1)) {
		twoB := new(big.Int).Mul(big.NewInt(2), b)
		result = new(big.Int).Add(twoB, a)
		a.Set(b)
		b.Set(result)
	}

	return &NumericSequence{
		Name:     "Pell numbers (A000129)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
