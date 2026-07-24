package sequences

import (
	"fmt"
	"math"
	"math/big"
)

// Smallest prime power >= n.
// URL: https://oeis.org/A000015
// Description: a(n) is the smallest prime power q >= n.

// GetSmallestPrimePowerA000015Sequence returns the A000015 sequence.
func GetSmallestPrimePowerA000015Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetSmallestPrimePowerA000015AtPosition(maxNumber)
	}
	return GenerateSmallestPrimePowerA000015Sequence(maxNumber)
}

// GenerateSmallestPrimePowerA000015Sequence generates the A000015 sequence up to maxNumber.
func GenerateSmallestPrimePowerA000015Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() <= 0 {
		return &NumericSequence{
			Name:     "Smallest prime power >= n (A000015)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}

	limit := maxNumber.Int64()
	if maxNumber.IsInt64() && limit > 0 {
		var sequence []*big.Int
		for n := int64(1); n <= limit; n++ {
			sequence = append(sequence, big.NewInt(SmallestPrimePowerGE(n)))
		}
		var result *big.Int
		if len(sequence) > 0 {
			result = sequence[len(sequence)-1]
		}
		return &NumericSequence{
			Name:     "Smallest prime power >= n (A000015)",
			Number:   maxNumber,
			Sequence: sequence,
			Result:   result,
		}, nil
	}

	return nil, fmt.Errorf("maxNumber too large for current implementation")
}

// GetSmallestPrimePowerA000015AtPosition returns the n-th term of A000015.
func GetSmallestPrimePowerA000015AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() <= 0 {
		return nil, fmt.Errorf("position must be positive")
	}

	if n.IsInt64() {
		val := big.NewInt(SmallestPrimePowerGE(n.Int64()))
		return &NumericSequence{
			Name:     "Smallest prime power >= n (A000015)",
			Number:   n,
			Sequence: []*big.Int{val},
			Result:   val,
		}, nil
	}

	return nil, fmt.Errorf("position too large for current implementation")
}

// SmallestPrimePowerGE finds the smallest prime power q >= n.
func SmallestPrimePowerGE(n int64) int64 {
	if n <= 1 {
		return 1
	}
	for q := n; ; q++ {
		if IsPrimePower(q) {
			return q
		}
	}
}

// IsPrimePower checks if n is a prime power p^k, k >= 1.
func IsPrimePower(n int64) bool {
	if n < 2 {
		return false
	}
	// A000961: Prime powers: 1 is not included, but for A000015(1)=1.
	// Actually A000015(1) = 1, and 1 is often considered a prime power p^0 or just handled separately.
	// Looking at A000015 data: 1, 2, 3, 4, 5, 7, 7, 8, 9, 11...
	// So 1 is included.

	// Check if n is prime
	bn := big.NewInt(n)
	if bn.ProbablyPrime(20) {
		return true
	}

	// Check for p^k, k > 1
	// n = p^k => k = log_p(n). Max k is log_2(n).
	for k := 2; (1 << uint(k)) <= n; k++ {
		// Calculate k-th root of n
		root := RoundRoot(n, k)
		if root < 2 {
			continue
		}
		// Check if root^k == n
		if Power(root, k) == n {
			// Check if root is prime
			if big.NewInt(root).ProbablyPrime(20) {
				return true
			}
		}
	}
	return false
}

// RoundRoot calculates floor(n^(1/k)).
func RoundRoot(n int64, k int) int64 {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 0
	}
	if k == 1 {
		return n
	}
	// Use math.Pow for a floating point approximation and then refine
	root := int64(math.Pow(float64(n), 1.0/float64(k)))
	// Refine to find floor
	if Power(root+1, k) <= n {
		root++
	} else if Power(root, k) > n {
		root--
	}
	return root
}

// Power calculates a^b for int64.
func Power(a int64, b int) int64 {
	res := int64(1)
	for i := 0; i < b; i++ {
		res *= a
	}
	return res
}
