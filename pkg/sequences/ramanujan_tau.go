package sequences

import (
	"fmt"
	"math/big"
)

// Ramanujan's tau function (A000594)
// URL: https://oeis.org/A000594
// Definition: Sum_{n=1..inf} tau(n)q^n = q * Product_{n=1..inf} (1-q^n)^24

// GetRamanujanTauSequence returns Ramanujan's tau sequence (OEIS A000594).
func GetRamanujanTauSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetRamanujanTauAtPosition(maxNumber)
	}
	return GenerateRamanujanTauSequence(maxNumber)
}

// GenerateRamanujanTauSequence generates the A000594 sequence up to maxNumber.
func GenerateRamanujanTauSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(1)) < 0 {
		return &NumericSequence{
			Name:     "Ramanujan's tau function (A000594)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}

	limit := int(maxNumber.Int64())
	if !maxNumber.IsInt64() || limit > 1000 { // Safety limit
		limit = 1000
	}

	results := calculateRamanujanTau(limit)
	sequence := make([]*big.Int, len(results)-1)
	for i := 1; i <= limit; i++ {
		sequence[i-1] = results[i]
	}

	return &NumericSequence{
		Name:     "Ramanujan's tau function (A000594)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[len(sequence)-1],
	}, nil
}

// GetRamanujanTauAtPosition returns the n-th term of A000594.
func GetRamanujanTauAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("position must be at least 1")
	}

	limit := int(n.Int64())
	if !n.IsInt64() || limit > 1000 {
		limit = 1000
	}

	results := calculateRamanujanTau(limit)
	var result *big.Int
	if int(n.Int64()) < len(results) {
		result = results[int(n.Int64())]
	} else {
		return nil, fmt.Errorf("position %s exceeds calculated limit", n.String())
	}

	return &NumericSequence{
		Name:     "Ramanujan's tau function (A000594)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateRamanujanTau computes the first n terms of A000594.
// It uses Euler's pentagonal number theorem: Product_{k=1..inf} (1-q^k) = Sum_{k=-inf..inf} (-1)^k q^(k(3k-1)/2)
// Then it raises this power series to the 24th power and multiplies by q.
func calculateRamanujanTau(n int) []*big.Int {
	if n < 1 {
		return []*big.Int{big.NewInt(0)}
	}

	// Coefficients of P(q) = Product_{k=1..inf} (1-q^k) up to q^n
	p := make([]*big.Int, n+1)
	for i := range p {
		p[i] = big.NewInt(0)
	}
	p[0] = big.NewInt(1)

	for k := 1; ; k++ {
		p1 := k * (3*k - 1) / 2
		p2 := k * (3*k + 1) / 2
		if p1 > n && p2 > n {
			break
		}
		coeff := big.NewInt(1)
		if k%2 != 0 {
			coeff = big.NewInt(-1)
		}
		if p1 <= n {
			p[p1].Set(coeff)
		}
		if p2 <= n {
			p[p2].Set(coeff)
		}
	}

	// Calculate P(q)^24
	// We can use the property (P(q)^2)^2... or just repeated multiplication.
	// Since 24 = 16 + 8, or just power by squaring.
	res := powerSeries(p, 24, n)

	// Ramanujan tau starts from q^1, so tau(n) is the coefficient of q^(n-1) in P(q)^24
	// or coefficient of q^n in q*P(q)^24.
	tau := make([]*big.Int, n+1)
	tau[0] = big.NewInt(0)
	for i := 1; i <= n; i++ {
		tau[i] = new(big.Int).Set(res[i-1])
	}

	return tau
}

// powerSeries calculates (s(q))^p mod q^(n+1)
func powerSeries(s []*big.Int, p int, n int) []*big.Int {
	res := make([]*big.Int, n+1)
	for i := range res {
		res[i] = big.NewInt(0)
	}
	res[0] = big.NewInt(1)

	base := make([]*big.Int, n+1)
	for i := range base {
		base[i] = new(big.Int).Set(s[i])
	}

	exp := p
	for exp > 0 {
		if exp%2 == 1 {
			res = multiplySeries(res, base, n)
		}
		base = multiplySeries(base, base, n)
		exp /= 2
	}

	return res
}

// multiplySeries multiplies two power series mod q^(n+1)
func multiplySeries(s1, s2 []*big.Int, n int) []*big.Int {
	res := make([]*big.Int, n+1)
	for i := range res {
		res[i] = big.NewInt(0)
	}

	for i := 0; i <= n; i++ {
		if s1[i].Sign() == 0 {
			continue
		}
		for j := 0; j <= n-i; j++ {
			if s2[j].Sign() == 0 {
				continue
			}
			term := new(big.Int).Mul(s1[i], s2[j])
			res[i+j].Add(res[i+j], term)
		}
	}
	return res
}
