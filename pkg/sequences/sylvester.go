package sequences

import (
	"fmt"
	"math/big"
)

// Sylvester's sequence: a(n+1) = a(n)^2 - a(n) + 1, with a(0) = 2.
// URL: https://oeis.org/A000058
// Also called Euclid numbers, because a(n) = a(0)*a(1)*...*a(n-1) + 1 for n > 0, with a(0) = 2.

// GetSylvesterSequence returns Sylvester's sequence up to maxNumber or the n-th term.
func GetSylvesterSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetSylvesterAtPosition(maxNumber)
	}
	return GenerateSylvesterSequence(maxNumber)
}

// GenerateSylvesterSequence generates the A000058 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(n) such that a(n) <= maxNumber.
func GenerateSylvesterSequence(maxNumber *big.Int) (*NumericSequence, error) {
	two := big.NewInt(2)
	if maxNumber.Cmp(two) < 0 {
		return nil, fmt.Errorf("max number must be at least 2 for this sequence")
	}

	sequence := []*big.Int{}
	current := big.NewInt(2)

	for current.Cmp(maxNumber) <= 0 {
		sequence = append(sequence, new(big.Int).Set(current))

		// Next term: a(n+1) = a(n)^2 - a(n) + 1
		next := new(big.Int).Mul(current, current)
		next.Sub(next, current)
		next.Add(next, big.NewInt(1))
		current = next
	}

	return &NumericSequence{
		Name:     "Sylvester's sequence (A000058)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[len(sequence)-1],
	}, nil
}

// GetSylvesterAtPosition returns the n-th term of A000058 (n >= 0).
func GetSylvesterAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be non-negative")
	}

	pos := n.Uint64()
	// To prevent excessive computation/memory, we could limit pos,
	// but let's follow the pattern of other sequences.

	current := big.NewInt(2)
	for i := uint64(0); i < pos; i++ {
		next := new(big.Int).Mul(current, current)
		next.Sub(next, current)
		next.Add(next, big.NewInt(1))
		current = next
	}

	return &NumericSequence{
		Name:     "Sylvester's sequence (A000058)",
		Number:   n,
		Sequence: []*big.Int{current},
		Result:   current,
	}, nil
}
