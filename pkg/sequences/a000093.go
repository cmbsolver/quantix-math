package sequences

import (
	"math"
	"math/big"
)

// A000093: a(n) = floor(n^(3/2)).
func GetA000093Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		n := float64(maxNumber.Int64())
		val := int64(math.Floor(math.Pow(n, 1.5)))
		res := big.NewInt(val)
		return &NumericSequence{
			Name:     "floor(n^(3/2)) (A000093)",
			Number:   maxNumber,
			Sequence: []*big.Int{res},
			Result:   res,
		}, nil
	}

	limit := maxNumber.Int64()
	var sequence []*big.Int
	for n := int64(0); n < limit; n++ {
		val := int64(math.Floor(math.Pow(float64(n), 1.5)))
		sequence = append(sequence, big.NewInt(val))
	}
	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	} else {
		result = big.NewInt(0)
	}
	return &NumericSequence{
		Name:     "floor(n^(3/2)) (A000093)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}
