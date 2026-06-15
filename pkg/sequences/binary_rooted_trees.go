package sequences

import (
	"fmt"
	"math/big"
)

// BinaryRootedTrees A002572: Number of partitions of 1 into n powers of 1/2; or the number of binary rooted trees.
// URL: https://oeis.org/A002572
// Recurrence: v(c, d) = if d < 0 or c < 0 then 0 elif d = c then 1 else sum(v(i, d-c), i=1..2*c)
// a(n) = v(1, n-1) with a(1) = 1.

// GetBinaryRootedTreesA002572Sequence returns the A002572 sequence.
func GetBinaryRootedTreesA002572Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetBinaryRootedTreesA002572AtPosition(maxNumber)
	}
	return GenerateBinaryRootedTreesA002572Sequence(maxNumber)
}

// GenerateBinaryRootedTreesA002572Sequence generates the A002572 sequence up to maxNumber.
func GenerateBinaryRootedTreesA002572Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	if n == 0 {
		return &NumericSequence{
			Name:     "Binary rooted trees (A002572)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   big.NewInt(0),
		}, nil
	}

	results := make([]*big.Int, n)
	memo := make(map[string]*big.Int)

	for i := 1; i <= n; i++ {
		if i == 1 {
			results[i-1] = big.NewInt(1)
		} else {
			results[i-1] = computeV(1, i-1, memo)
		}
	}

	return &NumericSequence{
		Name:     "Binary rooted trees (A002572)",
		Number:   maxNumber,
		Sequence: results,
		Result:   results[n-1],
	}, nil
}

// GetBinaryRootedTreesA002572AtPosition returns the n-th term of A002572.
func GetBinaryRootedTreesA002572AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() <= 0 {
		return nil, fmt.Errorf("position must be positive")
	}

	pos := int(n.Int64())
	var result *big.Int
	if pos == 1 {
		result = big.NewInt(1)
	} else {
		memo := make(map[string]*big.Int)
		result = computeV(1, pos-1, memo)
	}

	return &NumericSequence{
		Name:     "Binary rooted trees (A002572)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// computeV implements the recurrence: v(c, d) = if d == c then 1 else sum(v(i, d-c), i=1..2*c)
func computeV(c, d int, memo map[string]*big.Int) *big.Int {
	if d < 0 || c < 0 {
		return big.NewInt(0)
	}
	if d == c {
		return big.NewInt(1)
	}
	if d < c {
		return big.NewInt(0)
	}

	key := fmt.Sprintf("%d,%d", c, d)
	if val, ok := memo[key]; ok {
		return val
	}

	sum := big.NewInt(0)
	for i := 1; i <= 2*c; i++ {
		sum.Add(sum, computeV(i, d-c, memo))
	}

	memo[key] = sum
	return sum
}
