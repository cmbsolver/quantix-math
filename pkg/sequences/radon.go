package sequences

import (
	"fmt"
	"math/big"
)

// Radon function, also called Hurwitz-Radon numbers (OEIS A003484)
// URL: https://oeis.org/A003484
// If n = 2^k * (2m + 1), where 2m+1 is odd, and k = 4a + b with 0 <= b < 4,
// then a(n) = 8a + 2^b.

// GetRadonSequence returns the Radon numbers up to maxNumber (OEIS A003484).
func GetRadonSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetRadonAtPosition(maxNumber)
	}
	return GenerateRadonSequence(maxNumber)
}

// GenerateRadonSequence generates the A003484 sequence up to maxNumber.
func GenerateRadonSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("max number must be at least 1")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n)

	for i := 1; i <= n; i++ {
		sequence[i-1] = CalculateRadon(int64(i))
	}

	return &NumericSequence{
		Name:     "Radon function (A003484)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n-1],
	}, nil
}

// GetRadonAtPosition returns the n-th term of A003484.
func GetRadonAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("position must be at least 1")
	}

	result := CalculateRadon(n.Int64())

	return &NumericSequence{
		Name:     "Radon function (A003484)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// CalculateRadon calculates the Radon number for n.
// n = 2^k * odd_part
// k = 4a + b (0 <= b < 4)
// rho(n) = 8a + 2^b
func CalculateRadon(n int64) *big.Int {
	if n <= 0 {
		return big.NewInt(0)
	}

	k := 0
	temp := n
	for temp%2 == 0 {
		temp /= 2
		k++
	}

	a := k / 4
	b := k % 4

	// 2^b
	twoToB := int64(1) << uint(b)

	result := 8*int64(a) + twoToB
	return big.NewInt(result)
}
