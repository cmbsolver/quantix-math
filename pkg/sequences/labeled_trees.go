package sequences

import (
	"fmt"
	"math/big"
)

// Labeled Trees (OEIS A000272)
// URL: https://oeis.org/A000272
// a(n) = n^(n-2) for n >= 1, a(0) = 1.
// This sequence represents the number of trees on n labeled nodes.

// GetLabeledTreesSequence returns the number of labeled trees on n nodes (OEIS A000272).
func GetLabeledTreesSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetLabeledTreesAtPosition(maxNumber)
	}
	return GenerateLabeledTreesSequence(maxNumber)
}

// GenerateLabeledTreesSequence generates the A000272 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateLabeledTreesSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number must be at least 0")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		sequence[i] = calculateA000272(int64(i))
	}

	return &NumericSequence{
		Name:     "Labeled Trees (A000272)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetLabeledTreesAtPosition returns the n-th term of A000272 (n >= 0).
func GetLabeledTreesAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	result := calculateA000272(n.Int64())

	return &NumericSequence{
		Name:     "Labeled Trees (A000272)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateA000272 implements the formula for labeled trees: a(n) = n^(n-2).
// Cayley's formula: For n >= 1, there are n^(n-2) labeled trees on n vertices.
// By convention, a(0) = 1.
func calculateA000272(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	if n == 1 {
		return big.NewInt(1)
	}
	if n == 2 {
		return big.NewInt(1)
	}

	// result = n^(n-2)
	base := big.NewInt(n)
	exp := big.NewInt(n - 2)
	result := new(big.Int).Exp(base, exp, nil)
	return result
}
