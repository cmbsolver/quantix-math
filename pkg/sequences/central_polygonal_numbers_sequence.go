package sequences

import (
	"math/big"
)

// GetCentralPolygonalNumbersSequence generates the central polygonal numbers sequence.
func GetCentralPolygonalNumbersSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	retval := &NumericSequence{Name: "Central Polygonal Numbers", Number: new(big.Int).Set(maxNumber)}
	numberToCalculate := new(big.Int).Set(maxNumber)
	if isPositional {
		numberToCalculate = new(big.Int).SetUint64(^uint64(0)) // Max uint64 value
	}

	for n := big.NewInt(0); n.Cmp(numberToCalculate) <= 0; n.Add(n, big.NewInt(1)) {
		if !isPositional {
			item := new(big.Int).Mul(n, new(big.Int).Add(n, big.NewInt(1)))
			item.Div(item, big.NewInt(2))
			item.Add(item, big.NewInt(1))
			retval.Sequence = append(retval.Sequence, item)
		} else {
			if n.Cmp(maxNumber) == 0 {
				item := new(big.Int).Mul(n, new(big.Int).Sub(n, big.NewInt(1)))
				item.Div(item, big.NewInt(2))
				item.Add(item, big.NewInt(1))
				retval.Sequence = append(retval.Sequence, item)
				break
			}
		}
	}

	return retval, nil
}
