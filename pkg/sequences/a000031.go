package sequences

import (
	"fmt"
	"math/big"
)

// Number of n-bead necklaces with 2 colors when turning over is not allowed (OEIS A000031).
// URL: https://oeis.org/A000031
// Formula: a(0) = 1; for n > 0, a(n) = (1/n) * Sum_{d|n} phi(d) * 2^(n/d).

// GetA000031Sequence returns the A000031 sequence.
func GetA000031Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000031AtPosition(maxNumber)
	}

	return GenerateA000031Sequence(maxNumber)
}

// GenerateA000031Sequence generates A000031 from a(0) through a(maxNumber).
func GenerateA000031Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	nLimit := int(maxNumber.Int64())
	sequence := make([]*big.Int, nLimit+1)

	for i := 0; i <= nLimit; i++ {
		sequence[i] = calculateA000031(int64(i))
	}

	return &NumericSequence{
		Name:     "Necklaces with no turnover allowed (A000031)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[nLimit],
	}, nil
}

// GetA000031AtPosition returns the n-th term of A000031.
func GetA000031AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	result := calculateA000031(n.Int64())

	return &NumericSequence{
		Name:     "Necklaces with no turnover allowed (A000031)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateA000031 computes a(n) for OEIS A000031.
func calculateA000031(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	sum := big.NewInt(0)
	for d := int64(1); d <= n; d++ {
		if n%d == 0 {
			phiD := big.NewInt(EulerTotient(d))
			twoPow := new(big.Int).Exp(big.NewInt(2), big.NewInt(n/d), nil)
			term := new(big.Int).Mul(phiD, twoPow)
			sum.Add(sum, term)
		}
	}

	return new(big.Int).Div(sum, big.NewInt(n))
}
