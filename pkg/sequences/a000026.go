package sequences

import (
	"fmt"
	"math/big"
)

// A000026: Mosaic numbers or multiplicative projection of n: if n = Product (p_j^k_j) then a(n) = Product (p_j * k_j).
// URL: https://oeis.org/A000026
// n = Product (p_j^k_j) -> a(n) = Product (p_j * k_j).
// Multiplicative with a(p^e) = p*e.

// GetA000026Sequence returns the mosaic numbers (OEIS A000026).
func GetA000026Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000026AtPosition(maxNumber)
	}
	return GenerateA000026Sequence(maxNumber)
}

// GenerateA000026Sequence generates the A000026 sequence up to maxNumber.
func GenerateA000026Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("max number must be at least 1")
	}

	limit := int(maxNumber.Int64())
	// Use a reasonable limit for calculation if it exceeds Int64 or is too large
	if !maxNumber.IsInt64() || limit > 10000 {
		limit = 10000
	}

	sequence := make([]*big.Int, limit)
	for i := 1; i <= limit; i++ {
		sequence[i-1] = calculateA000026(int64(i))
	}

	return &NumericSequence{
		Name:     "Mosaic numbers (A000026)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[len(sequence)-1],
	}, nil
}

// GetA000026AtPosition returns the n-th term of A000026.
func GetA000026AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("position must be at least 1")
	}

	result := calculateA000026(n.Int64())

	return &NumericSequence{
		Name:     "Mosaic numbers (A000026)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateA000026 computes the n-th term of A000026.
// a(n) = Product (p_j * k_j) where n = Product (p_j^k_j)
func calculateA000026(n int64) *big.Int {
	if n == 1 {
		return big.NewInt(1)
	}

	result := big.NewInt(1)
	d := int64(2)
	temp := n

	for d*d <= temp {
		if temp%d == 0 {
			count := int64(0)
			for temp%d == 0 {
				count++
				temp /= d
			}
			// a(p^k) = p * k
			term := big.NewInt(d)
			term.Mul(term, big.NewInt(count))
			result.Mul(result, term)
		}
		d++
	}

	if temp > 1 {
		// temp is a prime p with k=1
		result.Mul(result, big.NewInt(temp))
	}

	return result
}
