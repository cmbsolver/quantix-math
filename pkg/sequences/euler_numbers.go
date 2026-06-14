package sequences

import (
	"fmt"
	"math/big"
)

// Euler numbers (A000364)
// URL: https://oeis.org/A000364
// Description: Euler (or secant or "Zig") numbers: e.g.f. (even powers only) sec(x) = 1/cos(x).
// Recurrence: E_n = sum_{k=0}^{n-1} (-1)^{n-k+1} * binom(2n, 2k) * E_k, with E_0 = 1.

// GetEulerNumbersSequence returns the sequence of Euler numbers.
func GetEulerNumbersSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetEulerNumberAtPosition(maxNumber)
	}
	return GenerateEulerNumbersSequence(maxNumber)
}

// GenerateEulerNumbersSequence generates Euler numbers up to maxNumber.
func GenerateEulerNumbersSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Euler numbers (A000364)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   nil,
		}, nil
	}

	var sequence []*big.Int
	memo := make([]*big.Int, 0)

	for n := 0; ; n++ {
		en := calculateEulerNumber(n, memo)
		if en.Cmp(maxNumber) > 0 {
			break
		}
		sequence = append(sequence, en)
		memo = append(memo, en)
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Euler numbers (A000364)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetEulerNumberAtPosition returns the n-th Euler number (0-indexed).
func GetEulerNumberAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be a non-negative integer")
	}

	if !n.IsInt64() {
		return nil, fmt.Errorf("position %s is too large", n.String())
	}

	pos := int(n.Int64())
	memo := make([]*big.Int, 0, pos+1)
	for i := 0; i <= pos; i++ {
		memo = append(memo, calculateEulerNumber(i, memo))
	}

	result := memo[pos]

	return &NumericSequence{
		Name:     "Euler numbers (A000364)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateEulerNumber computes the n-th Euler number using memoized previous values.
func calculateEulerNumber(n int, memo []*big.Int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	if n < len(memo) {
		return memo[n]
	}

	// E_n = sum_{k=0}^{n-1} (-1)^{n-k+1} * binom(2n, 2k) * E_k
	sum := new(big.Int)
	for k := 0; k < n; k++ {
		term := new(big.Int).Binomial(int64(2*n), int64(2*k))
		term.Mul(term, memo[k])

		if (n-k+1)%2 != 0 {
			sum.Sub(sum, term)
		} else {
			sum.Add(sum, term)
		}
	}

	// We want the absolute value for A000364
	return new(big.Int).Abs(sum)
}
