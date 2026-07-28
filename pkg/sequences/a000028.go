package sequences

import (
	"fmt"
	"math/big"
)

// A000028: Let k = p_1^e_1 p_2^e_2 p_3^e_3 ... be the prime factorization of n.
// Sequence gives k such that the sum of the numbers of 1's in the binary expansions of e_1, e_2, e_3, ... is odd.
// URL: https://oeis.org/A000028
// Formerly M0520 N0187

// GetA000028Sequence returns the sequence A000028.
// If isPositional is true, it returns the n-th term of the sequence.
// Otherwise, it returns all terms up to maxNumber terms.
func GetA000028Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000028AtPosition(maxNumber)
	}
	return GenerateA000028Sequence(maxNumber)
}

// GenerateA000028Sequence generates the first 'limit' terms of A000028.
func GenerateA000028Sequence(limit *big.Int) (*NumericSequence, error) {
	if limit.Sign() < 0 {
		return nil, fmt.Errorf("limit cannot be negative")
	}

	nLimit := int(limit.Int64())
	sequence := make([]*big.Int, 0, nLimit)

	for k := int64(1); len(sequence) < nLimit; k++ {
		if isA000028(k) {
			sequence = append(sequence, big.NewInt(k))
		}
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	} else {
		result = big.NewInt(0)
	}

	return &NumericSequence{
		Name:     "Binary weight of exponents is odd (A000028)",
		Number:   limit,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetA000028AtPosition returns the n-th term of A000028.
func GetA000028AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() <= 0 {
		return nil, fmt.Errorf("position must be positive")
	}

	nPos := n.Int64()
	count := int64(0)
	for k := int64(1); ; k++ {
		if isA000028(k) {
			count++
			if count == nPos {
				val := big.NewInt(k)
				return &NumericSequence{
					Name:     "Binary weight of exponents is odd (A000028)",
					Number:   n,
					Sequence: []*big.Int{val},
					Result:   val,
				}, nil
			}
		}
	}
}

// isA000028 checks if a number k belongs to sequence A000028.
func isA000028(k int64) bool {
	if k < 1 {
		return false
	}
	if k == 1 {
		// Based on the sequence data 2, 3, 4, 5... 1 is not in the sequence.
		// For k=1, the prime factorization is empty, so the sum is 0 (even).
		return false
	}

	exponents := factorizeA000028(k)
	totalOnes := 0
	for _, e := range exponents {
		totalOnes += countSetBits(int64(e))
	}
	return totalOnes%2 != 0
}

// factorizeA000028 returns the exponents of the prime factorization of n.
func factorizeA000028(n int64) []int {
	exponents := []int{}
	d := int64(2)
	temp := n
	for d*d <= temp {
		if temp%d == 0 {
			count := 0
			for temp%d == 0 {
				count++
				temp /= d
			}
			exponents = append(exponents, count)
		}
		d++
	}
	if temp > 1 {
		exponents = append(exponents, 1)
	}
	return exponents
}

// countSetBits returns the number of set bits (1s) in the binary representation of n.
func countSetBits(n int64) int {
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}
