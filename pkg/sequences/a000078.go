package sequences

import (
	"math/big"
)

// A000078: Tetranacci numbers: a(n) = a(n-1) + a(n-2) + a(n-3) + a(n-4) for n >= 4 with a(0)=a(1)=a(2)=0 and a(3)=1.
func GetA000078Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		n := maxNumber.Int64()
		val := tetranacci(n)
		return &NumericSequence{
			Name:     "Tetranacci numbers (A000078)",
			Number:   maxNumber,
			Sequence: []*big.Int{val},
			Result:   val,
		}, nil
	}

	limit := maxNumber.Int64()
	var sequence []*big.Int
	for n := int64(0); n < limit; n++ {
		sequence = append(sequence, tetranacci(n))
	}
	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	} else {
		result = big.NewInt(0)
	}
	return &NumericSequence{
		Name:     "Tetranacci numbers (A000078)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

func tetranacci(n int64) *big.Int {
	if n < 3 {
		return big.NewInt(0)
	}
	if n == 3 {
		return big.NewInt(1)
	}
	a, b, c, d := big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(1)
	for i := int64(4); i <= n; i++ {
		next := new(big.Int).Add(a, b)
		next.Add(next, c)
		next.Add(next, d)
		a, b, c, d = b, c, d, next
	}
	return d
}
