package sequences

import (
	"fmt"
	"math/big"
)

// Schroeder's second problem (OEIS A001003)
// URL: https://oeis.org/A001003
// Description: Schroeder's second problem (little Schroeder/super-Catalan numbers).
// Offset 0: 1, 1, 3, 11, 45, 197, ...
// Recurrence: a(0)=1, a(1)=1, and for n >= 2,
// a(n) = ((6*n-3)*a(n-1) - (n-2)*a(n-2)) / (n+1).

// GetSchroederSecondA001003Sequence returns the Schroeder second sequence (A001003).
func GetSchroederSecondA001003Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetSchroederSecondA001003AtPosition(maxNumber)
	}

	return GenerateSchroederSecondA001003Sequence(maxNumber)
}

// GenerateSchroederSecondA001003Sequence generates A001003 terms up to maxNumber.
func GenerateSchroederSecondA001003Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Schroeder's second problem (A001003)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   nil,
		}, nil
	}

	sequence := make([]*big.Int, 0)
	for n := 0; ; n++ {
		value := calculateSchroederSecondA001003(n)
		if value.Cmp(maxNumber) > 0 {
			break
		}
		sequence = append(sequence, value)
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Schroeder's second problem (A001003)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetSchroederSecondA001003AtPosition returns the term at index n in A001003.
func GetSchroederSecondA001003AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be non-negative")
	}

	result := calculateSchroederSecondA001003(int(n.Int64()))

	return &NumericSequence{
		Name:     "Schroeder's second problem (A001003)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateSchroederSecondA001003 computes a(n) for OEIS A001003.
func calculateSchroederSecondA001003(n int) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	if n <= 1 {
		return big.NewInt(1)
	}

	aPrev2 := big.NewInt(1) // a(0)
	aPrev1 := big.NewInt(1) // a(1)

	for k := 2; k <= n; k++ {
		term1 := new(big.Int).Mul(big.NewInt(int64(6*k-3)), aPrev1)
		term2 := new(big.Int).Mul(big.NewInt(int64(k-2)), aPrev2)
		numerator := new(big.Int).Sub(term1, term2)
		current := new(big.Int).Div(numerator, big.NewInt(int64(k+1)))

		aPrev2 = aPrev1
		aPrev1 = current
	}

	return aPrev1
}
