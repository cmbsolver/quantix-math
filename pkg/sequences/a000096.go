package sequences

import (
	"math/big"
)

// A000096: a(n) = n*(n+3)/2.
func GetA000096Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		n := new(big.Int).Set(maxNumber)
		val := new(big.Int).Add(n, big.NewInt(3))
		val.Mul(val, n)
		val.Div(val, big.NewInt(2))
		return &NumericSequence{
			Name:     "n*(n+3)/2 (A000096)",
			Number:   maxNumber,
			Sequence: []*big.Int{val},
			Result:   val,
		}, nil
	}

	limit := maxNumber.Int64()
	var sequence []*big.Int
	for i := int64(0); i < limit; i++ {
		n := big.NewInt(i)
		val := new(big.Int).Add(n, big.NewInt(3))
		val.Mul(val, n)
		val.Div(val, big.NewInt(2))
		sequence = append(sequence, val)
	}
	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	} else {
		result = big.NewInt(0)
	}
	return &NumericSequence{
		Name:     "n*(n+3)/2 (A000096)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}
