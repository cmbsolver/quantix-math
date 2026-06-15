package sequences

import (
	"fmt"
	"math/big"
)

// Binary partitions (OEIS A000123)
// URL: https://oeis.org/A000123
// a(n) is the number of partitions of 2n into powers of 2.
// Recurrence: a(n) = a(n-1) + a(floor(n/2)) with a(0) = 1.

// GetBinaryPartitionsSequence returns the A000123 sequence.
func GetBinaryPartitionsSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetBinaryPartitionsAtPosition(maxNumber)
	}
	return GenerateBinaryPartitionsSequence(maxNumber)
}

// GenerateBinaryPartitionsSequence generates the A000123 sequence up to maxNumber.
func GenerateBinaryPartitionsSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	// dp[i] = a(i)
	dp := make([]*big.Int, n+1)
	dp[0] = big.NewInt(1)

	for i := 1; i <= n; i++ {
		dp[i] = new(big.Int).Set(dp[i-1])
		dp[i].Add(dp[i], dp[i/2])
	}

	for i := 0; i <= n; i++ {
		sequence[i] = new(big.Int).Set(dp[i])
	}

	return &NumericSequence{
		Name:     "Binary partitions (A000123)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetBinaryPartitionsAtPosition returns the n-th term of A000123.
func GetBinaryPartitionsAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	target := int(n.Int64())
	dp := make([]*big.Int, target+1)
	dp[0] = big.NewInt(1)

	for i := 1; i <= target; i++ {
		dp[i] = new(big.Int).Set(dp[i-1])
		dp[i].Add(dp[i], dp[i/2])
	}

	result := dp[target]

	return &NumericSequence{
		Name:     "Binary partitions (A000123)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
