package sequences

import (
	"fmt"
	"math/big"
)

// A100640: Triangle read by rows: numerators of Cotesian numbers C(n,k) (0 <= k <= n).
// URL: https://oeis.org/A100640
// Description: Flattened triangle of numerators of Cotesian numbers C(n,k), read by rows.
//
// Formula (from OEIS Maple entry):
// C(n,0) = C(n,n) = (1/n!) * Sum_{a=1..n+1} n^a*s(n,a)/(a+1), for n > 0,
// C(0,0) = 0,
// and for 0 < k < n:
// C(n,k) = (1/n!)*binomial(n,k) *
// Sum_{a=1..k+1} Sum_{b=1..n-k+1}
// n^(a+b)*s(k,a)*s(n-k,b)/((b+1)*binomial(a+b+1,b+1)),
// where s(n,k) are signed Stirling numbers of the first kind.

// GetCotesianNumeratorsA100640Sequence returns the A100640 sequence.
func GetCotesianNumeratorsA100640Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetCotesianNumeratorsA100640AtPosition(maxNumber)
	}

	return GenerateCotesianNumeratorsA100640Sequence(maxNumber)
}

// GenerateCotesianNumeratorsA100640Sequence generates the first maxNumber flattened terms of A100640.
func GenerateCotesianNumeratorsA100640Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	if !maxNumber.IsInt64() {
		return nil, fmt.Errorf("max number too large for current implementation")
	}

	termCount := int(maxNumber.Int64())
	sequence := make([]*big.Int, termCount)

	for i := 0; i < termCount; i++ {
		n, k := a100640IndexToRowCol(i)
		sequence[i] = CalculateCotesianNumeratorA100640(n, k)
	}

	result := big.NewInt(0)
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Numerators of Cotesian numbers C(n,k) (A100640)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetCotesianNumeratorsA100640AtPosition returns the flattened n-th term of A100640.
func GetCotesianNumeratorsA100640AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	if !n.IsInt64() {
		return nil, fmt.Errorf("position too large for current implementation")
	}

	row, col := a100640IndexToRowCol(int(n.Int64()))
	value := CalculateCotesianNumeratorA100640(row, col)

	return &NumericSequence{
		Name:     "Numerators of Cotesian numbers C(n,k) (A100640)",
		Number:   n,
		Sequence: []*big.Int{value},
		Result:   value,
	}, nil
}

// CalculateCotesianNumeratorA100640 computes the numerator of Cotesian number C(n,k) for 0 <= k <= n.
func CalculateCotesianNumeratorA100640(n, k int) *big.Int {
	if n < 0 || k < 0 || k > n {
		return big.NewInt(0)
	}
	if n == 0 {
		return big.NewInt(0)
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
		return new(big.Int).Set(result.Num())
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

	return new(big.Int).Set(result.Num())
}

// a100640IndexToRowCol maps flattened index i to triangle coordinates (n, k).
func a100640IndexToRowCol(i int) (int, int) {
	row := 0
	for i > row {
		i -= (row + 1)
		row++
	}
	return row, i
}

// signedStirlingFirstKindTable builds signed Stirling numbers of the first kind up to n.
func signedStirlingFirstKindTable(n int) [][]*big.Int {
	table := make([][]*big.Int, n+1)
	for i := 0; i <= n; i++ {
		table[i] = make([]*big.Int, n+1)
		for j := 0; j <= n; j++ {
			table[i][j] = big.NewInt(0)
		}
	}
	table[0][0] = big.NewInt(1)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			left := new(big.Int).Set(table[i-1][j-1])
			right := new(big.Int).Mul(big.NewInt(int64(i-1)), table[i-1][j])
			table[i][j] = new(big.Int).Sub(left, right)
		}
	}

	return table
}
