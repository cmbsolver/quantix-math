package sequences

import (
	"fmt"
	"math/big"
)

// Unlabeled Directed Graphs (OEIS A000273)
// URL: https://oeis.org/A000273
// a(n) is the number of unlabeled directed graphs with n nodes.
// The number of simple directed graphs on n nodes.

// GetUnlabeledDigraphsSequence returns the number of unlabeled directed graphs on n nodes (OEIS A000273).
func GetUnlabeledDigraphsSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetUnlabeledDigraphsAtPosition(maxNumber)
	}
	return GenerateUnlabeledDigraphsSequence(maxNumber)
}

// GenerateUnlabeledDigraphsSequence generates the A000273 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateUnlabeledDigraphsSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number must be at least 0")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		val, err := calculateA000273(int64(i))
		if err != nil {
			return nil, err
		}
		sequence[i] = val
	}

	return &NumericSequence{
		Name:     "Unlabeled Directed Graphs (A000273)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetUnlabeledDigraphsAtPosition returns the n-th term of A000273 (n >= 0).
func GetUnlabeledDigraphsAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	result, err := calculateA000273(n.Int64())
	if err != nil {
		return nil, err
	}

	return &NumericSequence{
		Name:     "Unlabeled Directed Graphs (A000273)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateA000273 implements the formula for unlabeled directed graphs using Polya Enumeration Theorem.
// a(n) = 1/n! * sum_{p in S_n} 2^c(p)
// where c(p) is the number of cycles of directed edges induced by permutation p.
// For a permutation with cycle structure (j_1, j_2, ..., j_n), where j_k is the number of cycles of length k:
// c(p) = sum_{k=1}^n sum_{l=1}^n gcd(k, l) * j_k * j_l - sum_{k=1}^n j_k
func calculateA000273(n int64) (*big.Int, error) {
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

		// Calculate the number of cycles of edges c(p)
		// c(p) = sum_{k,l} gcd(k,l) * j_k * j_l - sum_k j_k
		cp := calculateEdgeCycles(p)

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

// cyclePartition represents a partition of n as a map where key is cycle length and value is count of cycles.
type cyclePartition map[int]int

func generateCyclePartitions(n int) []cyclePartition {
	var result []cyclePartition
	var p cyclePartition = make(map[int]int)
	findCyclePartitions(n, n, p, &result)
	return result
}

func findCyclePartitions(n, max int, current cyclePartition, result *[]cyclePartition) {
	if n == 0 {
		pCopy := make(cyclePartition)
		for k, v := range current {
			pCopy[k] = v
		}
		*result = append(*result, pCopy)
		return
	}

	for i := min(n, max); i >= 1; i-- {
		current[i]++
		findCyclePartitions(n-i, i, current, result)
		current[i]--
		if current[i] == 0 {
			delete(current, i)
		}
	}
}

func cyclePermutationCount(n int, p cyclePartition) *big.Int {
	res := factorial(n)
	for k, jk := range p {
		// k^jk
		powK := new(big.Int).Exp(big.NewInt(int64(k)), big.NewInt(int64(jk)), nil)
		res.Div(res, powK)
		// jk!
		res.Div(res, factorial(jk))
	}
	return res
}

func calculateEdgeCycles(p cyclePartition) int64 {
	var cp int64 = 0
	// sum_{k,l} gcd(k,l) * j_k * j_l
	for k, jk := range p {
		for l, jl := range p {
			cp += int64(gcd(k, l) * jk * jl)
		}
	}
	// - sum_k j_k
	for _, jk := range p {
		cp -= int64(jk)
	}
	return cp
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func factorial(n int) *big.Int {
	res := big.NewInt(1)
	for i := 2; i <= n; i++ {
		res.Mul(res, big.NewInt(int64(i)))
	}
	return res
}
