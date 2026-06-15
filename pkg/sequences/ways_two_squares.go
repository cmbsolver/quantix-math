package sequences

import (
	"fmt"
	"math/big"
)

// Number of ways of writing n as a sum of at most two nonzero squares, where order matters (A002654)
// URL: https://oeis.org/A002654
// Description: a(n) = (number of divisors of n of form 4k+1) - (number of divisors of n of form 4k+3).

// GetWaysTwoSquaresSequence returns the A002654 sequence.
func GetWaysTwoSquaresSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetWaysTwoSquaresAtPosition(maxNumber)
	}
	return GenerateWaysTwoSquaresSequence(maxNumber)
}

// GenerateWaysTwoSquaresSequence generates the A002654 sequence up to maxNumber.
func GenerateWaysTwoSquaresSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() <= 0 {
		return &NumericSequence{
			Name:     "Ways as sum of at most two nonzero squares (A002654)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}

	limit := int(maxNumber.Int64())
	if !maxNumber.IsInt64() {
		return nil, fmt.Errorf("max number too large")
	}

	sequence := make([]*big.Int, limit)
	for i := 1; i <= limit; i++ {
		sequence[i-1] = calculateWaysTwoSquares(int64(i))
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Ways as sum of at most two nonzero squares (A002654)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetWaysTwoSquaresAtPosition returns the n-th term of A002654.
func GetWaysTwoSquaresAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() <= 0 {
		return nil, fmt.Errorf("position must be positive")
	}

	if !n.IsInt64() {
		return nil, fmt.Errorf("position too large")
	}

	val := calculateWaysTwoSquares(n.Int64())

	return &NumericSequence{
		Name:     "Ways as sum of at most two nonzero squares (A002654)",
		Number:   n,
		Sequence: []*big.Int{val},
		Result:   val,
	}, nil
}

// calculateWaysTwoSquares computes a(n) = count(d|n, d=1 mod 4) - count(d|n, d=3 mod 4).
func calculateWaysTwoSquares(n int64) *big.Int {
	count1 := 0
	count3 := 0
	for i := int64(1); i*i <= n; i++ {
		if n%i == 0 {
			// i is a divisor
			if i%4 == 1 {
				count1++
			} else if i%4 == 3 {
				count3++
			}

			j := n / i
			if j != i {
				// j is also a divisor
				if j%4 == 1 {
					count1++
				} else if j%4 == 3 {
					count3++
				}
			}
		}
	}
	return big.NewInt(int64(count1 - count3))
}
