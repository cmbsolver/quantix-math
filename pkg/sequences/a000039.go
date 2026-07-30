package sequences

import (
	"fmt"
	"math/big"
)

// A000039 (OEIS): Coefficient of q^(2n) in the series expansion of Ramanujan's mock theta function f(q).
// URL: https://oeis.org/A000039
// Definition used: f(q) = 1 + Sum_{m>=1} q^(m^2) / Product_{k=1..m} (1+q^k)^2.

// GetA000039Sequence returns sequence A000039 values.
// If isPositional is true, maxNumber is treated as an index n (offset 0).
// If isPositional is false, maxNumber is treated as the number of terms to generate from n=0.
func GetA000039Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000039AtPosition(maxNumber)
	}
	return GenerateA000039Sequence(maxNumber)
}

// GenerateA000039Sequence generates the first maxNumber terms of A000039.
func GenerateA000039Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Mock theta f(q) even-power coefficients (A000039)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}
	if !maxNumber.IsInt64() {
		return nil, fmt.Errorf("maxNumber too large for current implementation")
	}

	termCount := maxNumber.Int64()
	if termCount == 0 {
		return &NumericSequence{
			Name:     "Mock theta f(q) even-power coefficients (A000039)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   big.NewInt(0),
		}, nil
	}

	maxDegree := int(2 * termCount)
	coeffs := a000039MockThetaFCoefficients(maxDegree)

	sequence := make([]*big.Int, 0, termCount)
	for n := int64(0); n < termCount; n++ {
		sequence = append(sequence, new(big.Int).Set(coeffs[2*n]))
	}

	result := big.NewInt(0)
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Mock theta f(q) even-power coefficients (A000039)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetA000039AtPosition returns a(n) for A000039 where n >= 0.
func GetA000039AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be non-negative")
	}
	if !n.IsInt64() {
		return nil, fmt.Errorf("position too large for current implementation")
	}

	index := n.Int64()
	coeffs := a000039MockThetaFCoefficients(int(2 * (index + 1)))
	term := new(big.Int).Set(coeffs[2*index])

	return &NumericSequence{
		Name:     "Mock theta f(q) even-power coefficients (A000039)",
		Number:   n,
		Sequence: []*big.Int{term},
		Result:   term,
	}, nil
}

// a000039MockThetaFCoefficients computes coefficients of f(q) up to q^maxDegree.
func a000039MockThetaFCoefficients(maxDegree int) []*big.Int {
	coeffs := make([]*big.Int, maxDegree+1)
	for i := range coeffs {
		coeffs[i] = big.NewInt(0)
	}
	coeffs[0] = big.NewInt(1)

	for m := 1; m*m <= maxDegree; m++ {
		shift := m * m
		remaining := maxDegree - shift

		term := make([]*big.Int, remaining+1)
		for i := range term {
			term[i] = big.NewInt(0)
		}
		term[0] = big.NewInt(1)

		for k := 1; k <= m; k++ {
			next := make([]*big.Int, remaining+1)
			for i := range next {
				next[i] = big.NewInt(0)
			}
			for degree := 0; degree <= remaining; degree++ {
				if term[degree].Sign() == 0 {
					continue
				}
				for mul := 0; degree+k*mul <= remaining; mul++ {
					factor := big.NewInt(int64(mul + 1))
					if mul%2 == 1 {
						factor.Neg(factor)
					}
					contrib := new(big.Int).Mul(term[degree], factor)
					next[degree+k*mul].Add(next[degree+k*mul], contrib)
				}
			}
			term = next
		}

		for degree := 0; degree <= remaining; degree++ {
			coeffs[shift+degree].Add(coeffs[shift+degree], term[degree])
		}
	}

	return coeffs
}
