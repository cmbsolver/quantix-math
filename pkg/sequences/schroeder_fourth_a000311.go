package sequences

import (
	"fmt"
	"math/big"
)

// Schroeder's fourth problem (OEIS A000311)
// URL: https://oeis.org/A000311
// Description: Schroeder's fourth problem; also series-reduced rooted trees with n labeled leaves;
// also number of total partitions of n.
// a(1) = 1; for n > 1, a(n) = sum_{k=1..n-1} binom(n-1, k) * a(n-k) * E(k),
// where E(1) = 1 and E(k) = 2*a(k) for k > 1.

// GetSchroederFourthSequence returns the Schroeder's fourth problem sequence (A000311).
func GetSchroederFourthSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetSchroederFourthAtPosition(maxNumber)
	}
	return GenerateSchroederFourthSequence(maxNumber)
}

// GenerateSchroederFourthSequence generates A000311 sequence up to maxNumber.
func GenerateSchroederFourthSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Schroeder's fourth problem (A000311)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   nil,
		}, nil
	}

	var sequence []*big.Int
	// We'll start from n=1 as a(0) is often defined as 0 or 1, but the main sequence starts from a(1)=1.
	// Based on research, it starts 1, 1, 4, 26, 236... for n=1, 2, 3, 4, 5...
	// If maxNumber is small, we should include a(0)=0 if we want to follow the offset 0.
	// But usually these labeled problems start from 1. The OEIS says offset 0, and a(0)=0.
	sequence = append(sequence, big.NewInt(0)) // a(0) = 0

	for n := 1; ; n++ {
		val := calculateSchroederFourth(n)
		if val.Cmp(maxNumber) > 0 {
			break
		}
		sequence = append(sequence, val)
		if n > 500 { // Safety break
			break
		}
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Schroeder's fourth problem (A000311)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetSchroederFourthAtPosition returns the n-th term of A000311.
func GetSchroederFourthAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be non-negative")
	}

	pos := int(n.Int64())
	result := calculateSchroederFourth(pos)

	return &NumericSequence{
		Name:     "Schroeder's fourth problem (A000311)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateSchroederFourth computes the n-th term of A000311.
func calculateSchroederFourth(n int) *big.Int {
	if n <= 0 {
		return big.NewInt(0)
	}
	if n == 1 {
		return big.NewInt(1)
	}

	// We use DP to store previous values of a(n) and E(n)
	a := make([]*big.Int, n+1)
	e := make([]*big.Int, n+1)

	a[1] = big.NewInt(1)
	e[1] = big.NewInt(1)

	for i := 2; i <= n; i++ {
		sum := big.NewInt(0)
		for k := 1; k < i; k++ {
			// binom(i-1, k) * a[i-k] * e[k]
			term := new(big.Int).Binomial(int64(i-1), int64(k))
			term.Mul(term, a[i-k])
			term.Mul(term, e[k])
			sum.Add(sum, term)
		}
		a[i] = new(big.Int).Set(sum)
		e[i] = new(big.Int).Mul(a[i], big.NewInt(2))
	}

	return a[n]
}
