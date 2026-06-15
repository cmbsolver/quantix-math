package sequences

import (
	"fmt"
	"math/big"
)

// Sum of Divisors Sequence (OEIS A000203)
// URL: https://oeis.org/A000203
// a(n) = sigma(n), the sum of the divisors of n. Also called sigma_1(n).

// GetSumDivisorsSequence returns the sum of divisors of n (OEIS A000203).
func GetSumDivisorsSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetSumDivisorsAtPosition(maxNumber)
	}
	return GenerateSumDivisorsSequence(maxNumber)
}

// GenerateSumDivisorsSequence generates the A000203 sequence up to maxNumber.
// It returns a(1), a(2), ..., a(maxNumber).
func GenerateSumDivisorsSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("max number must be at least 1 for this sequence")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n)

	for i := 1; i <= n; i++ {
		sum := calculateSumDivisors(int64(i))
		sequence[i-1] = big.NewInt(sum)
	}

	return &NumericSequence{
		Name:     "Sum of Divisors (A000203)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n-1],
	}, nil
}

// GetSumDivisorsAtPosition returns the n-th term of A000203 (n >= 1).
func GetSumDivisorsAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("position must be at least 1")
	}

	val := n.Int64()
	result := big.NewInt(calculateSumDivisors(val))

	return &NumericSequence{
		Name:     "Sum of Divisors (A000203)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateSumDivisors calculates the sum of divisors of n using prime factorization.
// sigma(n) = Product ( (p_i^(e_i+1) - 1) / (p_i - 1) ) where n = Product p_i^e_i
func calculateSumDivisors(n int64) int64 {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	var sum int64 = 1
	d := n
	for i := int64(2); i*i <= d; i++ {
		if d%i == 0 {
			var pPower int64 = 1
			var pSum int64 = 1
			for d%i == 0 {
				pPower *= i
				pSum += pPower
				d /= i
			}
			sum *= pSum
		}
	}
	if d > 1 {
		sum *= (d + 1)
	}
	return sum
}
