package sequences

import (
	"fmt"
	"math/big"
)

// Modular function j (A000521)
// URL: https://oeis.org/A000521
// Description: Coefficients of modular function j as power series in q = e^(2 Pi i t).
// j(q) = 1/q + 744 + 196884q + 21493760q^2 + ...
// The implementation uses the relation j(q) * Delta(q) = E_4(q)^3, where Delta(q) is the modular discriminant
// and E_4(q) is the Eisenstein series of weight 4.

// GetModularJCoefficientsSequence returns the modular function j coefficients (OEIS A000521).
func GetModularJCoefficientsSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetModularJCoefficientAtPosition(maxNumber)
	}
	return GenerateModularJCoefficientsSequence(maxNumber)
}

// GenerateModularJCoefficientsSequence generates the A000521 sequence up to position n.
func GenerateModularJCoefficientsSequence(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < -1 {
		return &NumericSequence{
			Name:     "Modular function j (A000521)",
			Number:   n,
			Sequence: []*big.Int{},
		}, nil
	}

	limit := int(n.Int64())
	if !n.IsInt64() || limit > 100 { // Safety limit for computational reasons
		limit = 100
	}

	coeffs := calculateModularJCoefficients(limit)
	sequence := make([]*big.Int, len(coeffs))
	copy(sequence, coeffs)

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Modular function j (A000521)",
		Number:   n,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetModularJCoefficientAtPosition returns the n-th coefficient of A000521.
func GetModularJCoefficientAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < -1 {
		return nil, fmt.Errorf("position must be at least -1")
	}

	limit := int(n.Int64())
	if !n.IsInt64() || limit > 500 {
		limit = 500
	}

	coeffs := calculateModularJCoefficients(limit)
	idx := int(n.Int64()) + 1 // offset is -1
	var val *big.Int
	if idx >= 0 && idx < len(coeffs) {
		val = coeffs[idx]
	} else {
		return nil, fmt.Errorf("position %s exceeds limit", n.String())
	}

	return &NumericSequence{
		Name:     "Modular function j (A000521)",
		Number:   n,
		Sequence: []*big.Int{val},
		Result:   val,
	}, nil
}

// calculateModularJCoefficients computes coefficients c(n) for n = -1 to limit.
func calculateModularJCoefficients(limit int) []*big.Int {
	if limit < -1 {
		return nil
	}

	// We need Ramanujan tau(k) for k = 1 to limit+1
	tau := calculateRamanujanTau(limit + 2) // We need tau[1] to tau[limit+2]

	// We need Eisenstein series E4(q) coefficients P_k for E4(q)^3
	// E4(q) = 1 + 240 * Sum_{n=1..inf} sigma_3(n) q^n
	e4 := make([]*big.Int, limit+2)
	e4[0] = big.NewInt(1)
	for k := 1; k <= limit+1; k++ {
		s3 := sigma3(k)
		e4[k] = new(big.Int).Mul(big.NewInt(240), s3)
	}

	// E4(q)^3
	p := powerSeries(e4, 3, limit+1)

	// Recurrence: sum_{i=-1}^{k-1} c(i) * tau(k-i) = P_k
	// c(k-1) * tau(1) = P_k - sum_{i=-1}^{k-2} c(i) * tau(k-i)
	// Since tau(1) = 1:
	// c(k-1) = P_k - sum_{i=-1}^{k-2} c(i) * tau(k-i)

	c := make([]*big.Int, limit+2) // c[0] is c(-1), c[1] is c(0), etc.
	for k := 0; k <= limit+1; k++ {
		// Calculate c(k-1)
		sum := big.NewInt(0)
		for i := -1; i <= k-2; i++ {
			term := new(big.Int).Mul(c[i+1], tau[k-i])
			sum.Add(sum, term)
		}
		c[k] = new(big.Int).Sub(p[k], sum)
	}

	return c
}

// sigma3 computes the sum of cubes of divisors of n.
func sigma3(n int) *big.Int {
	sum := big.NewInt(0)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			// i is a divisor
			val := new(big.Int).SetInt64(int64(i))
			sum.Add(sum, new(big.Int).Exp(val, big.NewInt(3), nil))
			if i*i != n {
				// n/i is another divisor
				val2 := new(big.Int).SetInt64(int64(n / i))
				sum.Add(sum, new(big.Int).Exp(val2, big.NewInt(3), nil))
			}
		}
	}
	return sum
}
