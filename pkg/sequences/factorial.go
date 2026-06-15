package sequences

import (
	"fmt"
	"math/big"
)

// Factorial numbers (OEIS A000142)
// URL: https://oeis.org/A000142
// a(n) = n! = 1*2*3*...*n.
// a(0) = 1.

// GetFactorialSequence returns the factorial sequence up to maxNumber (OEIS A000142).
func GetFactorialSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetFactorialAtPosition(maxNumber)
	}
	return GenerateFactorialSequence(maxNumber)
}

// GenerateFactorialSequence generates the A000142 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(n) where a(n) <= maxNumber.
func GenerateFactorialSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Factorial numbers (A000142)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   nil,
		}, nil
	}

	var sequence []*big.Int
	current := big.NewInt(1)
	sequence = append(sequence, new(big.Int).Set(current))

	for i := int64(1); ; i++ {
		next := new(big.Int).Mul(current, big.NewInt(i))
		if next.Cmp(maxNumber) > 0 {
			break
		}
		current = next
		sequence = append(sequence, new(big.Int).Set(current))
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Factorial numbers (A000142)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetFactorialAtPosition returns the n-th term of A000142 (n >= 0).
func GetFactorialAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be a non-negative integer")
	}

	result := CalculateFactorial(n.Int64())

	return &NumericSequence{
		Name:     "Factorial numbers (A000142)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// CalculateFactorial calculates the n-th factorial number.
// a(n) = n!
func CalculateFactorial(n int64) *big.Int {
	res := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		res.Mul(res, big.NewInt(i))
	}
	return res
}
