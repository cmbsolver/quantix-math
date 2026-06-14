package sequences

import (
	"fmt"
	"math/big"
)

// Ways to make change for n cents using coins of 1, 2, 5, 10 cents (OEIS A000008).
// URL: https://oeis.org/A000008
// a(n) is the number of ways to make change for n using {1, 2, 5, 10}.

// GetWaysToMakeChangeSequence returns the number of ways of making change for n cents using coins of 1, 2, 5, 10 cents (OEIS A000008).
func GetWaysToMakeChangeSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetWaysToMakeChangeAtPosition(maxNumber)
	}
	return GenerateWaysToMakeChangeSequence(maxNumber)
}

// GenerateWaysToMakeChangeSequence generates the A000008 sequence up to maxNumber.
func GenerateWaysToMakeChangeSequence(maxNumber *big.Int) (*NumericSequence, error) {
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

// GetWaysToMakeChangeAtPosition returns the n-th term of A000008.
func GetWaysToMakeChangeAtPosition(n *big.Int) (*NumericSequence, error) {
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
