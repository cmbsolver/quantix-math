package sequences

import (
	"fmt"
	"math/big"
)

// Sets of Lists (OEIS A000262)
// URL: https://oeis.org/A000262
// a(n) is the number of "sets of lists": number of partitions of {1,...,n} into any number of lists,
// where a list means an ordered subset.
// a(n) = sum_{k=1}^n L(n, k) where L(n, k) is the Lah number L(n, k) = n!(n-1)! / (k!(k-1)!(n-k)!).

// GetSetsOfListsSequence returns the sets of lists sequence (OEIS A000262).
func GetSetsOfListsSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetSetsOfListsAtPosition(maxNumber)
	}
	return GenerateSetsOfListsSequence(maxNumber)
}

// GenerateSetsOfListsSequence generates the A000262 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateSetsOfListsSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number must be at least 0")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		sequence[i] = CalculateA000262(int64(i))
	}

	return &NumericSequence{
		Name:     "Sets of Lists (A000262)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetSetsOfListsAtPosition returns the n-th term of A000262 (n >= 0).
func GetSetsOfListsAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	result := CalculateA000262(n.Int64())

	return &NumericSequence{
		Name:     "Sets of Lists (A000262)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// CalculateA000262 calculates the n-th term of OEIS A000262.
func CalculateA000262(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	total := new(big.Int)
	// a(n) = sum_{k=1}^n L(n, k)
	// L(n, k) = n!/k! * binom(n-1, k-1)

	factN := factorial(int(n))

	for k := int64(1); k <= n; k++ {
		// n!/k!
		term := new(big.Int).Div(factN, factorial(int(k)))

		// * binom(n-1, k-1)
		term.Mul(term, binomial(n-1, k-1))

		total.Add(total, term)
	}

	return total
}
