package sequences

import (
	"fmt"
	"math/big"
)

// Stern's diatomic series (OEIS A002487)
// URL: https://oeis.org/A002487
// a(0) = 0, a(1) = 1; for n > 0: a(2n) = a(n), a(2n+1) = a(n) + a(n+1).

// GetSternSequence returns Stern's diatomic series up to maxNumber or at a specific position.
func GetSternSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetSternAtPosition(maxNumber)
	}
	return GenerateSternSequence(maxNumber)
}

// GenerateSternSequence generates the A002487 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateSternSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number must be at least 0 for this sequence")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		sequence[i] = sternFusc(int64(i))
	}

	return &NumericSequence{
		Name:     "Stern's diatomic series (A002487)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetSternAtPosition returns the n-th term of A002487 (n >= 0).
func GetSternAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	val := n.Int64()
	result := sternFusc(val)

	return &NumericSequence{
		Name:     "Stern's diatomic series (A002487)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// sternFusc calculates the n-th term of Stern's diatomic series (fusc function).
// Efficient iterative version:
// a = 1, b = 0
// while n > 0:
//
//	if n is odd: b = a + b
//	else: a = a + b
//	n = n / 2
//
// return b
func sternFusc(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	var a, b int64 = 1, 0
	for n > 0 {
		if n%2 == 1 {
			b += a
		} else {
			a += b
		}
		n /= 2
	}
	return big.NewInt(b)
}

// IsSternNumber checks if a number exists in Stern's diatomic series.
// Since a(n)/a(n+1) covers all positive rationals, every positive integer n exists in the sequence.
// Specifically, a(2^(n-1)) = 1, and the maximum value in the diatomic array grows.
func IsSternNumber(n *big.Int) (bool, string) {
	if n.Sign() < 0 {
		return false, ""
	}
	if n.Sign() == 0 {
		return true, "a(0) = 0"
	}
	// Every positive integer appears in the sequence.
	// For example, a(2^n-1) = n. (Wait, let's verify this)
	// a(1)=1, a(3)=a(1)+a(2)=1+1=2, a(7)=a(3)+a(4)=2+1=3, a(15)=a(7)+a(8)=3+1=4.
	// So a(2^n-1) = n.
	pos := new(big.Int).Exp(big.NewInt(2), n, nil)
	pos.Sub(pos, big.NewInt(1))
	return true, fmt.Sprintf("a(%s) = %s", pos.String(), n.String())
}
