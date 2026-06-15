package sequences

import (
	"fmt"
	"math/big"
)

// Bell numbers (OEIS A000110)
// URL: https://oeis.org/A000110
// Bell or exponential numbers: number of ways to partition a set of n labeled elements.

// GetBellSequence returns the Bell numbers sequence up to maxNumber (OEIS A000110).
func GetBellSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetBellNumberAtPosition(maxNumber)
	}
	return GenerateBellSequence(maxNumber)
}

// GenerateBellSequence generates the A000110 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateBellSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number must be at least 0")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		sequence[i] = calculateBell(int64(i))
	}

	return &NumericSequence{
		Name:     "Bell numbers (A000110)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetBellNumberAtPosition returns the n-th term of A000110 (n >= 0).
func GetBellNumberAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	result := calculateBell(n.Int64())

	return &NumericSequence{
		Name:     "Bell numbers (A000110)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateBell calculates the n-th Bell number using the recurrence relation:
// B(n+1) = sum_{k=0..n} C(n, k) * B(k), with B(0) = 1.
func calculateBell(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	bells := make([]*big.Int, n+1)
	bells[0] = big.NewInt(1)

	for i := int64(0); i < n; i++ {
		nextBell := big.NewInt(0)
		for k := int64(0); k <= i; k++ {
			// C(i, k) * B(k)
			term := new(big.Int).Binomial(i, k)
			term.Mul(term, bells[k])
			nextBell.Add(nextBell, term)
		}
		bells[i+1] = nextBell
	}

	return bells[n]
}
