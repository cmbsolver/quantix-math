package sequences

import (
	"fmt"
	"math/big"
)

// Euler zigzag numbers (A000111)
// URL: https://oeis.org/A000111
// Description: Euler or up/down numbers: e.g.f. sec(x) + tan(x). Also for n >= 2, half the number of alternating permutations on n letters (A001250).
// Also known as alternating permutations or zigzag numbers.

// GetEulerZigzagSequence returns the sequence of Euler zigzag numbers.
func GetEulerZigzagSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetEulerZigzagAtPosition(maxNumber)
	}
	return GenerateEulerZigzagSequence(maxNumber)
}

// GenerateEulerZigzagSequence generates Euler zigzag numbers up to maxNumber.
func GenerateEulerZigzagSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Euler zigzag numbers (A000111)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   nil,
		}, nil
	}

	var sequence []*big.Int
	memo := make([]*big.Int, 0)

	for n := 0; ; n++ {
		ez := calculateEulerZigzag(n, memo)
		if ez.Cmp(maxNumber) > 0 {
			break
		}
		sequence = append(sequence, ez)
		memo = append(memo, ez)
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Euler zigzag numbers (A000111)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetEulerZigzagAtPosition returns the n-th Euler zigzag number (0-indexed).
func GetEulerZigzagAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be a non-negative integer")
	}

	if !n.IsInt64() {
		return nil, fmt.Errorf("position %s is too large", n.String())
	}

	pos := int(n.Int64())
	memo := make([]*big.Int, 0, pos+1)
	for i := 0; i <= pos; i++ {
		memo = append(memo, calculateEulerZigzag(i, memo))
	}

	result := memo[pos]

	return &NumericSequence{
		Name:     "Euler zigzag numbers (A000111)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateEulerZigzag computes the n-th Euler zigzag number using memoized previous values.
// The (n+1)th Zigzag number is: a(n+1) = \dfrac{\sum_{k=0}^{n} (\binom{n}{k} a(k) a(n-k))}{2}
// For n=0, a(0)=1. For n=1, a(1)=1.
func calculateEulerZigzag(n int, memo []*big.Int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	if n == 1 {
		return big.NewInt(1)
	}
	if n < len(memo) {
		return memo[n]
	}

	// a(n) = (sum_{k=0}^{n-1} binom(n-1, k) * a(k) * a(n-1-k)) / 2
	nm1 := n - 1
	sum := new(big.Int)
	for k := 0; k <= nm1; k++ {
		term := new(big.Int).Binomial(int64(nm1), int64(k))
		term.Mul(term, memo[k])
		term.Mul(term, memo[nm1-k])
		sum.Add(sum, term)
	}

	return new(big.Int).Div(sum, big.NewInt(2))
}
