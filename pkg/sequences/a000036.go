package sequences

import (
	"fmt"
	"math"
	"math/big"
)

// A000036 (OEIS): closest integer to P(A000099(n)), where
// A(m) = #{(i,j): i^2 + j^2 <= m}, V(m) = Pi*m, and P(m) = A(m) - V(m).
// URL: https://oeis.org/A000036

// GetA000036Sequence returns sequence A000036 values.
// If isPositional is true, maxNumber is treated as a 1-based position n.
// If isPositional is false, maxNumber is treated as the number of terms to generate.
func GetA000036Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000036AtPosition(maxNumber)
	}
	return GenerateA000036Sequence(maxNumber)
}

// GenerateA000036Sequence generates the first maxNumber terms of A000036.
func GenerateA000036Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Record values of |P(n)| (A000036)",
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
		term, err := a000036Term(n)
		if err != nil {
			return nil, err
		}
		sequence = append(sequence, term)
	}

	result := big.NewInt(0)
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Record values of |P(n)| (A000036)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetA000036AtPosition returns the n-th term of A000036, with n starting at 1.
func GetA000036AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() <= 0 {
		return nil, fmt.Errorf("position must be positive")
	}
	if !n.IsInt64() {
		return nil, fmt.Errorf("position too large for current implementation")
	}

	term, err := a000036Term(n.Int64())
	if err != nil {
		return nil, err
	}

	return &NumericSequence{
		Name:     "Record values of |P(n)| (A000036)",
		Number:   n,
		Sequence: []*big.Int{term},
		Result:   term,
	}, nil
}

func a000036Term(index int64) (*big.Int, error) {
	a000099, ok := oeisLookups["A000099"]
	if !ok {
		return nil, fmt.Errorf("lookup for A000099 not found")
	}
	if index < 1 || index > int64(len(a000099)) {
		return nil, fmt.Errorf("position out of lookup range")
	}

	recordN, ok := new(big.Int).SetString(a000099[index-1], 10)
	if !ok || !recordN.IsInt64() {
		return nil, fmt.Errorf("invalid A000099 term at position %d", index)
	}

	p := latticePointDiscrepancy(recordN.Int64())
	return big.NewInt(int64(math.Round(p))), nil
}

func latticePointDiscrepancy(n int64) float64 {
	a := latticePointCount(n)
	return float64(a) - math.Pi*float64(n)
}

func latticePointCount(n int64) int64 {
	r := int64(math.Sqrt(float64(n)))
	var count int64
	for x := -r; x <= r; x++ {
		yMax := int64(math.Sqrt(float64(n - x*x)))
		count += 2*yMax + 1
	}
	return count
}
