package sequences

import (
	"math/big"
)

// A000073: Tribonacci numbers: a(n) = a(n-1) + a(n-2) + a(n-3) for n >= 3 with a(0) = a(1) = 0 and a(2) = 1.
func GetA000073Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		n := maxNumber.Int64()
		val := tribonacci(n)
		return &NumericSequence{
			Name:     "Tribonacci numbers (A000073)",
			Number:   maxNumber,
			Sequence: []*big.Int{val},
			Result:   val,
		}, nil
	}

	limit := maxNumber.Int64()
	var sequence []*big.Int
	for n := int64(0); n < limit; n++ {
		sequence = append(sequence, tribonacci(n))
	}
	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	} else {
		result = big.NewInt(0)
	}
	return &NumericSequence{
		Name:     "Tribonacci numbers (A000073)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

func tribonacci(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	if n == 1 {
		return big.NewInt(0)
	}
	if n == 2 {
		return big.NewInt(1)
	}
	a, b, c := big.NewInt(0), big.NewInt(0), big.NewInt(1)
	for i := int64(3); i <= n; i++ {
		next := new(big.Int).Add(a, b)
		next.Add(next, c)
		a, b, c = b, c, next
	}
	return c
}
