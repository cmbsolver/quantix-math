package sequences

import (
	"math"
	"math/big"
)

// A000062: A Beatty sequence: a(n) = floor(n/(e-2)).
func GetA000062Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	eMinus2 := math.E - 2.0
	if isPositional {
		n := float64(maxNumber.Int64())
		val := int64(math.Floor(n / eMinus2))
		res := big.NewInt(val)
		return &NumericSequence{
			Name:     "Beatty sequence floor(n/(e-2)) (A000062)",
			Number:   maxNumber,
			Sequence: []*big.Int{res},
			Result:   res,
		}, nil
	}

	limit := maxNumber.Int64()
	var sequence []*big.Int
	for n := int64(1); n <= limit; n++ {
		val := int64(math.Floor(float64(n) / eMinus2))
		sequence = append(sequence, big.NewInt(val))
	}
	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	} else {
		result = big.NewInt(0)
	}
	return &NumericSequence{
		Name:     "Beatty sequence floor(n/(e-2)) (A000062)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}
