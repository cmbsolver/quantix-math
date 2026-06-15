package sequences

import (
	"fmt"
	"math/big"
)

// ParitySequence represents the parity of n (A000035).
// a(n) = n mod 2.
// Period 2: repeat [0, 1].
// URL: https://oeis.org/A000035
type ParitySequence struct{}

// GetParitySequence generates the parity sequence up to maxNumber or at a specific position.
// If isPositional is true, it returns the n-th term (where n is maxNumber).
// If isPositional is false, it returns all terms up to maxNumber.
func GetParitySequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	numericSequence := &NumericSequence{
		Name:   "n mod 2; parity of n (A000035)",
		Number: new(big.Int).Set(maxNumber),
	}

	if isPositional {
		// Calculate a(n) = n mod 2
		result := new(big.Int).Mod(maxNumber, big.NewInt(2))
		numericSequence.Sequence = append(numericSequence.Sequence, result)
	} else {
		// Calculate all terms from a(0) to a(maxNumber)
		n := maxNumber.Int64()
		for i := int64(0); i <= n; i++ {
			numericSequence.Sequence = append(numericSequence.Sequence, big.NewInt(i%2))
		}
	}

	return numericSequence, nil
}
