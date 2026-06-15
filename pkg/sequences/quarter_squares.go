package sequences

import (
	"fmt"
	"math/big"
)

// Quarter-squares (OEIS A002620)
// URL: https://oeis.org/A002620
// a(n) = floor(n^2/4).

// GetQuarterSquaresSequence returns the quarter-squares sequence (OEIS A002620).
func GetQuarterSquaresSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetQuarterSquaresAtPosition(maxNumber)
	}
	return GenerateQuarterSquaresSequence(maxNumber)
}

// GenerateQuarterSquaresSequence generates the A002620 sequence up to maxNumber (a(0), a(1), ..., a(maxNumber)).
func GenerateQuarterSquaresSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		sequence[i] = calculateQuarterSquare(int64(i))
	}

	result := sequence[n]

	return &NumericSequence{
		Name:     "Quarter-squares (A002620)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetQuarterSquaresAtPosition returns the n-th term of A002620 (n >= 0).
func GetQuarterSquaresAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	result := calculateQuarterSquare(n.Int64())

	return &NumericSequence{
		Name:     "Quarter-squares (A002620)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateQuarterSquare calculates floor(n^2/4).
func calculateQuarterSquare(n int64) *big.Int {
	bigN := big.NewInt(n)
	nSquared := new(big.Int).Mul(bigN, bigN)
	return new(big.Int).Div(nSquared, big.NewInt(4))
}
