package sequences

import (
	"math/big"
)

// A000069: Odious numbers: numbers with an odd number of 1's in their binary expansion.
func GetA000069Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		n := maxNumber.Int64()
		// We want the n-th odious number.
		// For odious numbers, they are numbers with odd parity.
		// a(n) = 2n or 2n+1. Actually it's simple:
		// The n-th odious number (starting n=0) is the smallest k such that there are n+1 odious numbers <= k.
		count := int64(0)
		for k := int64(0); ; k++ {
			if isOdious(k) {
				if count == n {
					val := big.NewInt(k)
					return &NumericSequence{
						Name:     "Odious numbers (A000069)",
						Number:   maxNumber,
						Sequence: []*big.Int{val},
						Result:   val,
					}, nil
				}
				count++
			}
		}
	}

	limit := maxNumber.Int64()
	var sequence []*big.Int
	count := int64(0)
	for k := int64(0); count < limit; k++ {
		if isOdious(k) {
			sequence = append(sequence, big.NewInt(k))
			count++
		}
	}
	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	} else {
		result = big.NewInt(0)
	}
	return &NumericSequence{
		Name:     "Odious numbers (A000069)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

func isOdious(n int64) bool {
	ones := 0
	for n > 0 {
		if n&1 == 1 {
			ones++
		}
		n >>= 1
	}
	return ones%2 != 0
}
