package sequences

import (
	"fmt"
	"math/big"
)

// Partitions into 2 squares (OEIS A000161)
// URL: https://oeis.org/A000161
// Number of partitions of n into 2 squares.
// Also described as: Number of ways of writing n as a sum of 2 (possibly zero) squares when order does not matter.

// GetPartitionsInto2SquaresSequence returns the number of partitions of n into 2 squares (OEIS A000161).
func GetPartitionsInto2SquaresSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetPartitionsInto2SquaresAtPosition(maxNumber)
	}
	return GeneratePartitionsInto2SquaresSequence(maxNumber)
}

// GeneratePartitionsInto2SquaresSequence generates the A000161 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GeneratePartitionsInto2SquaresSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number must be at least 0")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		count := countPartitionsInto2Squares(int64(i))
		sequence[i] = big.NewInt(count)
	}

	return &NumericSequence{
		Name:     "Partitions into 2 squares (A000161)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetPartitionsInto2SquaresAtPosition returns the n-th term of A000161 (n >= 0).
func GetPartitionsInto2SquaresAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	val := n.Int64()
	result := big.NewInt(countPartitionsInto2Squares(val))

	return &NumericSequence{
		Name:     "Partitions into 2 squares (A000161)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// countPartitionsInto2Squares calculates the number of ways to write n as a sum of 2 squares.
// x^2 + y^2 = n, where 0 <= x <= y.
func countPartitionsInto2Squares(n int64) int64 {
	if n < 0 {
		return 0
	}
	var count int64 = 0
	for x := int64(0); x*x <= n/2; x++ {
		y2 := n - x*x
		// Check if y2 is a perfect square
		if y2 >= 0 {
			root := int64(0)
			// Simple sqrt for smallish n
			// For larger n we might want something more efficient,
			// but A000161 typically isn't called for huge n in this UI context.
			for i := int64(0); i*i <= y2; i++ {
				if i*i == y2 {
					root = i
					break
				}
			}
			if root*root == y2 {
				// We found a pair (x, y) such that x^2 + y^2 = n and x <= y.
				// Since x*x <= n/2, y*y must be >= n/2, so x <= y is guaranteed if y exists.
				count++
			}
		}
	}
	return count
}
