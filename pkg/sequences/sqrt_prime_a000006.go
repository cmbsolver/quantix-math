package sequences

import (
	"fmt"
	"math/big"
)

// Integer part of square root of n-th prime.
// URL: https://oeis.org/A000006
// Description: a(n) = floor(sqrt(prime(n))).

// GetSqrtPrimeA000006Sequence returns the A000006 sequence up to maxNumber or the n-th term.
// If isPositional is true, it returns the n-th term where n = maxNumber.
// If isPositional is false, it returns the first maxNumber terms of the sequence.
func GetSqrtPrimeA000006Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetSqrtPrimeA000006AtPosition(maxNumber)
	}
	return GenerateSqrtPrimeA000006Sequence(maxNumber)
}

// GenerateSqrtPrimeA000006Sequence generates the A000006 sequence up to maxNumber (first n terms).
func GenerateSqrtPrimeA000006Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() <= 0 {
		return &NumericSequence{
			Name:     "Integer part of square root of n-th prime (A000006)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}

	limit := maxNumber.Int64()
	if maxNumber.IsInt64() && limit > 0 {
		var sequence []*big.Int
		count := int64(0)
		p := big.NewInt(1)
		for count < limit {
			p = nextPrime(p)
			sqrtP := new(big.Int).Sqrt(p)
			sequence = append(sequence, sqrtP)
			count++
		}
		var result *big.Int
		if len(sequence) > 0 {
			result = sequence[len(sequence)-1]
		}
		return &NumericSequence{
			Name:     "Integer part of square root of n-th prime (A000006)",
			Number:   maxNumber,
			Sequence: sequence,
			Result:   result,
		}, nil
	}

	return nil, fmt.Errorf("maxNumber too large for current implementation")
}

// GetSqrtPrimeA000006AtPosition returns the n-th term of A000006.
func GetSqrtPrimeA000006AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() <= 0 {
		return nil, fmt.Errorf("position must be positive")
	}

	if n.IsInt64() {
		limit := n.Int64()
		p := big.NewInt(1)
		for i := int64(0); i < limit; i++ {
			p = nextPrime(p)
		}
		val := new(big.Int).Sqrt(p)
		return &NumericSequence{
			Name:     "Integer part of square root of n-th prime (A000006)",
			Number:   n,
			Sequence: []*big.Int{val},
			Result:   val,
		}, nil
	}

	return nil, fmt.Errorf("position too large for current implementation")
}

func nextPrime(n *big.Int) *big.Int {
	next := new(big.Int).Add(n, big.NewInt(1))
	for !next.ProbablyPrime(20) {
		next.Add(next, big.NewInt(1))
	}
	return next
}
