package sequences

import (
	"fmt"
	"math/big"
)

// Powers of 2: a(n) = 2^n.
// URL: https://oeis.org/A000079
// a(n) = 2^n.

// GetPowersOf2Sequence returns the powers of 2 sequence (OEIS A000079).
func GetPowersOf2Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetPowersOf2SequenceAtPosition(maxNumber)
	}
	return GeneratePowersOf2Sequence(maxNumber)
}

// GeneratePowersOf2Sequence generates the A000079 sequence up to maxNumber (a(0), a(1), ..., a(maxNumber)).
func GeneratePowersOf2Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	base := big.NewInt(2)
	for i := 0; i <= n; i++ {
		sequence[i] = new(big.Int).Exp(base, big.NewInt(int64(i)), nil)
	}

	result := sequence[n]

	return &NumericSequence{
		Name:     "Powers of 2 (A000079)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetPowersOf2SequenceAtPosition returns the n-th term of A000079 (n >= 0).
func GetPowersOf2SequenceAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	base := big.NewInt(2)
	result := new(big.Int).Exp(base, n, nil)

	return &NumericSequence{
		Name:     "Powers of 2 (A000079)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
