package sequences

import (
	"fmt"
	"math/big"
)

// Labeled Rooted Trees (OEIS A000169)
// URL: https://oeis.org/A000169
// Number of labeled rooted trees with n nodes: n^(n-1).

// GetLabeledRootedTreesSequence returns the number of labeled rooted trees with n nodes (OEIS A000169).
func GetLabeledRootedTreesSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetLabeledRootedTreesAtPosition(maxNumber)
	}
	return GenerateLabeledRootedTreesSequence(maxNumber)
}

// GenerateLabeledRootedTreesSequence generates the A000169 sequence up to maxNumber.
// It returns a(1), a(2), ..., a(maxNumber).
func GenerateLabeledRootedTreesSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("max number must be at least 1 for this sequence")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n)

	for i := 1; i <= n; i++ {
		sequence[i-1] = calculateLabeledRootedTrees(int64(i))
	}

	return &NumericSequence{
		Name:     "Labeled Rooted Trees (A000169)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n-1],
	}, nil
}

// GetLabeledRootedTreesAtPosition returns the n-th term of A000169 (n >= 1).
func GetLabeledRootedTreesAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("position must be at least 1")
	}

	result := calculateLabeledRootedTrees(n.Int64())

	return &NumericSequence{
		Name:     "Labeled Rooted Trees (A000169)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateLabeledRootedTrees calculates n^(n-1).
func calculateLabeledRootedTrees(n int64) *big.Int {
	if n <= 0 {
		return big.NewInt(0)
	}
	if n == 1 {
		return big.NewInt(1)
	}

	bn := big.NewInt(n)
	exponent := big.NewInt(n - 1)

	result := new(big.Int).Exp(bn, exponent, nil)
	return result
}
