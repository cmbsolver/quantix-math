package sequences

import (
	"fmt"
	"math/big"
)

// Number of necklaces with n beads of 2 colors, allowing turning over (these are also called bracelets).
// URL: https://oeis.org/A000029
// Description: a(n) is the number of necklaces with n beads of 2 colors, allowing turning over.
// Offset: 0, 1
// a(0) = 1; for n > 0: a(n) = 1/(2n) * sum_{d|n} phi(d) * 2^(n/d) + correction(n)
// where correction(n) = 2^((n-1)/2) if n is odd, and 2^(n/2-1) + 2^(n/2-2) if n is even.

// GetA000029Sequence returns the number of necklaces with n beads of 2 colors (OEIS A000029).
func GetA000029Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000029AtPosition(maxNumber)
	}
	return GenerateA000029Sequence(maxNumber)
}

// GenerateA000029Sequence generates the A000029 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateA000029Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	nLimit := int(maxNumber.Int64())
	sequence := make([]*big.Int, nLimit+1)

	for i := 0; i <= nLimit; i++ {
		sequence[i] = calculateA000029(int64(i))
	}

	return &NumericSequence{
		Name:     "Bracelets with n beads of 2 colors (A000029)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[nLimit],
	}, nil
}

// GetA000029AtPosition returns the n-th term of A000029.
func GetA000029AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	val := n.Int64()
	result := calculateA000029(val)

	return &NumericSequence{
		Name:     "Bracelets with n beads of 2 colors (A000029)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateA000029 calculates the n-th term of A000029.
func calculateA000029(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	// a(n) = (A000031(n) + A164090(n)) / 2
	// A000031(n) = 1/n * sum_{d|n} phi(d) * 2^(n/d)
	// A164090(n):
	// if n odd: 2^((n+1)/2)
	// if n even: 2^(n/2) + 2^(n/2-1) = 3 * 2^(n/2-1)

	sumA000031 := big.NewInt(0)
	for d := int64(1); d <= n; d++ {
		if n%d == 0 {
			phiD := big.NewInt(EulerTotient(d))
			twoPow := new(big.Int).Exp(big.NewInt(2), big.NewInt(n/d), nil)
			term := new(big.Int).Mul(phiD, twoPow)
			sumA000031.Add(sumA000031, term)
		}
	}
	// Note: sumA000031 is n * A000031(n)

	var a164090 *big.Int
	if n%2 != 0 {
		a164090 = new(big.Int).Exp(big.NewInt(2), big.NewInt((n+1)/2), nil)
	} else {
		term1 := new(big.Int).Exp(big.NewInt(2), big.NewInt(n/2), nil)
		term2 := new(big.Int).Exp(big.NewInt(2), big.NewInt(n/2-1), nil)
		a164090 = new(big.Int).Add(term1, term2)
	}

	// a(n) = (sumA000031/n + a164090) / 2
	// To avoid float, a(n) = (sumA000031 + n*a164090) / (2n)

	numerator := new(big.Int).Add(sumA000031, new(big.Int).Mul(big.NewInt(n), a164090))
	denominator := new(big.Int).Mul(big.NewInt(2), big.NewInt(n))

	return new(big.Int).Div(numerator, denominator)
}
