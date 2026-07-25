package sequences

import (
	"fmt"
	"math/big"
)

// A000008: Number of ways of making change for n cents using coins of 1, 2, 5, 10 cents.
// URL: https://oeis.org/A000008
// Number of partitions of n into parts 1, 2, 5, and 10.

// GetA000008Sequence returns the number of ways of making change for n cents using coins of 1, 2, 5, 10 cents (OEIS A000008).
// It can return either the entire sequence up to maxNumber or just the value at maxNumber position.
func GetA000008Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000008AtPosition(maxNumber)
	}
	return GenerateA000008Sequence(maxNumber)
}

// GenerateA000008Sequence generates the A000008 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateA000008Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	coins := []int{1, 2, 5, 10}

	dp := make([]*big.Int, n+1)
	for i := range dp {
		dp[i] = big.NewInt(0)
	}
	dp[0] = big.NewInt(1)

	for _, coin := range coins {
		for i := coin; i <= n; i++ {
			dp[i].Add(dp[i], dp[i-coin])
		}
	}

	sequence := make([]*big.Int, n+1)
	for i := 0; i <= n; i++ {
		sequence[i] = new(big.Int).Set(dp[i])
	}

	return &NumericSequence{
		Name:     "Ways to Make Change (A000008)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   dp[n],
	}, nil
}

// GetA000008AtPosition returns the n-th term of A000008.
func GetA000008AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	pos := int(n.Int64())
	coins := []int{1, 2, 5, 10}

	dp := make([]*big.Int, pos+1)
	for i := range dp {
		dp[i] = big.NewInt(0)
	}
	dp[0] = big.NewInt(1)

	for _, coin := range coins {
		for i := coin; i <= pos; i++ {
			dp[i].Add(dp[i], dp[i-coin])
		}
	}

	result := new(big.Int).Set(dp[pos])

	return &NumericSequence{
		Name:     "Ways to Make Change (A000008)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
