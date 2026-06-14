package sequences

import (
	"fmt"
	"math/big"
)

// Mersenne numbers (OEIS A000225)
// URL: https://oeis.org/A000225
// a(n) = 2^n - 1.

// GetMersenneNumbersSequence returns the Mersenne numbers sequence (OEIS A000225).
// If isPositional is true, it returns the n-th term.
// Otherwise, it returns all terms up to maxNumber.
func GetMersenneNumbersSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetMersenneNumbersAtPosition(maxNumber)
	}
	return GenerateMersenneNumbersSequence(maxNumber)
}

// GenerateMersenneNumbersSequence generates the A000225 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(n) such that a(n) <= maxNumber.
func GenerateMersenneNumbersSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	var sequence []*big.Int
	two := big.NewInt(2)
	one := big.NewInt(1)

	// a(n) = 2^n - 1
	// We can generate it iteratively: a(n) = 2 * a(n-1) + 1
	current := big.NewInt(0) // a(0) = 2^0 - 1 = 0

	for current.Cmp(maxNumber) <= 0 {
		sequence = append(sequence, new(big.Int).Set(current))

		// next = 2 * current + 1
		next := new(big.Int).Mul(current, two)
		next.Add(next, one)
		current = next
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	} else {
		result = big.NewInt(0)
	}

	return &NumericSequence{
		Name:     "Mersenne numbers (A000225)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetMersenneNumbersAtPosition returns the n-th term of A000225 (n >= 0).
func GetMersenneNumbersAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	// a(n) = 2^n - 1
	two := big.NewInt(2)
	result := new(big.Int).Exp(two, n, nil)
	result.Sub(result, big.NewInt(1))

	return &NumericSequence{
		Name:     "Mersenne numbers (A000225)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
