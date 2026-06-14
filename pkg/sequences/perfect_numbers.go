package sequences

import (
	"fmt"
	"math/big"
)

// Perfect numbers (A000396)
// URL: https://oeis.org/A000396
// Description: Numbers n such that the sum of the divisors of n is 2n.
// Formula: Even perfect numbers are of the form 2^{p-1}(2^p - 1) where 2^p - 1 is a Mersenne prime.

// mersenneExponents are the exponents p such that 2^p - 1 is a Mersenne prime (OEIS A000043).
var mersenneExponents = []int64{
	2, 3, 5, 7, 13, 17, 19, 31, 61, 89, 107, 127, 521, 607, 1279, 2203, 2281, 3217, 4253, 4423,
	9689, 9941, 11213, 19937, 21701, 23209, 44497, 86243, 110503, 132049, 216091, 756839, 859433,
	1257787, 1398269, 2976221, 3021377, 6972593, 13466917, 20996011, 24036583, 25964951, 30402457,
	32582657, 37156667, 42643801, 43112609, 57885161, 74207281, 77232917, 82589933, 136279841,
}

// GetPerfectNumbersSequence returns the sequence of perfect numbers.
func GetPerfectNumbersSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetPerfectNumberAtPosition(maxNumber)
	}
	return GeneratePerfectNumbersSequence(maxNumber)
}

// GeneratePerfectNumbersSequence generates perfect numbers up to maxNumber.
func GeneratePerfectNumbersSequence(maxNumber *big.Int) (*NumericSequence, error) {
	var sequence []*big.Int
	for _, p := range mersenneExponents {
		pn := calculatePerfectNumber(p)
		if pn.Cmp(maxNumber) > 0 {
			break
		}
		sequence = append(sequence, pn)
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Perfect numbers (A000396)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetPerfectNumberAtPosition returns the n-th perfect number.
func GetPerfectNumberAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() <= 0 {
		return nil, fmt.Errorf("position must be a positive integer")
	}

	idx := int(n.Int64()) - 1
	if !n.IsInt64() || idx >= len(mersenneExponents) {
		return nil, fmt.Errorf("position %s exceeds known perfect numbers", n.String())
	}

	pn := calculatePerfectNumber(mersenneExponents[idx])

	return &NumericSequence{
		Name:     "Perfect numbers (A000396)",
		Number:   n,
		Sequence: []*big.Int{pn},
		Result:   pn,
	}, nil
}

// calculatePerfectNumber computes 2^{p-1}(2^p - 1).
func calculatePerfectNumber(p int64) *big.Int {
	// 2^{p-1}
	twoPMinus1 := new(big.Int).Lsh(big.NewInt(1), uint(p-1))
	// 2^p - 1
	mersennePrime := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), uint(p)), big.NewInt(1))
	// 2^{p-1} * (2^p - 1)
	return new(big.Int).Mul(twoPMinus1, mersennePrime)
}
