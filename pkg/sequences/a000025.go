package sequences

import (
	"fmt"
	"math/big"
)

// A000025: Coefficients of the 3rd-order mock theta function f(q).
// URL: https://oeis.org/A000025
// Generating function: f(q) = 1 + Sum_{n>=1} (q^(n^2) / Product_{i=1..n} (1 + q^i)^2)
// Formerly M0433 N0164

// GetA000025Sequence returns the coefficients of the 3rd-order mock theta function f(q) (OEIS A000025).
func GetA000025Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000025AtPosition(maxNumber)
	}
	return GenerateA000025Sequence(maxNumber)
}

// GenerateA000025Sequence generates the A000025 sequence up to maxNumber.
func GenerateA000025Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := int(maxNumber.Int64())
	// Use a reasonable limit for calculation
	if !maxNumber.IsInt64() || limit > 1000 {
		limit = 1000
	}

	results := calculateA000025(limit)
	sequence := make([]*big.Int, limit+1)
	for i := 0; i <= limit; i++ {
		sequence[i] = results[i]
	}

	return &NumericSequence{
		Name:     "Mock Theta Function f(q) (A000025)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[limit],
	}, nil
}

// GetA000025AtPosition returns the n-th term of A000025.
func GetA000025AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	limit := int(n.Int64())
	if !n.IsInt64() || limit > 1000 {
		limit = 1000
	}

	results := calculateA000025(limit)
	var result *big.Int
	if int(n.Int64()) < len(results) {
		result = results[int(n.Int64())]
	} else {
		// If exceeds limit, we would need a more efficient way or just return error
		return nil, fmt.Errorf("position %s exceeds calculated limit", n.String())
	}

	return &NumericSequence{
		Name:     "Mock Theta Function f(q) (A000025)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateA000025 computes the first n+1 terms of A000025.
// f(q) = 1 + Sum_{k>=1} (q^(k^2) / Product_{i=1..k} (1 + q^i)^2)
func calculateA000025(n int) []*big.Int {
	coeffs := make([]*big.Int, n+1)
	for i := range coeffs {
		coeffs[i] = big.NewInt(0)
	}
	// Initial 1
	coeffs[0] = big.NewInt(1)

	for k := 1; k*k <= n; k++ {
		k2 := k * k
		// Compute 1 / Product_{i=1..k} (1 + q^i)^2 up to q^(n-k2)
		// Let G_k(q) = Product_{i=1..k} (1 + q^i)
		// We want coeffs of q^m in 1/G_k(q)^2
		
		termCoeffs := calculateInverseProductSquared(k, n-k2)
		
		// Add q^k2 * termCoeffs to coeffs
		for m := 0; m <= n-k2; m++ {
			coeffs[m+k2].Add(coeffs[m+k2], termCoeffs[m])
		}
	}

	return coeffs
}

// calculateInverseProductSquared computes the coefficients of 1 / Product_{i=1..k} (1 + q^i)^2 up to q^limit.
func calculateInverseProductSquared(k int, limit int) []*big.Int {
	// 1 / Product_{i=1..k} (1 + q^i)^2 = Product_{i=1..k} (1 + q^i)^{-2}
	// (1 + q^i)^{-2} = Sum_{j>=0} (-1)^j * (j+1) * q^(i*j)
	
	coeffs := make([]*big.Int, limit+1)
	for i := range coeffs {
		coeffs[i] = big.NewInt(0)
	}
	coeffs[0] = big.NewInt(1)

	for i := 1; i <= k; i++ {
		nextCoeffs := make([]*big.Int, limit+1)
		for j := range nextCoeffs {
			nextCoeffs[j] = big.NewInt(0)
		}
		
		// Multiply current coeffs by (1 + q^i)^{-2} = 1 - 2q^i + 3q^{2i} - 4q^{3i} + ...
		for m := 0; m <= limit; m++ {
			if coeffs[m].Sign() == 0 {
				continue
			}
			for j := 0; i*j <= limit-m; j++ {
				// term = (-1)^j * (j+1)
				term := big.NewInt(int64(j + 1))
				if j%2 != 0 {
					term.Neg(term)
				}
				
				val := new(big.Int).Mul(coeffs[m], term)
				nextCoeffs[m+i*j].Add(nextCoeffs[m+i*j], val)
			}
		}
		coeffs = nextCoeffs
	}
	
	return coeffs
}
