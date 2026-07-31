package sequences

import (
	"fmt"
	"math/big"
)

// A010025: Crystal ball sequence for squashed {D_5}^* lattice, perhaps the smallest example of a "non-superficial" lattice.
// URL: https://oeis.org/A010025
// Description: a(n) = (8*n^5 + 15*n^4 + 30*n^3 + 45*n^2 + 46*n + 12) / 12 for n >= 0.

// GetA010025Sequence returns the A010025 sequence up to maxNumber terms or the n-th term.
func GetA010025Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA010025AtPosition(maxNumber)
	}

	return GenerateA010025Sequence(maxNumber)
}

// GenerateA010025Sequence generates the A010025 sequence up to maxNumber terms.
func GenerateA010025Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n)

	for i := 0; i < n; i++ {
		val := CalculateA010025(i)
		sequence[i] = new(big.Int).SetUint64(val)
	}

	result := big.NewInt(0)
	if n > 0 {
		result = sequence[n-1]
	}

	return &NumericSequence{
		Name:     "Crystal ball sequence for squashed {D_5}^* lattice (A010025)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetA010025AtPosition returns the n-th term of A010025.
func GetA010025AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	val := CalculateA010025(int(n.Int64()))

	return &NumericSequence{
		Name:     "Crystal ball sequence for squashed {D_5}^* lattice (A010025)",
		Number:   n,
		Sequence: []*big.Int{new(big.Int).SetUint64(val)},
		Result:   new(big.Int).SetUint64(val),
	}, nil
}

// CalculateA010025 calculates the n-th term of A010025.
func CalculateA010025(n int) uint64 {
	if n < 0 {
		return 0
	}

	n64 := uint64(n)
	n2 := n64 * n64
	n3 := n2 * n64
	n4 := n3 * n64
	n5 := n4 * n64

	numerator := 8*n5 + 15*n4 + 30*n3 + 45*n2 + 46*n64 + 12

	return numerator / 12
}
