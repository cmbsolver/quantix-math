// Package sequences provides functions to calculate various integer sequences.
package sequences

import (
	"fmt"
	"math/big"
)

// Number of groups of order n (OEIS A000001)
// URL: https://oeis.org/A000001
// a(n) is the number of nonisomorphic groups of order n.
// This sequence is a core sequence in group theory.

// a000001Data contains the values of A000001 for n = 0 to 100.
// Values taken from OEIS A000001.
var a000001Data = []int64{
	0, 1, 1, 1, 2, 1, 2, 1, 5, 2, 2, 1, 5, 1, 2, 1, 14, 1, 5, 1, 5, 2, 2, 1, 15, 2, 2, 5, 4, 1, 4, 1, 51, 1, 2, 1, 14, 1, 2, 2, 14, 1, 6, 1, 4, 2, 2, 1, 52, 2, 5, 1, 5, 1, 15, 2, 13, 2, 2, 1, 13, 1, 2, 4, 267, 1, 4, 1, 5, 1, 4, 1, 50, 1, 2, 3, 4, 1, 6, 1, 52, 15, 2, 1, 15, 1, 2, 1, 12, 1, 10, 1, 4, 2, 2, 1, 231, 1, 5, 2, 16,
}

// CalculateGroupsOrderN calculates the number of groups of order n.
// For large n, this is a hard problem. This implementation uses formulas for specific forms of n
// and a lookup table for other values up to 100.
func CalculateGroupsOrderN(n int64) int64 {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// Use lookup table if available
	if int(n) < len(a000001Data) {
		return a000001Data[n]
	}

	// Try formulas for specific forms
	factors := factorize(n)

	// Case: p^k
	if len(factors) == 1 {
		p := getPrimeFactor(n)
		k := factors[0]
		switch k {
		case 1: // p
			return 1
		case 2: // p^2
			return 2
		case 3: // p^3
			return 5
		case 4: // p^4
			if p == 2 {
				return 14
			}
			return 15
		}
	}

	// Case: p*q (p < q)
	if len(factors) == 2 && factors[0] == 1 && factors[1] == 1 {
		p, q := getTwoPrimeFactors(n)
		if p > q {
			p, q = q, p
		}
		if (q-1)%p == 0 {
			return 2
		}
		return 1
	}

	// Fallback/Unknown
	return -1 // Indicates not implemented for this n
}

func getPrimeFactor(n int64) int64 {
	d := int64(2)
	temp := n
	for d*d <= temp {
		if temp%d == 0 {
			return d
		}
		d++
	}
	return temp
}

func getTwoPrimeFactors(n int64) (int64, int64) {
	var p1 int64
	d := int64(2)
	temp := n
	for d*d <= temp {
		if temp%d == 0 {
			p1 = d
			return p1, n / p1
		}
		d++
	}
	return 1, n
}

// GetGroupsOrderNSequence returns the Number of groups of order n sequence (OEIS A000001).
// Since calculating a(n) for large n is a hard problem, this implementation uses a lookup table
// for known values up to 100 and returns an error for larger values.
func GetGroupsOrderNSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetGroupsOrderNAtPosition(maxNumber)
	}
	return GenerateGroupsOrderNSequence(maxNumber)
}

// GenerateGroupsOrderNSequence generates the A000001 sequence up to maxNumber (a(0), a(1), ..., a(maxNumber)).
func GenerateGroupsOrderNSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := int(maxNumber.Int64())
	sequence := make([]*big.Int, limit+1)
	for i := 0; i <= limit; i++ {
		val := CalculateGroupsOrderN(int64(i))
		if val == -1 {
			return nil, fmt.Errorf("sequence A000001 is not implemented for n=%d", i)
		}
		sequence[i] = big.NewInt(val)
	}

	return &NumericSequence{
		Name:     "Number of Groups (A000001)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[len(sequence)-1],
	}, nil
}

// GetGroupsOrderNAtPosition returns the n-th term of A000001 (n >= 0).
func GetGroupsOrderNAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	pos := n.Int64()
	val := CalculateGroupsOrderN(pos)
	if val == -1 {
		return nil, fmt.Errorf("sequence A000001 is not implemented for n=%d", pos)
	}

	result := big.NewInt(val)

	return &NumericSequence{
		Name:     "Number of Groups (A000001)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
