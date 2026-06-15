package sequences

import (
	"fmt"
	"math/big"
)

// Sqrt3ConvergentsDenominators Sequence (OEIS A002530)
// URL: https://oeis.org/A002530
// a(n) = 4*a(n-2) - a(n-4) for n > 3, with a(0)=0, a(1)=1, a(2)=1, a(3)=3.
// Denominators of continued fraction convergents to sqrt(3), for n >= 1.

// GetSqrt3ConvergentsDenominatorsSequence returns the denominators of continued fraction convergents to sqrt(3) (OEIS A002530).
func GetSqrt3ConvergentsDenominatorsSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetSqrt3ConvergentsDenominatorsAtPosition(maxNumber)
	}
	return GenerateSqrt3ConvergentsDenominatorsSequence(maxNumber)
}

// GenerateSqrt3ConvergentsDenominatorsSequence generates the A002530 sequence up to maxNumber (a(0), a(1), ..., a(maxNumber)).
func GenerateSqrt3ConvergentsDenominatorsSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		sequence[i] = calculateSqrt3ConvergentsDenominator(int64(i))
	}

	result := big.NewInt(0)
	if n >= 0 {
		result = sequence[n]
	}

	return &NumericSequence{
		Name:     "Sqrt(3) Convergents Denominators (A002530)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetSqrt3ConvergentsDenominatorsAtPosition returns the n-th term of A002530 (n >= 0).
func GetSqrt3ConvergentsDenominatorsAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	result := calculateSqrt3ConvergentsDenominator(n.Int64())

	return &NumericSequence{
		Name:     "Sqrt(3) Convergents Denominators (A002530)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateSqrt3ConvergentsDenominator calculates the n-th term of A002530.
func calculateSqrt3ConvergentsDenominator(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	if n == 1 {
		return big.NewInt(1)
	}
	if n == 2 {
		return big.NewInt(1)
	}
	if n == 3 {
		return big.NewInt(3)
	}

	// a(n) = 4*a(n-2) - a(n-4)
	a0 := big.NewInt(0)
	a1 := big.NewInt(1)
	a2 := big.NewInt(1)
	a3 := big.NewInt(3)

	var current *big.Int
	for i := int64(4); i <= n; i++ {
		current = new(big.Int).Mul(big.NewInt(4), a2)
		current.Sub(current, a0)

		// Shift
		a0, a1, a2, a3 = a1, a2, a3, current
	}
	return a3
}
