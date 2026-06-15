package sequences

import (
	"fmt"
	"math/big"
)

// Rooted unlabeled trees with n nodes (OEIS A000081)
// URL: https://oeis.org/A000081
// a(n) is the number of rooted unlabeled trees with n nodes.
// Recurrence: a(n+1) = (1/n) * sum_{k=1}^n (sum_{d|k} d * a(d)) * a(n-k+1) for n >= 1, a(1) = 1.

// GetRootedUnlabeledTreesSequence returns the number of rooted unlabeled trees with n nodes (OEIS A000081).
func GetRootedUnlabeledTreesSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetRootedUnlabeledTreesAtPosition(maxNumber)
	}
	return GenerateRootedUnlabeledTreesSequence(maxNumber)
}

// GenerateRootedUnlabeledTreesSequence generates the A000081 sequence up to maxNumber.
func GenerateRootedUnlabeledTreesSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number must be at least 0")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	if n >= 0 {
		sequence[0] = big.NewInt(0)
	}
	if n >= 1 {
		sequence[1] = big.NewInt(1)
	}

	if n > 1 {
		// precompute sums of divisors terms
		// s(k) = sum_{d|k} d * a(d)
		s := make([]*big.Int, n) // we need s(1) to s(n-1)

		for i := 2; i <= n; i++ {
			// calculate a(i) using values a(1)...a(i-1)
			// a(i) = (1/(i-1)) * sum_{k=1}^{i-1} s(k) * a(i-k)

			// first update s(i-1)
			idx := i - 1
			s[idx] = new(big.Int)
			for d := 1; d <= idx; d++ {
				if idx%d == 0 {
					term := new(big.Int).Mul(big.NewInt(int64(d)), sequence[d])
					s[idx].Add(s[idx], term)
				}
			}

			sum := new(big.Int)
			for k := 1; k <= idx; k++ {
				term := new(big.Int).Mul(s[k], sequence[idx-k+1])
				sum.Add(sum, term)
			}
			sequence[i] = new(big.Int).Div(sum, big.NewInt(int64(idx)))
		}
	}

	return &NumericSequence{
		Name:     "Rooted unlabeled trees (A000081)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetRootedUnlabeledTreesAtPosition returns the n-th term of A000081.
func GetRootedUnlabeledTreesAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	// For A000081, we need all previous terms to calculate the n-th term.
	seq, err := GenerateRootedUnlabeledTreesSequence(n)
	if err != nil {
		return nil, err
	}

	return &NumericSequence{
		Name:     "Rooted unlabeled trees (A000081)",
		Number:   n,
		Sequence: []*big.Int{seq.Result},
		Result:   seq.Result,
	}, nil
}
