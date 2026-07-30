package sequences

import (
	"fmt"
	"math/big"
)

// Initial digit of n (OEIS A000030)
// URL: https://oeis.org/A000030
// Initial digit of n: a(n) is the first digit of n.
// a(0) = 0.

// GetInitialDigitA000030Sequence returns the A000030 sequence (initial digit of n).
func GetInitialDigitA000030Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetInitialDigitA000030AtPosition(maxNumber)
	}
	return GenerateInitialDigitA000030Sequence(maxNumber)
}

// GenerateInitialDigitA000030Sequence generates the A000030 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateInitialDigitA000030Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		sequence[i] = big.NewInt(int64(getInitialDigit(int64(i))))
	}

	result := big.NewInt(0)
	if n < len(sequence) {
		result = sequence[n]
	}

	return &NumericSequence{
		Name:     "Initial digit (A000030)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetInitialDigitA000030AtPosition returns the n-th term of A000030.
func GetInitialDigitA000030AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be non-negative")
	}

	s := n.String()
	firstDigit := int64(s[0] - '0')
	val := big.NewInt(firstDigit)
	return &NumericSequence{
		Name:     "Initial digit (A000030)",
		Number:   n,
		Sequence: []*big.Int{val},
		Result:   val,
	}, nil
}

func getInitialDigit(n int64) int {
	if n == 0 {
		return 0
	}
	s := fmt.Sprintf("%d", n)
	return int(s[0] - '0')
}
