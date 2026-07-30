package sequences

import (
	"fmt"
	"math/big"
)

// Dying rabbits: a(0) = 1; for 1 <= n <= 12, a(n) = Fibonacci(n);
// for n >= 13, a(n) = a(n-1) + a(n-2) - a(n-13) (OEIS A000044).
// URL: https://oeis.org/A000044

// GetA000044Sequence returns the A000044 sequence.
func GetA000044Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000044AtPosition(maxNumber)
	}

	return GenerateA000044Sequence(maxNumber)
}

// GenerateA000044Sequence generates A000044 from a(0) through a(maxNumber).
func GenerateA000044Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := maxNumber.Int64()
	sequence := generateA000044Terms(limit)

	return &NumericSequence{
		Name:     "Dying rabbits (A000044)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[limit],
	}, nil
}

// GetA000044AtPosition returns the n-th term of A000044.
func GetA000044AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	sequence := generateA000044Terms(n.Int64())
	result := sequence[n.Int64()]

	return &NumericSequence{
		Name:     "Dying rabbits (A000044)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// generateA000044Terms computes terms a(0) through a(limit) for OEIS A000044.
func generateA000044Terms(limit int64) []*big.Int {
	sequence := make([]*big.Int, limit+1)
	sequence[0] = big.NewInt(1)

	if limit == 0 {
		return sequence
	}

	sequence[1] = big.NewInt(1)
	if limit >= 2 {
		sequence[2] = big.NewInt(1)
	}

	for i := int64(3); i <= limit && i <= 12; i++ {
		sequence[i] = new(big.Int).Add(sequence[i-1], sequence[i-2])
	}

	for i := int64(13); i <= limit; i++ {
		sequence[i] = new(big.Int).Add(sequence[i-1], sequence[i-2])
		sequence[i].Sub(sequence[i], sequence[i-13])
	}

	return sequence
}
