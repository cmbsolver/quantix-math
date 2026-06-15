package sequences

import (
	"fmt"
	"math/big"
)

// Composite numbers: numbers n of the form x*y for x > 1 and y > 1 (OEIS A002808).
// URL: https://oeis.org/A002808
// The sequence starts: 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20, 21, 22, 24, 25, 26, 27, 28, 30, 32, 33, 34, 35, 36, 38, 39, 40, 42, 44, 45, 46, 48, 49, 50, 51, 52, 54, 55, 56, 57, 58, 60, 62, 63, 64, 65, 66, 68, 69, 70, 72, 74, 75, 76, 77, 78, 80, 81, 82, 84, 85, 86, 87, 88

// GetCompositesSequence returns the composite numbers up to maxNumber (OEIS A002808).
func GetCompositesSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetCompositesAtPosition(maxNumber)
	}
	return GenerateCompositesSequence(maxNumber)
}

// GenerateCompositesSequence generates the A002808 sequence up to maxNumber.
// It returns all composite numbers <= maxNumber.
func GenerateCompositesSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(4)) < 0 {
		return &NumericSequence{
			Name:     "Composite numbers (A002808)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   big.NewInt(0),
		}, nil
	}

	n := maxNumber.Int64()
	var sequence []*big.Int

	for i := int64(4); i <= n; i++ {
		if isComposite(i) {
			sequence = append(sequence, big.NewInt(i))
		}
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	} else {
		result = big.NewInt(0)
	}

	return &NumericSequence{
		Name:     "Composite numbers (A002808)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetCompositesAtPosition returns the n-th composite number (n >= 1).
func GetCompositesAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("position must be at least 1")
	}

	pos := n.Int64()
	var count int64 = 0
	var i int64 = 4
	for {
		if isComposite(i) {
			count++
			if count == pos {
				result := big.NewInt(i)
				return &NumericSequence{
					Name:     "Composite numbers (A002808)",
					Number:   n,
					Sequence: []*big.Int{result},
					Result:   result,
				}, nil
			}
		}
		i++
	}
}

// isComposite checks if a number is composite.
// A number n is composite if n > 1 and it is not prime.
func isComposite(n int64) bool {
	if n <= 3 {
		return false
	}
	// Use IsPrime from the package if available, or simple primality test
	return !IsPrime(big.NewInt(n))
}
