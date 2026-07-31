package sequences

import (
	"fmt"
	"math/big"
)

// A100641: Triangle read by rows: denominators of Cotesian numbers C(n,k) (0 <= k <= n).
// URL: https://oeis.org/A100641
// Description: Flattened triangle of denominators of Cotesian numbers C(n,k), read by rows.
//
// Formula (from OEIS Maple entry):
// C(n,0) = C(n,n) = (1/n!) * Sum_{a=1..n+1} n^a*s(n,a)/(a+1), for n > 0,
// C(0,0) = 1,
// and for 0 < k < n:
// C(n,k) = (1/n!)*binomial(n,k) *
// Sum_{a=1..k+1} Sum_{b=1..n-k+1}
// n^(a+b)*s(k,a)*s(n-k,b)/((b+1)*binomial(a+b+1,b+1)),
// where s(n,k) are signed Stirling numbers of the first kind.

// GetCotesianDenominatorsA100641Sequence returns the A100641 sequence.
func GetCotesianDenominatorsA100641Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetCotesianDenominatorsA100641AtPosition(maxNumber)
	}

	return GenerateCotesianDenominatorsA100641Sequence(maxNumber)
}

// GenerateCotesianDenominatorsA100641Sequence generates the first maxNumber flattened terms of A100641.
func GenerateCotesianDenominatorsA100641Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	if !maxNumber.IsInt64() {
		return nil, fmt.Errorf("max number too large for current implementation")
	}

	termCount := int(maxNumber.Int64())
	sequence := make([]*big.Int, termCount)

	for i := 0; i < termCount; i++ {
		n, k := a100641IndexToRowCol(i)
		sequence[i] = CalculateCotesianDenominatorA100641(n, k)
	}

	result := big.NewInt(0)
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Denominators of Cotesian numbers C(n,k) (A100641)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetCotesianDenominatorsA100641AtPosition returns the flattened n-th term of A100641.
func GetCotesianDenominatorsA100641AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	if !n.IsInt64() {
		return nil, fmt.Errorf("position too large for current implementation")
	}

	row, col := a100641IndexToRowCol(int(n.Int64()))
	value := CalculateCotesianDenominatorA100641(row, col)

	return &NumericSequence{
		Name:     "Denominators of Cotesian numbers C(n,k) (A100641)",
		Number:   n,
		Sequence: []*big.Int{value},
		Result:   value,
	}, nil
}

// CalculateCotesianDenominatorA100641 computes the denominator of Cotesian number C(n,k) for 0 <= k <= n.
func CalculateCotesianDenominatorA100641(n, k int) *big.Int {
	if n < 0 || k < 0 || k > n {
		return big.NewInt(0)
	}
	if n == 0 {
		return big.NewInt(1)
	}

	stirlingN := signedStirlingFirstKindTable(n)
	nFactorial := CalculateFactorial(int64(n))

	if k == 0 || k == n {
		sum := new(big.Rat)
		for a := 1; a <= n+1; a++ {
			if a > n {
				continue
			}
			s := stirlingN[n][a]
			if s.Sign() == 0 {
				continue
			}

			nPow := new(big.Int).Exp(big.NewInt(int64(n)), big.NewInt(int64(a)), nil)
			numerator := new(big.Int).Mul(nPow, s)
			term := new(big.Rat).SetFrac(numerator, big.NewInt(int64(a+1)))
			sum.Add(sum, term)
		}

		result := new(big.Rat).SetFrac(sum.Num(), new(big.Int).Mul(sum.Denom(), nFactorial))
		return new(big.Int).Set(result.Denom())
	}

	stirlingK := signedStirlingFirstKindTable(k)
	stirlingNK := signedStirlingFirstKindTable(n - k)

	sum := new(big.Rat)
	for a := 1; a <= k+1; a++ {
		if a > k {
			continue
		}
		sa := stirlingK[k][a]
		if sa.Sign() == 0 {
			continue
		}

		for b := 1; b <= n-k+1; b++ {
			if b > n-k {
				continue
			}
			sb := stirlingNK[n-k][b]
			if sb.Sign() == 0 {
				continue
			}

			nPow := new(big.Int).Exp(big.NewInt(int64(n)), big.NewInt(int64(a+b)), nil)
			numerator := new(big.Int).Mul(nPow, sa)
			numerator.Mul(numerator, sb)

			denominator := new(big.Int).Binomial(int64(a+b+1), int64(b+1))
			denominator.Mul(denominator, big.NewInt(int64(b+1)))

			term := new(big.Rat).SetFrac(numerator, denominator)
			sum.Add(sum, term)
		}
	}

	prefactorNumerator := new(big.Int).Binomial(int64(n), int64(k))
	result := new(big.Rat).Mul(new(big.Rat).SetFrac(prefactorNumerator, nFactorial), sum)

	return new(big.Int).Set(result.Denom())
}

// a100641IndexToRowCol maps flattened index i to triangle coordinates (n, k).
func a100641IndexToRowCol(i int) (int, int) {
	row := 0
	for i > row {
		i -= (row + 1)
		row++
	}
	return row, i
}
