package sequences

import (
	"fmt"
	"math/big"
)

// Unlabeled Graphs (OEIS A000088)
// URL: https://oeis.org/A000088
// a(n) is the number of simple graphs on n unlabeled nodes.

// GetUnlabeledGraphsSequence returns the number of simple graphs on n unlabeled nodes (OEIS A000088).
func GetUnlabeledGraphsSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetUnlabeledGraphsAtPosition(maxNumber)
	}
	return GenerateUnlabeledGraphsSequence(maxNumber)
}

// GenerateUnlabeledGraphsSequence generates the A000088 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateUnlabeledGraphsSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number must be at least 0")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		val, err := calculateA000088(int64(i))
		if err != nil {
			return nil, err
		}
		sequence[i] = val
	}

	return &NumericSequence{
		Name:     "Unlabeled Graphs (A000088)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetUnlabeledGraphsAtPosition returns the n-th term of A000088 (n >= 0).
func GetUnlabeledGraphsAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	result, err := calculateA000088(n.Int64())
	if err != nil {
		return nil, err
	}

	return &NumericSequence{
		Name:     "Unlabeled Graphs (A000088)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateA000088 implements the formula for unlabeled graphs using Polya Enumeration Theorem on edges.
// a(n) = 1/n! * sum_{p in S_n} 2^c(p)
// where c(p) is the number of cycles of pairs of nodes induced by permutation p.
// For a permutation p with cycle structure (j_1, j_2, ..., j_n):
// c(p) = sum_{k=1}^n j_k * floor(k/2) + sum_{k=1}^n k * binomial(j_k, 2) + sum_{1<=k<l<=n} gcd(k, l) * j_k * j_l
func calculateA000088(n int64) (*big.Int, error) {
	if n == 0 {
		return big.NewInt(1), nil
	}
	if n == 1 {
		return big.NewInt(1), nil
	}

	total := new(big.Int)
	partitions := generateCyclePartitions(int(n))

	for _, p := range partitions {
		// Calculate the number of permutations with this cycle structure
		// count = n! / (prod (k^j_k * j_k!))
		count := cyclePermutationCount(int(n), p)

		// Calculate the number of cycles of pairs c(p)
		cp := calculatePairCycles(p)

		// Term: count * 2^cp
		term := new(big.Int).Lsh(big.NewInt(1), uint(cp))
		term.Mul(term, count)
		total.Add(total, term)
	}

	// Result = total / n!
	factN := factorial(int(n))
	result := new(big.Int).Div(total, factN)

	return result, nil
}

// calculatePairCycles calculates the number of cycles of edges induced by a permutation of nodes.
func calculatePairCycles(p cyclePartition) int64 {
	var cp int64 = 0

	// sum_{k=1}^n j_k * floor(k/2)
	for k, jk := range p {
		cp += int64(jk * (k / 2))
	}

	// sum_{k=1}^n k * binomial(j_k, 2)
	for k, jk := range p {
		if jk >= 2 {
			cp += int64(k * jk * (jk - 1) / 2)
		}
	}

	// sum_{1<=k<l<=n} gcd(k, l) * j_k * j_l
	// We can iterate over all pairs k, l and only add if k < l
	keys := make([]int, 0, len(p))
	for k := range p {
		keys = append(keys, k)
	}

	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			k := keys[i]
			l := keys[j]
			cp += int64(gcd(k, l) * p[k] * p[l])
		}
	}

	return cp
}
