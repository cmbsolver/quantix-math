package sequences

import (
	"fmt"
	"math/big"
)

// A000602: Number of n-carbon alkanes C(n)H(2n+2) ignoring stereoisomers.
// URL: https://oeis.org/A000602
// Calculated using generating functions for rooted ternary trees (A000598).

// GetAlkanesSequence returns the number of alkanes sequence (OEIS A000602).
func GetAlkanesSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetAlkanesAtPosition(maxNumber)
	}
	return GenerateAlkanesSequence(maxNumber)
}

// GenerateAlkanesSequence generates the A000602 sequence up to maxNumber.
func GenerateAlkanesSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := int(maxNumber.Int64())
	if !maxNumber.IsInt64() || limit > 500 { // Safety limit
		limit = 500
	}

	results := calculateAlkanes(limit)
	sequence := make([]*big.Int, limit+1)
	for i := 0; i <= limit; i++ {
		sequence[i] = results[i]
	}

	return &NumericSequence{
		Name:     "Number of Alkanes (A000602)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[len(sequence)-1],
	}, nil
}

// GetAlkanesAtPosition returns the n-th term of A000602.
func GetAlkanesAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	limit := int(n.Int64())
	if !n.IsInt64() || limit > 500 {
		limit = 500
	}

	results := calculateAlkanes(limit)
	var result *big.Int
	if int(n.Int64()) < len(results) {
		result = results[int(n.Int64())]
	} else {
		return nil, fmt.Errorf("position %s exceeds calculated limit", n.String())
	}

	return &NumericSequence{
		Name:     "Number of Alkanes (A000602)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateAlkanes computes the first n+1 terms of A000602 using g.f.
// B(x) = 1 + (x/6)*(B(x)^3 + 3*B(x)*B(x^2) + 2*B(x^3))
// A(x) = B(x) - (B(x)^2 - B(x^2))/2 + (x/24)*(B(x)^4 + 6*B(x)^2*B(x^2) + 8*B(x)*B(x^3) + 3*B(x^2)^2 + 6*B(x^4))
func calculateAlkanes(n int) []*big.Int {
	// Rooted ternary trees B(x) (A000598)
	b := make([]*big.Int, n+1)
	for i := range b {
		b[i] = big.NewInt(0)
	}
	b[0] = big.NewInt(1)

	// Solve B(x) iteratively
	for i := 1; i <= n; i++ {
		// B(x) = 1 + x * S3(B, x)
		// Coeff of x^i in B(x) is coeff of x^{i-1} in S3(B, x)
		// S3(B, x) = (B(x)^3 + 3*B(x)*B(x^2) + 2*B(x^3)) / 6

		targetIdx := i - 1

		// Coeff of x^targetIdx in B(x)^3
		cB3 := big.NewInt(0)
		for j := 0; j <= targetIdx; j++ {
			for k := 0; k <= targetIdx-j; k++ {
				term := new(big.Int).Mul(b[j], b[k])
				term.Mul(term, b[targetIdx-j-k])
				cB3.Add(cB3, term)
			}
		}

		// Coeff of x^targetIdx in 3*B(x)*B(x^2)
		cB1B2 := big.NewInt(0)
		for j := 0; j <= targetIdx; j++ {
			// B(x^2) has non-zero coeffs only at even indices
			k2 := targetIdx - j
			if k2%2 == 0 {
				k := k2 / 2
				if k < i {
					term := new(big.Int).Mul(b[j], b[k])
					cB1B2.Add(cB1B2, term)
				}
			}
		}
		cB1B2.Mul(cB1B2, big.NewInt(3))

		// Coeff of x^targetIdx in 2*B(x^3)
		cB3only := big.NewInt(0)
		if targetIdx%3 == 0 {
			k := targetIdx / 3
			if k < i {
				cB3only.Set(b[k])
			}
		}
		cB3only.Mul(cB3only, big.NewInt(2))

		sum := new(big.Int).Add(cB3, cB1B2)
		sum.Add(sum, cB3only)
		b[i].Div(sum, big.NewInt(6))
	}

	// Now calculate A(x)
	// A(x) = B(x) - (B(x)^2 - B(x^2))/2 + (x/24)*(B(x)^4 + 6*B(x)^2*B(x^2) + 8*B(x)*B(x^3) + 3*B(x^2)^2 + 6*B(x^4))
	a := make([]*big.Int, n+1)
	for i := 0; i <= n; i++ {
		term1 := b[i]

		// Coeff of x^i in (B(x)^2 - B(x^2))/2
		cB2 := big.NewInt(0)
		for j := 0; j <= i; j++ {
			term := new(big.Int).Mul(b[j], b[i-j])
			cB2.Add(cB2, term)
		}
		if i%2 == 0 {
			cB2.Sub(cB2, b[i/2])
		}
		term2 := new(big.Int).Div(cB2, big.NewInt(2))

		// Coeff of x^i in x*S4(B, x)
		// Coeff of x^{i-1} in S4(B, x)
		term3 := big.NewInt(0)
		if i > 0 {
			targetIdx := i - 1

			// B(x)^4
			cB4 := big.NewInt(0)
			for j := 0; j <= targetIdx; j++ {
				for k := 0; k <= targetIdx-j; k++ {
					for l := 0; l <= targetIdx-j-k; l++ {
						term := new(big.Int).Mul(b[j], b[k])
						term.Mul(term, b[l])
						term.Mul(term, b[targetIdx-j-k-l])
						cB4.Add(cB4, term)
					}
				}
			}

			// 6*B(x)^2*B(x^2)
			cB2B2 := big.NewInt(0)
			for j := 0; j <= targetIdx; j++ {
				for k := 0; k <= targetIdx-j; k++ {
					l2 := targetIdx - j - k
					if l2%2 == 0 {
						l := l2 / 2
						term := new(big.Int).Mul(b[j], b[k])
						term.Mul(term, b[l])
						cB2B2.Add(cB2B2, term)
					}
				}
			}
			cB2B2.Mul(cB2B2, big.NewInt(6))

			// 8*B(x)*B(x^3)
			cB1B3 := big.NewInt(0)
			for j := 0; j <= targetIdx; j++ {
				k3 := targetIdx - j
				if k3%3 == 0 {
					k := k3 / 3
					term := new(big.Int).Mul(b[j], b[k])
					cB1B3.Add(cB1B3, term)
				}
			}
			cB1B3.Mul(cB1B3, big.NewInt(8))

			// 3*B(x^2)^2
			cB2_2 := big.NewInt(0)
			for j2 := 0; j2 <= targetIdx; j2 += 2 {
				k2 := targetIdx - j2
				if k2%2 == 0 {
					term := new(big.Int).Mul(b[j2/2], b[k2/2])
					cB2_2.Add(cB2_2, term)
				}
			}
			cB2_2.Mul(cB2_2, big.NewInt(3))

			// 6*B(x^4)
			cB4only := big.NewInt(0)
			if targetIdx%4 == 0 {
				cB4only.Set(b[targetIdx/4])
			}
			cB4only.Mul(cB4only, big.NewInt(6))

			sumS4 := new(big.Int).Add(cB4, cB2B2)
			sumS4.Add(sumS4, cB1B3)
			sumS4.Add(sumS4, cB2_2)
			sumS4.Add(sumS4, cB4only)

			term3.Div(sumS4, big.NewInt(24))
		}

		res := new(big.Int).Sub(term1, term2)
		res.Add(res, term3)
		a[i] = res
	}

	return a
}
