package sequences

import (
	"fmt"
	"math/big"
)

// Partitions of n (OEIS A000041)
// URL: https://oeis.org/A000041
// a(n) is the number of partitions of n (the partition numbers).

// GetPartitionsSequence returns the A000041 sequence.
func GetPartitionsSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetPartitionsAtPosition(maxNumber)
	}
	return GeneratePartitionsSequence(maxNumber)
}

// GeneratePartitionsSequence generates the A000041 sequence up to maxNumber.
func GeneratePartitionsSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	// Using dynamic programming to count partitions of n.
	// dp[i] = number of partitions of i.
	dp := make([]*big.Int, n+1)
	for i := range dp {
		dp[i] = big.NewInt(0)
	}
	dp[0] = big.NewInt(1)

	for part := 1; part <= n; part++ {
		for i := part; i <= n; i++ {
			dp[i].Add(dp[i], dp[i-part])
		}
	}

	for i := 0; i <= n; i++ {
		sequence[i] = new(big.Int).Set(dp[i])
	}

	return &NumericSequence{
		Name:     "Partitions of n (A000041)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetPartitionsAtPosition returns the n-th term of A000041.
func GetPartitionsAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	target := int(n.Int64())
	dp := make([]*big.Int, target+1)
	for i := range dp {
		dp[i] = big.NewInt(0)
	}
	dp[0] = big.NewInt(1)

	for part := 1; part <= target; part++ {
		for i := part; i <= target; i++ {
			dp[i].Add(dp[i], dp[i-part])
		}
	}

	result := dp[target]

	return &NumericSequence{
		Name:     "Partitions of n (A000041)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
