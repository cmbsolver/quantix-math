package sequences

import (
	"fmt"
	"math/big"
)

// Number of Divisors Sequence (OEIS A000005)
// URL: https://oeis.org/A000005
// d(n) (also called tau(n) or sigma_0(n)), the number of divisors of n.

// GetDivisorCountSequence returns the number of divisors of n (OEIS A000005).
func GetDivisorCountSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetDivisorCountAtPosition(maxNumber)
	}
	return GenerateDivisorCountSequence(maxNumber)
}

// GenerateDivisorCountSequence generates the A000005 sequence up to maxNumber.
// It returns a(1), a(2), ..., a(maxNumber).
func GenerateDivisorCountSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("max number must be at least 1 for this sequence")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n)

	for i := 1; i <= n; i++ {
		count := countDivisors(int64(i))
		sequence[i-1] = big.NewInt(count)
	}

	return &NumericSequence{
		Name:     "Number of Divisors (A000005)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n-1],
	}, nil
}

// GetDivisorCountAtPosition returns the n-th term of A000005 (n >= 1).
func GetDivisorCountAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("position must be at least 1")
	}

	val := n.Int64()
	result := big.NewInt(countDivisors(val))

	return &NumericSequence{
		Name:     "Number of Divisors (A000005)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// countDivisors calculates the number of divisors of n using prime factorization.
// d(n) = Product (e_i + 1) where n = Product p_i^e_i
func countDivisors(n int64) int64 {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	var count int64 = 1
	d := n
	for i := int64(2); i*i <= d; i++ {
		if d%i == 0 {
			var exponent int64 = 0
			for d%i == 0 {
				exponent++
				d /= i
			}
			count *= (exponent + 1)
		}
	}
	if d > 1 {
		count *= 2
	}
	return count
}
