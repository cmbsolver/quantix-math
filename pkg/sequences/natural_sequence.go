package sequences

import (
	"math/big"
)

// GetNaturalSequence generates the natural number sequence.
func GetNaturalSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	retval := &NumericSequence{Name: "Natural", Number: new(big.Int).Set(maxNumber)}
	numberToCalculate := new(big.Int).Set(maxNumber)
	if isPositional {
		numberToCalculate = new(big.Int).SetUint64(^uint64(0)) // Max uint64 value
	}

	for n := big.NewInt(0); n.Cmp(numberToCalculate) <= 0; n.Add(n, big.NewInt(1)) {
		if !isPositional {
			retval.Sequence = append(retval.Sequence, new(big.Int).Set(n))
		} else {
			if n.Cmp(maxNumber) == 0 {
				retval.Sequence = append(retval.Sequence, new(big.Int).Set(n))
				break
			}
		}
	}

	return retval, nil
}
