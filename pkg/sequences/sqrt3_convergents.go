package sequences

import (
	"fmt"
	"math/big"
)

// Sqrt3Convergents A002531: Numerators of continued fraction convergents to sqrt(3).
// URL: https://oeis.org/A002531
// Recurrence:
// a(0) = 1, a(1) = 1
// a(2n) = a(2n-1) + a(2n-2)
// a(2n+1) = 2*a(2n) + a(2n-1)

// GetSqrt3ConvergentsA002531Sequence returns the A002531 sequence.
func GetSqrt3ConvergentsA002531Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetSqrt3ConvergentsA002531AtPosition(maxNumber)
	}
	return GenerateSqrt3ConvergentsA002531Sequence(maxNumber)
}

// GenerateSqrt3ConvergentsA002531Sequence generates the A002531 sequence up to maxNumber.
func GenerateSqrt3ConvergentsA002531Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	results := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		if i == 0 {
			results[i] = big.NewInt(1)
		} else if i == 1 {
			results[i] = big.NewInt(1)
		} else if i%2 == 0 {
			// a(2k) = a(2k-1) + a(2k-2)
			results[i] = new(big.Int).Add(results[i-1], results[i-2])
		} else {
			// a(2k+1) = 2*a(2k) + a(2k-1)
			twoAk := new(big.Int).Mul(big.NewInt(2), results[i-1])
			results[i] = new(big.Int).Add(twoAk, results[i-2])
		}
	}

	return &NumericSequence{
		Name:     "Numerators of convergents to sqrt(3) (A002531)",
		Number:   maxNumber,
		Sequence: results,
		Result:   results[n],
	}, nil
}

// GetSqrt3ConvergentsA002531AtPosition returns the n-th term of A002531.
func GetSqrt3ConvergentsA002531AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	pos := int(n.Int64())
	// Use iterative approach to find the term
	aPrev2 := big.NewInt(1) // a(0)
	if pos == 0 {
		return &NumericSequence{
			Name:     "Numerators of convergents to sqrt(3) (A002531)",
			Number:   n,
			Sequence: []*big.Int{aPrev2},
			Result:   aPrev2,
		}, nil
	}

	aPrev1 := big.NewInt(1) // a(1)
	if pos == 1 {
		return &NumericSequence{
			Name:     "Numerators of convergents to sqrt(3) (A002531)",
			Number:   n,
			Sequence: []*big.Int{aPrev1},
			Result:   aPrev1,
		}, nil
	}

	var current *big.Int
	for i := 2; i <= pos; i++ {
		if i%2 == 0 {
			current = new(big.Int).Add(aPrev1, aPrev2)
		} else {
			twoAk := new(big.Int).Mul(big.NewInt(2), aPrev1)
			current = new(big.Int).Add(twoAk, aPrev2)
		}
		aPrev2.Set(aPrev1)
		aPrev1.Set(current)
	}

	return &NumericSequence{
		Name:     "Numerators of convergents to sqrt(3) (A002531)",
		Number:   n,
		Sequence: []*big.Int{aPrev1},
		Result:   aPrev1,
	}, nil
}
