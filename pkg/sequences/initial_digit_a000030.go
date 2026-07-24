package sequences

import (
	"fmt"
	"math/big"
)

// Initial digit of n.
// URL: https://oeis.org/A000030
// Description: a(n) is the first digit of n.

// GetInitialDigitA000030Sequence returns the A000030 sequence.
func GetInitialDigitA000030Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetInitialDigitA000030AtPosition(maxNumber)
	}
	return GenerateInitialDigitA000030Sequence(maxNumber)
}

// GenerateInitialDigitA000030Sequence generates the A000030 sequence up to maxNumber.
func GenerateInitialDigitA000030Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Initial digit (A000030)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}

	limit := maxNumber.Int64()
	if maxNumber.IsInt64() && limit >= 0 {
		var sequence []*big.Int
		for n := int64(0); n < limit; n++ {
			sequence = append(sequence, big.NewInt(int64(getInitialDigit(n))))
		}
		var result *big.Int
		if len(sequence) > 0 {
			result = sequence[len(sequence)-1]
		} else {
			result = big.NewInt(0)
		}
		return &NumericSequence{
			Name:     "Initial digit (A000030)",
			Number:   maxNumber,
			Sequence: sequence,
			Result:   result,
		}, nil
	}

	return nil, fmt.Errorf("maxNumber too large for current implementation")
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
