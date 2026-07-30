package sequences

import (
	"fmt"
	"math"
	"math/big"
)

// A000037 (OEIS): numbers that are not squares (the nonsquares).
// URL: https://oeis.org/A000037

// GetA000037Sequence returns sequence A000037 values.
// If isPositional is true, maxNumber is treated as a 1-based position n.
// If isPositional is false, maxNumber is treated as the number of terms to generate.
func GetA000037Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000037AtPosition(maxNumber)
	}
	return GenerateA000037Sequence(maxNumber)
}

// GenerateA000037Sequence generates the first maxNumber terms of A000037.
func GenerateA000037Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Nonsquares (A000037)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}
	if !maxNumber.IsInt64() {
		return nil, fmt.Errorf("maxNumber too large for current implementation")
	}

	limit := maxNumber.Int64()
	sequence := make([]*big.Int, 0, limit)
	for n := int64(1); n <= limit; n++ {
		sequence = append(sequence, a000037Term(n))
	}

	result := big.NewInt(0)
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Nonsquares (A000037)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetA000037AtPosition returns the n-th term of A000037, with n starting at 1.
func GetA000037AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() <= 0 {
		return nil, fmt.Errorf("position must be positive")
	}
	if !n.IsInt64() {
		return nil, fmt.Errorf("position too large for current implementation")
	}

	term := a000037Term(n.Int64())
	return &NumericSequence{
		Name:     "Nonsquares (A000037)",
		Number:   n,
		Sequence: []*big.Int{term},
		Result:   term,
	}, nil
}

// a000037Term computes a(n) = n + floor(1/2 + sqrt(n)) for n >= 1.
func a000037Term(n int64) *big.Int {
	return big.NewInt(n + int64(math.Floor(0.5+math.Sqrt(float64(n)))))
}
