package sequences

import (
	"fmt"
	"math/big"
)

// Sum of odd divisors of n (A000593)
// URL: https://oeis.org/A000593
// Description: a(n) = sum_{d|n, d is odd} d.

// GetSumOddDivisorsSequence returns the sum of odd divisors sequence (OEIS A000593).
func GetSumOddDivisorsSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetSumOddDivisorsAtPosition(maxNumber)
	}
	return GenerateSumOddDivisorsSequence(maxNumber)
}

// GenerateSumOddDivisorsSequence generates the A000593 sequence up to maxNumber.
func GenerateSumOddDivisorsSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() <= 0 {
		return &NumericSequence{
			Name:     "Sum of Odd Divisors (A000593)",
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
		sequence[i-1] = calculateSumOddDivisors(int64(i))
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Sum of Odd Divisors (A000593)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetSumOddDivisorsAtPosition returns the n-th term of A000593.
func GetSumOddDivisorsAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() <= 0 {
		return nil, fmt.Errorf("position must be positive")
	}

	if !n.IsInt64() {
		return nil, fmt.Errorf("position too large")
	}

	val := calculateSumOddDivisors(n.Int64())

	return &NumericSequence{
		Name:     "Sum of Odd Divisors (A000593)",
		Number:   n,
		Sequence: []*big.Int{val},
		Result:   val,
	}, nil
}

// calculateSumOddDivisors computes the sum of odd divisors of n.
func calculateSumOddDivisors(n int64) *big.Int {
	sum := big.NewInt(0)
	for i := int64(1); i*i <= n; i++ {
		if n%i == 0 {
			// i is a divisor
			if i%2 != 0 {
				sum.Add(sum, big.NewInt(i))
			}
			j := n / i
			if j != i {
				// j is also a divisor
				if j%2 != 0 {
					sum.Add(sum, big.NewInt(j))
				}
			}
		}
	}
	return sum
}
