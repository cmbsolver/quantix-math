package sequences

import (
	"fmt"
	"math/big"
)

// Kendall-Mann numbers (OEIS A000140)
// URL: https://oeis.org/A000140
// Kendall-Mann numbers: the most common number of inversions in a permutation on n letters is floor(n(n-1)/4); a(n) is the number of permutations with this many inversions.
// Row maxima of A008302: coefficients in expansion of Product_{i=0..n-1} (1 + x + ... + x^i).

// GetKendallMannSequence returns the A000140 sequence.
func GetKendallMannSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetKendallMannAtPosition(maxNumber)
	}
	return GenerateKendallMannSequence(maxNumber)
}

// GenerateKendallMannSequence generates the A000140 sequence up to maxNumber.
func GenerateKendallMannSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := int(maxNumber.Int64())
	sequence := make([]*big.Int, limit)

	for n := 1; n <= limit; n++ {
		val := calculateKendallMann(n)
		sequence[n-1] = val
	}

	var result *big.Int
	if limit > 0 {
		result = sequence[limit-1]
	} else {
		result = big.NewInt(0)
	}

	return &NumericSequence{
		Name:     "Kendall-Mann numbers (A000140)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetKendallMannAtPosition returns the n-th term of A000140.
func GetKendallMannAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() <= 0 {
		return nil, fmt.Errorf("position must be positive")
	}

	val := calculateKendallMann(int(n.Int64()))

	return &NumericSequence{
		Name:     "Kendall-Mann numbers (A000140)",
		Number:   n,
		Sequence: []*big.Int{val},
		Result:   val,
	}, nil
}

// calculateKendallMann calculates the n-th term of A000140.
// a(n) is the coefficient of x^k in Product_{i=1..n} (1 + x + ... + x^{i-1})
// where k = floor(n(n-1)/4).
func calculateKendallMann(n int) *big.Int {
	if n <= 0 {
		return big.NewInt(0)
	}
	if n == 1 {
		return big.NewInt(1)
	}

	k := n * (n - 1) / 4
	// dp[j] will store the coefficient of x^j
	dp := make([]*big.Int, k+1)
	for i := range dp {
		dp[i] = big.NewInt(0)
	}
	dp[0] = big.NewInt(1)

	// Multiply by (1 + x + ... + x^{i-1}) for i = 1 to n
	// i=1 is (1), which doesn't change dp[0]=1.
	for i := 2; i <= n; i++ {
		newDp := make([]*big.Int, k+1)
		for j := range newDp {
			newDp[j] = big.NewInt(0)
		}
		// Multiply current dp by (1 + x + ... + x^{i-1})
		// newDp[j] = sum_{m=0}^{min(j, i-1)} dp[j-m]
		for j := 0; j <= k; j++ {
			for m := 0; m < i && j-m >= 0; m++ {
				newDp[j].Add(newDp[j], dp[j-m])
			}
		}
		dp = newDp
	}

	return dp[k]
}
