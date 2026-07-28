package sequences

import (
	"fmt"
	"math/big"
)

// A000022: Number of centered hydrocarbons with n atoms.
// URL: https://oeis.org/A000022
// Calculated using generating functions for rooted ternary trees (A000598).
// Centered hydrocarbons are alkanes whose carbon tree has a centroid.
// A carbon tree is a tree where each node has degree at most 4.
// For centered hydrocarbons, there is a unique node (centroid) such that
// no component of the tree remaining after its removal has more than n/2 nodes.

// GetA000022Sequence returns the number of centered hydrocarbons sequence (OEIS A000022).
func GetA000022Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000022AtPosition(maxNumber)
	}
	return GenerateA000022Sequence(maxNumber)
}

// GenerateA000022Sequence generates the A000022 sequence up to maxNumber.
func GenerateA000022Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := int(maxNumber.Int64())
	if !maxNumber.IsInt64() || limit > 500 { // Safety limit
		limit = 500
	}

	results := calculateA000022(limit)
	sequence := make([]*big.Int, limit+1)
	for i := 0; i <= limit; i++ {
		sequence[i] = results[i]
	}

	return &NumericSequence{
		Name:     "Centered Hydrocarbons (A000022)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[len(sequence)-1],
	}, nil
}

// GetA000022AtPosition returns the n-th term of A000022.
func GetA000022AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	limit := int(n.Int64())
	if !n.IsInt64() || limit > 500 {
		limit = 500
	}

	results := calculateA000022(limit)
	var result *big.Int
	if int(n.Int64()) < len(results) {
		result = results[int(n.Int64())]
	} else {
		return nil, fmt.Errorf("position %s exceeds calculated limit", n.String())
	}

	return &NumericSequence{
		Name:     "Centered Hydrocarbons (A000022)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateA000022 computes the first n+1 terms of A000022 using the Rains and Sloane algorithm.
// S3[f, x] := (f(x)^3 + 3*f(x)*f(x^2) + 2*f(x^3)) / 6
// S4[f, x] := (f(x)^4 + 6*f(x)^2*f(x^2) + 8*f(x)*f(x^3) + 3*f(x^2)^2 + 6*f(x^4)) / 24
// T[h, x] is the g.f. for rooted ternary trees of height at most h.
// T[-1, x] = 1
// T[h, x] = 1 + x * S3[T[h-1], x]
// The g.f. for centered hydrocarbons (A000022) is:
// Sum_{h >= 0} (x * S4[T[h], x] - x * S4[T[h-1], x] - (T[h](x) - T[h-1](x)) * (T[h](x) - 1)) + x
func calculateA000022(n int) []*big.Int {
	if n < 0 {
		return nil
	}

	// T[h] stores coefficients of x^k for h-th height limit
	// We need height up to n/2 approx, but let's just go up to n.
	maxH := n/2 + 1
	if maxH < 1 {
		maxH = 1
	}

	t := make([][]*big.Int, maxH+1)
	for h := range t {
		t[h] = make([]*big.Int, n+1)
		for k := range t[h] {
			t[h][k] = big.NewInt(0)
		}
	}

	// T[-1, x] = 1 (effectively t[0] here as h ranges from -1 to maxH-1)
	// Let's use index h directly.
	// t[h] will correspond to height h-1 in the formula, so t[0] is T[-1, x]
	t[0][0] = big.NewInt(1)

	for h := 1; h <= maxH; h++ {
		// T[h-1, x] = 1 + x * S3[T[h-2], x]
		// Formula uses T[-1] = 1, so T[0] = 1 + x*S3[1, x] = 1 + x*(1+3+2)/6 = 1 + x
		t[h][0] = big.NewInt(1)
		prevT := t[h-1]

		// Coeff of x^k in S3[prevT, x]
		for k := 0; k < n; k++ {
			// S3 = (T^3 + 3*T(x)*T(x^2) + 2*T(x^3)) / 6
			
			// T(x)^3
			cB3 := big.NewInt(0)
			for i := 0; i <= k; i++ {
				for j := 0; j <= k-i; j++ {
					term := new(big.Int).Mul(prevT[i], prevT[j])
					term.Mul(term, prevT[k-i-j])
					cB3.Add(cB3, term)
				}
			}

			// 3*T(x)*T(x^2)
			cB1B2 := big.NewInt(0)
			for i := 0; i <= k; i++ {
				rem := k - i
				if rem%2 == 0 {
					term := new(big.Int).Mul(prevT[i], prevT[rem/2])
					cB1B2.Add(cB1B2, term)
				}
			}
			cB1B2.Mul(cB1B2, big.NewInt(3))

			// 2*T(x^3)
			cB3only := big.NewInt(0)
			if k%3 == 0 {
				cB3only.Set(prevT[k/3])
			}
			cB3only.Mul(cB3only, big.NewInt(2))

			sum := new(big.Int).Add(cB3, cB1B2)
			sum.Add(sum, cB3only)
			
			// Coeff of x^{k+1} in T[h-1, x]
			t[h][k+1].Div(sum, big.NewInt(6))
		}
	}

	res := make([]*big.Int, n+1)
	for i := range res {
		res[i] = big.NewInt(0)
	}

	// Sum_{h = 1 to maxH} (x * S4[T[h-1], x] - x * S4[T[h-2], x] - (T[h-1](x) - T[h-2](x)) * (T[h-1](x) - 1))
	for h := 1; h <= maxH; h++ {
		th1 := t[h]
		th2 := t[h-1]

		// x * S4[th1, x]
		s4_1 := calculateS4(th1, n)
		// x * S4[th2, x]
		s4_2 := calculateS4(th2, n)

		// th1 - th2
		diffT := make([]*big.Int, n+1)
		for i := 0; i <= n; i++ {
			diffT[i] = new(big.Int).Sub(th1[i], th2[i])
		}

		// th1 - 1
		th1m1 := make([]*big.Int, n+1)
		for i := 0; i <= n; i++ {
			th1m1[i] = new(big.Int).Set(th1[i])
		}
		th1m1[0].Sub(th1m1[0], big.NewInt(1))

		// (th1 - th2) * (th1 - 1)
		prod := make([]*big.Int, n+1)
		for i := 0; i <= n; i++ {
			prod[i] = big.NewInt(0)
			for j := 0; j <= i; j++ {
				term := new(big.Int).Mul(diffT[j], th1m1[i-j])
				prod[i].Add(prod[i], term)
			}
		}

		for i := 0; i <= n; i++ {
			term := big.NewInt(0)
			if i > 0 {
				term.Sub(s4_1[i-1], s4_2[i-1])
			}
			term.Sub(term, prod[i])
			res[i].Add(res[i], term)
		}
	}

	if n >= 1 {
		res[1].Add(res[1], big.NewInt(1))
	}

	return res
}

func calculateS4(t []*big.Int, n int) []*big.Int {
	s4 := make([]*big.Int, n+1)
	for k := 0; k <= n; k++ {
		// S4 = (T^4 + 6*T^2*T(x^2) + 8*T*T(x^3) + 3*T(x^2)^2 + 6*T(x^4)) / 24

		// T(x)^4
		cB4 := big.NewInt(0)
		for i := 0; i <= k; i++ {
			for j := 0; j <= k-i; j++ {
				for l := 0; l <= k-i-j; l++ {
					term := new(big.Int).Mul(t[i], t[j])
					term.Mul(term, t[l])
					term.Mul(term, t[k-i-j-l])
					cB4.Add(cB4, term)
				}
			}
		}

		// 6*T(x)^2*T(x^2)
		cB2B2 := big.NewInt(0)
		for i := 0; i <= k; i++ {
			for j := 0; j <= k-i; j++ {
				rem := k - i - j
				if rem%2 == 0 {
					term := new(big.Int).Mul(t[i], t[j])
					term.Mul(term, t[rem/2])
					cB2B2.Add(cB2B2, term)
				}
			}
		}
		cB2B2.Mul(cB2B2, big.NewInt(6))

		// 8*T(x)*T(x^3)
		cB1B3 := big.NewInt(0)
		for i := 0; i <= k; i++ {
			rem := k - i
			if rem%3 == 0 {
				term := new(big.Int).Mul(t[i], t[rem/3])
				cB1B3.Add(cB1B3, term)
			}
		}
		cB1B3.Mul(cB1B3, big.NewInt(8))

		// 3*T(x^2)^2
		cB2_2 := big.NewInt(0)
		for i := 0; i <= k; i += 2 {
			rem := k - i
			if rem%2 == 0 {
				term := new(big.Int).Mul(t[i/2], t[rem/2])
				cB2_2.Add(cB2_2, term)
			}
		}
		cB2_2.Mul(cB2_2, big.NewInt(3))

		// 6*T(x^4)
		cB4only := big.NewInt(0)
		if k%4 == 0 {
			cB4only.Set(t[k/4])
		}
		cB4only.Mul(cB4only, big.NewInt(6))

		sum := new(big.Int).Add(cB4, cB2B2)
		sum.Add(sum, cB1B3)
		sum.Add(sum, cB2_2)
		sum.Add(sum, cB4only)

		s4[k] = new(big.Int).Div(sum, big.NewInt(24))
	}
	return s4
}
