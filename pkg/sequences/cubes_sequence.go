package sequences

import (
	"math/big"
)

// GetCubesSequence generates the cubes sequence.
func GetCubesSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	retval := &NumericSequence{Name: "Cubes", Number: new(big.Int).Set(maxNumber)}
	numberToCalculate := new(big.Int).Set(maxNumber)
	if isPositional {
		numberToCalculate = new(big.Int).SetUint64(^uint64(0)) // Max uint64 value
	}

	for n := big.NewInt(0); n.Cmp(numberToCalculate) <= 0; n.Add(n, big.NewInt(1)) {
		cube := new(big.Int).Mul(n, n)
		cube.Mul(cube, n)
		if !isPositional {
			retval.Sequence = append(retval.Sequence, cube)
		} else {
			if n.Cmp(maxNumber) == 0 {
				retval.Sequence = append(retval.Sequence, cube)
				break
			}
		}
	}

	return retval, nil
}
