package sequences

import (
	"fmt"
	"math/big"
)

// Partitions into distinct parts (OEIS A000009)
// URL: https://oeis.org/A000009
// a(n) is the number of partitions of n into distinct parts.
// It is also the number of partitions of n into odd parts.

// GetPartitionsDistinctSequence returns the A000009 sequence.
func GetPartitionsDistinctSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetPartitionsDistinctAtPosition(maxNumber)
	}
	return GeneratePartitionsDistinctSequence(maxNumber)
}

// GeneratePartitionsDistinctSequence generates the A000009 sequence up to maxNumber.
func GeneratePartitionsDistinctSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	// Using dynamic programming to count partitions into distinct parts.
	// dp[i] = number of partitions of i into distinct parts.
	dp := make([]*big.Int, n+1)
	for i := range dp {
		dp[i] = big.NewInt(0)
	}
	dp[0] = big.NewInt(1)

	for i := 1; i <= n; i++ {
		for j := n; j >= i; j-- {
			dp[j].Add(dp[j], dp[j-i])
		}
	}

	for i := 0; i <= n; i++ {
		sequence[i] = new(big.Int).Set(dp[i])
	}

	return &NumericSequence{
		Name:     "Partitions into distinct parts (A000009)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetPartitionsDistinctAtPosition returns the n-th term of A000009.
func GetPartitionsDistinctAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	target := int(n.Int64())
	dp := make([]*big.Int, target+1)
	for i := range dp {
		dp[i] = big.NewInt(0)
	}
	dp[0] = big.NewInt(1)

	for i := 1; i <= target; i++ {
		for j := target; j >= i; j-- {
			dp[j].Add(dp[j], dp[j-i])
		}
	}

	result := dp[target]

	return &NumericSequence{
		Name:     "Partitions into distinct parts (A000009)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
