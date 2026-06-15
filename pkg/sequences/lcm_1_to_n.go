package sequences

import (
	"fmt"
	"math/big"
)

// Least common multiple (or LCM) of {1, 2, ..., n} (OEIS A003418)
// URL: https://oeis.org/A003418
// a(n) = lcm(1, 2, ..., n) for n >= 1, a(0) = 1.

// GetLCM1ToNSequence returns the least common multiple of {1, 2, ..., n} (OEIS A003418).
func GetLCM1ToNSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetLCM1ToNAtPosition(maxNumber)
	}
	return GenerateLCM1ToNSequence(maxNumber)
}

// GenerateLCM1ToNSequence generates the A003418 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateLCM1ToNSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number must be non-negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	currentLCM := big.NewInt(1)
	sequence[0] = new(big.Int).Set(currentLCM)

	for i := 1; i <= n; i++ {
		currentLCM = lcm(currentLCM, big.NewInt(int64(i)))
		sequence[i] = new(big.Int).Set(currentLCM)
	}

	return &NumericSequence{
		Name:     "Least common multiple of {1, 2, ..., n} (A003418)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetLCM1ToNAtPosition returns the n-th term of A003418 (n >= 0).
func GetLCM1ToNAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be non-negative")
	}

	val := n.Int64()
	result := big.NewInt(1)
	for i := int64(1); i <= val; i++ {
		result = lcm(result, big.NewInt(i))
	}

	return &NumericSequence{
		Name:     "Least common multiple of {1, 2, ..., n} (A003418)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// lcm calculates the least common multiple of two big.Ints.
func lcm(a, b *big.Int) *big.Int {
	if a.Sign() == 0 || b.Sign() == 0 {
		return big.NewInt(0)
	}
	// lcm(a, b) = |a*b| / gcd(a, b)
	temp := new(big.Int).Mul(a, b)
	temp.Abs(temp)
	g := new(big.Int).GCD(nil, nil, a, b)
	return temp.Div(temp, g)
}

// IsLCM1ToN checks if the given number is in the sequence A003418.
func IsLCM1ToN(n *big.Int) (bool, string) {
	if n.Sign() < 0 {
		return false, ""
	}
	if n.Cmp(big.NewInt(1)) == 0 {
		return true, "0 or 1"
	}

	currentLCM := big.NewInt(1)
	for i := int64(2); ; i++ {
		currentLCM = lcm(currentLCM, big.NewInt(i))
		cmp := currentLCM.Cmp(n)
		if cmp == 0 {
			// Find if there are more indices with same LCM
			start := i
			for {
				nextLCM := lcm(currentLCM, big.NewInt(i+1))
				if nextLCM.Cmp(currentLCM) == 0 {
					i++
				} else {
					break
				}
			}
			if i == start {
				return true, fmt.Sprintf("%d", i)
			}
			return true, fmt.Sprintf("%d-%d", start, i)
		}
		if cmp > 0 {
			return false, ""
		}
		// Safety break for very large numbers if needed, but big.Int handles it.
		// However, a(n) grows fast (~ e^n), so we won't loop too many times.
	}
}
