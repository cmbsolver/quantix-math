package sequences

import (
	"fmt"
	"math/big"
)

// Motzkin numbers (OEIS A001006)
// URL: https://oeis.org/A001006
// Description: Number of ways of drawing any number of nonintersecting chords joining n labeled points on a circle.
// Offset 0: 1, 1, 2, 4, 9, 21, 51, 127, ...
// Recurrence: a(0)=1, a(1)=1, and for n >= 2,
// a(n) = ((2*n+1)*a(n-1) + (3*n-3)*a(n-2)) / (n+2).

// GetMotzkinNumbersA001006Sequence returns the Motzkin numbers sequence (A001006).
func GetMotzkinNumbersA001006Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetMotzkinNumbersA001006AtPosition(maxNumber)
	}

	return GenerateMotzkinNumbersA001006Sequence(maxNumber)
}

// GenerateMotzkinNumbersA001006Sequence generates A001006 terms up to maxNumber.
func GenerateMotzkinNumbersA001006Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Motzkin numbers (A001006)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   nil,
		}, nil
	}

	sequence := make([]*big.Int, 0)
	for n := 0; ; n++ {
		value := calculateMotzkinNumbersA001006(n)
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
		Name:     "Motzkin numbers (A001006)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetMotzkinNumbersA001006AtPosition returns the term at index n in A001006.
func GetMotzkinNumbersA001006AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be non-negative")
	}

	result := calculateMotzkinNumbersA001006(int(n.Int64()))

	return &NumericSequence{
		Name:     "Motzkin numbers (A001006)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateMotzkinNumbersA001006 computes a(n) for OEIS A001006.
func calculateMotzkinNumbersA001006(n int) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	if n <= 1 {
		return big.NewInt(1)
	}

	aPrev2 := big.NewInt(1) // a(0)
	aPrev1 := big.NewInt(1) // a(1)

	for k := 2; k <= n; k++ {
		term1 := new(big.Int).Mul(big.NewInt(int64(2*k+1)), aPrev1)
		term2 := new(big.Int).Mul(big.NewInt(int64(3*k-3)), aPrev2)
		numerator := new(big.Int).Add(term1, term2)
		current := new(big.Int).Div(numerator, big.NewInt(int64(k+2)))

		aPrev2 = aPrev1
		aPrev1 = current
	}

	return aPrev1
}
