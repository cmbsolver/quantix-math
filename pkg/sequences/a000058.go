package sequences

import (
	"math/big"
)

// A000058: Sylvester's sequence: a(n+1) = a(n)^2 - a(n) + 1, with a(0) = 2.
func GetA000058Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		n := maxNumber.Int64()
		val := big.NewInt(2)
		for i := int64(0); i < n; i++ {
			// val = val*val - val + 1
			sq := new(big.Int).Mul(val, val)
			val = new(big.Int).Sub(sq, val)
			val.Add(val, big.NewInt(1))
		}
		return &NumericSequence{
			Name:     "Sylvester's sequence (A000058)",
			Number:   maxNumber,
			Sequence: []*big.Int{val},
			Result:   val,
		}, nil
	}

	limit := maxNumber.Int64()
	var sequence []*big.Int
	val := big.NewInt(2)
	for n := int64(0); n < limit; n++ {
		sequence = append(sequence, new(big.Int).Set(val))
		sq := new(big.Int).Mul(val, val)
		val = new(big.Int).Sub(sq, val)
		val.Add(val, big.NewInt(1))
	}
	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	} else {
		result = big.NewInt(0)
	}
	return &NumericSequence{
		Name:     "Sylvester's sequence (A000058)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}
