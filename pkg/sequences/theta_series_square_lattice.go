package sequences

import (
	"fmt"
	"math/big"
)

// Theta series of square lattice (OEIS A004018)
// URL: https://oeis.org/A004018
// Number of ways of writing n as a sum of 2 squares.
// Often denoted by r(n) or r_2(n).
// a(n) is the number of pairs (x, y) such that x^2 + y^2 = n, where x and y can be any integers (positive, negative, or zero).

// GetThetaSeriesSquareLatticeSequence returns the Theta series of square lattice (OEIS A004018).
func GetThetaSeriesSquareLatticeSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetThetaSeriesSquareLatticeAtPosition(maxNumber)
	}
	return GenerateThetaSeriesSquareLatticeSequence(maxNumber)
}

// GenerateThetaSeriesSquareLatticeSequence generates the A004018 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateThetaSeriesSquareLatticeSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number must be at least 0")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		count := countWaysToSumOf2Squares(int64(i))
		sequence[i] = big.NewInt(count)
	}

	return &NumericSequence{
		Name:     "Theta series of square lattice (A004018)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetThetaSeriesSquareLatticeAtPosition returns the n-th term of A004018 (n >= 0).
func GetThetaSeriesSquareLatticeAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	val := n.Int64()
	result := big.NewInt(countWaysToSumOf2Squares(val))

	return &NumericSequence{
		Name:     "Theta series of square lattice (A004018)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// countWaysToSumOf2Squares calculates the number of ways to write n as a sum of 2 squares.
// x^2 + y^2 = n, where x, y are any integers.
func countWaysToSumOf2Squares(n int64) int64 {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1 // Only (0,0)
	}
	var count int64 = 0
	for x := int64(0); x*x <= n; x++ {
		y2 := n - x*x
		if y2 >= 0 {
			root := int64(0)
			// Simple sqrt for smallish n
			for i := int64(0); i*i <= y2; i++ {
				if i*i == y2 {
					root = i
					break
				}
			}
			if root*root == y2 {
				// We found a pair (x, y) such that x^2 + y^2 = n
				// Cases for x and y being zero or non-zero to account for signs

				if x == 0 && root == 0 {
					// (0,0) - already handled above but kept for completeness
					count += 1
				} else if x == 0 || root == 0 {
					// (0, y), (0, -y), (x, 0), (-x, 0)
					// Since we only iterate x >= 0, we found (0, root) or (x, 0)
					// If x=0, root must be > 0. We have (0, root) and (0, -root).
					// If root=0, x must be > 0. We have (x, 0) and (-x, 0).
					count += 2
				} else {
					// (x, root), (x, -root), (-x, root), (-x, -root)
					// Both x and root are > 0.
					count += 4
				}
			}
		}
	}
	return count
}
