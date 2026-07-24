package sequences

import (
	"math/big"
)

// A000051: a(n) = 2^n + 1.
func GetA000051Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		val := new(big.Int).Exp(big.NewInt(2), maxNumber, nil)
		val.Add(val, big.NewInt(1))
		return &NumericSequence{
			Name:     "2^n + 1 (A000051)",
			Number:   maxNumber,
			Sequence: []*big.Int{val},
			Result:   val,
		}, nil
	}

	limit := maxNumber.Int64()
	var sequence []*big.Int
	for n := int64(0); n < limit; n++ {
		val := new(big.Int).Exp(big.NewInt(2), big.NewInt(n), nil)
		val.Add(val, big.NewInt(1))
		sequence = append(sequence, val)
	}
	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	} else {
		result = big.NewInt(0)
	}
	return &NumericSequence{
		Name:     "2^n + 1 (A000051)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}
