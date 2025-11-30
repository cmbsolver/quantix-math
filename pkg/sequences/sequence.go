package sequences

import (
	"math"
	"math/big"
)

// GetCakeSequence generates the Cake sequence.
func GetCakeSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	retval := &NumericSequence{Name: "Cake", Number: new(big.Int).Set(maxNumber)}
	numberToCalculate := new(big.Int).Set(maxNumber)
	if isPositional {
		numberToCalculate = new(big.Int).SetUint64(math.MaxUint64)
	}

	for n := new(big.Int).SetUint64(0); n.Cmp(numberToCalculate) <= 0; n.Add(n, big.NewInt(1)) {
		if !isPositional {
			next := new(big.Int).Div(
				new(big.Int).Add(
					new(big.Int).Add(
						new(big.Int).Mul(new(big.Int).Set(n), new(big.Int).Mul(new(big.Int).Set(n), new(big.Int).Set(n))),
						new(big.Int).Mul(new(big.Int).SetUint64(5), new(big.Int).Set(n)),
					),
					new(big.Int).SetUint64(6),
				),
				new(big.Int).SetUint64(6),
			)
			retval.Sequence = append(retval.Sequence, next)
		} else {
			if n.Cmp(maxNumber) == 0 {
				next := new(big.Int).Div(
					new(big.Int).Add(
						new(big.Int).Add(
							new(big.Int).Mul(new(big.Int).Set(n), new(big.Int).Mul(new(big.Int).Set(n), new(big.Int).Set(n))),
							new(big.Int).Mul(new(big.Int).SetUint64(5), new(big.Int).Set(n)),
						),
						new(big.Int).SetUint64(6),
					),
					new(big.Int).SetUint64(6),
				)
				retval.Sequence = append(retval.Sequence, next)
				break
			}
		}
	}

	return retval, nil
}

// GetCatalanSequence generates the Catalan sequence.
func GetCatalanSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	retval := &NumericSequence{Name: "Catalan", Number: new(big.Int).Set(maxNumber)}
	numberToCalculate := new(big.Int).Set(maxNumber)
	if isPositional {
		numberToCalculate = new(big.Int).SetUint64(math.MaxUint64)
	}

	n := new(big.Int).SetUint64(0)
	catalan := new(big.Int).SetUint64(1)

	for catalan.Cmp(numberToCalculate) <= 0 {
		catalan = new(big.Int).Div(
			new(big.Int).Mul(
				new(big.Int).Mul(big.NewInt(2), new(big.Int).Add(new(big.Int).Mul(big.NewInt(2), n), big.NewInt(1))),
				catalan,
			),
			new(big.Int).Add(n, big.NewInt(2)),
		)
		n.Add(n, big.NewInt(1))

		if !isPositional {
			if catalan.Cmp(maxNumber) > 0 {
				break
			} else {
				retval.Sequence = append(retval.Sequence, new(big.Int).Set(catalan))
			}
		} else {
			if n.Cmp(maxNumber) == 0 {
				retval.Sequence = append(retval.Sequence, new(big.Int).Set(catalan))
				break
			}
		}
	}

	return retval, nil
}
