package sequences

import (
	"fmt"
	"math/big"
)

// Number of primitive n-bead necklaces (turning over is allowed) where complements are equivalent (OEIS A000046).
// URL: https://oeis.org/A000046

// GetA000046Sequence returns the A000046 sequence.
func GetA000046Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000046AtPosition(maxNumber)
	}

	return GenerateA000046Sequence(maxNumber)
}

// GenerateA000046Sequence generates A000046 from a(0) through a(maxNumber).
func GenerateA000046Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := maxNumber.Int64()
	sequence := make([]*big.Int, limit+1)
	for i := int64(0); i <= limit; i++ {
		sequence[i] = calculateA000046(i)
	}

	return &NumericSequence{
		Name:     "Primitive necklaces (turnover, complement equivalent) (A000046)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[limit],
	}, nil
}

// GetA000046AtPosition returns the n-th term of A000046.
func GetA000046AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	result := calculateA000046(n.Int64())

	return &NumericSequence{
		Name:     "Primitive necklaces (turnover, complement equivalent) (A000046)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateA000046 computes a(n) for OEIS A000046 using the formula from OEIS:
// a(0) = 1, and for n > 0,
// a(n) = ( ((1/(2n)) * Sum_{odd d|n} mu(d)*2^(n/d)) + Sum_{d|n} mu(n/d)*2^floor(d/2) ) / 2.
func calculateA000046(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	firstNumerator := new(big.Int)
	for _, d := range getOddDivisors(n) {
		mu := getMobius(d)
		if mu == 0 {
			continue
		}

		term := new(big.Int).Lsh(big.NewInt(1), uint(n/d))
		if mu < 0 {
			firstNumerator.Sub(firstNumerator, term)
		} else {
			firstNumerator.Add(firstNumerator, term)
		}
	}

	firstPart := new(big.Int).Div(firstNumerator, big.NewInt(2*n))

	secondPart := new(big.Int)
	for _, d := range getDivisorsA000046(n) {
		mu := getMobius(n / d)
		if mu == 0 {
			continue
		}

		term := new(big.Int).Lsh(big.NewInt(1), uint(d/2))
		if mu < 0 {
			secondPart.Sub(secondPart, term)
		} else {
			secondPart.Add(secondPart, term)
		}
	}

	result := new(big.Int).Add(firstPart, secondPart)
	result.Div(result, big.NewInt(2))

	return result
}

func getDivisorsA000046(n int64) []int64 {
	divs := make([]int64, 0)
	for i := int64(1); i*i <= n; i++ {
		if n%i == 0 {
			divs = append(divs, i)
			j := n / i
			if j != i {
				divs = append(divs, j)
			}
		}
	}
	return divs
}
