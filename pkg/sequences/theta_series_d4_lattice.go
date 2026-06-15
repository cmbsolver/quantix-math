package sequences

import (
	"fmt"
	"math/big"
)

// Theta series of D_4 lattice (OEIS A004011)
// URL: https://oeis.org/A004011
// Description: a(0) = 1; for n > 0, a(n) = 24 * (sum of odd divisors of n).
// Also Fourier coefficients of Eisenstein series E_{gamma,2}.

// GetThetaSeriesD4LatticeSequence returns the Theta series of D_4 lattice (OEIS A004011).
func GetThetaSeriesD4LatticeSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetThetaSeriesD4LatticeAtPosition(maxNumber)
	}
	return GenerateThetaSeriesD4LatticeSequence(maxNumber)
}

// GenerateThetaSeriesD4LatticeSequence generates the A004011 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateThetaSeriesD4LatticeSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number must be at least 0")
	}

	if !maxNumber.IsInt64() {
		return nil, fmt.Errorf("max number too large")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		sequence[i] = calculateThetaSeriesD4Lattice(int64(i))
	}

	return &NumericSequence{
		Name:     "Theta series of D_4 lattice (A004011)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetThetaSeriesD4LatticeAtPosition returns the n-th term of A004011 (n >= 0).
func GetThetaSeriesD4LatticeAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	if !n.IsInt64() {
		return nil, fmt.Errorf("position too large")
	}

	val := calculateThetaSeriesD4Lattice(n.Int64())

	return &NumericSequence{
		Name:     "Theta series of D_4 lattice (A004011)",
		Number:   n,
		Sequence: []*big.Int{val},
		Result:   val,
	}, nil
}

// calculateThetaSeriesD4Lattice computes a(n) for A004011.
// a(0) = 1; if n > 0, a(n) = 24 * (sum of odd divisors of n).
func calculateThetaSeriesD4Lattice(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	// Sum of odd divisors
	sumOdd := calculateSumOddDivisors(n)

	// result = 24 * sumOdd
	result := new(big.Int).Mul(big.NewInt(24), sumOdd)
	return result
}
