package sequences

import (
	"fmt"
	"math/big"
)

// Period 2: repeat [1, 2].
// URL: https://oeis.org/A000034
// Description: a(n) = 1 + (n mod 2).

// GetPeriod12A000034Sequence returns the A000034 sequence.
func GetPeriod12A000034Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetPeriod12A000034AtPosition(maxNumber)
	}
	return GeneratePeriod12A000034Sequence(maxNumber)
}

// GeneratePeriod12A000034Sequence generates the A000034 sequence up to maxNumber.
func GeneratePeriod12A000034Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Period 2: 1, 2 (A000034)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}

	limit := maxNumber.Int64()
	if maxNumber.IsInt64() && limit >= 0 {
		var sequence []*big.Int
		for n := int64(0); n < limit; n++ {
			val := 1 + (n % 2)
			sequence = append(sequence, big.NewInt(val))
		}
		var result *big.Int
		if len(sequence) > 0 {
			result = sequence[len(sequence)-1]
		} else {
			result = big.NewInt(0)
		}
		return &NumericSequence{
			Name:     "Period 2: 1, 2 (A000034)",
			Number:   maxNumber,
			Sequence: sequence,
			Result:   result,
		}, nil
	}

	return nil, fmt.Errorf("maxNumber too large for current implementation")
}

// GetPeriod12A000034AtPosition returns the n-th term of A000034.
func GetPeriod12A000034AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be non-negative")
	}

	rem := new(big.Int).Mod(n, big.NewInt(2))
	val := new(big.Int).Add(big.NewInt(1), rem)
	return &NumericSequence{
		Name:     "Period 2: 1, 2 (A000034)",
		Number:   n,
		Sequence: []*big.Int{val},
		Result:   val,
	}, nil
}
