package sequences

import (
	"fmt"
	"math/big"
)

// A001037: Number of degree-n irreducible polynomials over GF(2);
// number of n-bead necklaces with beads of 2 colors when turning over is not allowed
// and with primitive period n; number of binary Lyndon words of length n.
// URL: https://oeis.org/A001037

// GetA001037Sequence returns the A001037 sequence based on either positional or max-value semantics.
func GetA001037Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA001037AtPosition(maxNumber)
	}
	return GenerateA001037Sequence(maxNumber)
}

// GenerateA001037Sequence generates all A001037 terms less than or equal to maxNumber.
func GenerateA001037Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	sequence := make([]*big.Int, 0)
	for n := int64(0); ; n++ {
		term := calculateA001037(n)
		if term.Cmp(maxNumber) > 0 {
			break
		}
		sequence = append(sequence, term)
	}

	if len(sequence) == 0 {
		return &NumericSequence{
			Name:     "Binary Lyndon words count (A001037)",
			Number:   maxNumber,
			Sequence: sequence,
			Result:   big.NewInt(0),
		}, nil
	}

	return &NumericSequence{
		Name:     "Binary Lyndon words count (A001037)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   new(big.Int).Set(sequence[len(sequence)-1]),
	}, nil
}

// GetA001037AtPosition returns the n-th A001037 term using OEIS offset 0.
func GetA001037AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position must be >= 0")
	}

	if !n.IsInt64() {
		return nil, fmt.Errorf("position is too large")
	}

	index := n.Int64()
	result := calculateA001037(index)
	return &NumericSequence{
		Name:     "Binary Lyndon words count (A001037)",
		Number:   n,
		Sequence: []*big.Int{new(big.Int).Set(result)},
		Result:   result,
	}, nil
}

// calculateA001037 computes a(n) from the OEIS formula:
// a(0) = 1 and for n >= 1,
// a(n) = (1/n) * Sum_{d|n} mu(n/d) * 2^d.
func calculateA001037(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	divisors := getDivisorsA001037(n)
	sum := big.NewInt(0)

	for _, d := range divisors {
		mu := getMobius(n / d)
		if mu == 0 {
			continue
		}

		term := new(big.Int).Exp(big.NewInt(2), big.NewInt(d), nil)
		if mu > 0 {
			sum.Add(sum, term)
		} else {
			sum.Sub(sum, term)
		}
	}

	return sum.Div(sum, big.NewInt(n))
}

// getDivisorsA001037 returns all positive divisors of n.
func getDivisorsA001037(n int64) []int64 {
	divisors := make([]int64, 0)
	for i := int64(1); i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}
	return divisors
}
