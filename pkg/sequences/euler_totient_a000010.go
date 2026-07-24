package sequences

import (
	"fmt"
	"math/big"
)

// Euler totient function phi(n): count numbers <= n and prime to n.
// URL: https://oeis.org/A000010
// Description: a(n) is the number of positive integers k <= n such that gcd(n, k) = 1.

// GetEulerTotientA000010Sequence returns the Euler totient sequence (OEIS A000010).
func GetEulerTotientA000010Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetEulerTotientA000010AtPosition(maxNumber)
	}
	return GenerateEulerTotientA000010Sequence(maxNumber)
}

// GenerateEulerTotientA000010Sequence generates the A000010 sequence up to maxNumber.
func GenerateEulerTotientA000010Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() <= 0 {
		return &NumericSequence{
			Name:     "Euler totient function phi(n) (A000010)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}

	limit := maxNumber.Int64()
	if maxNumber.IsInt64() && limit > 0 {
		var sequence []*big.Int
		for n := int64(1); n <= limit; n++ {
			sequence = append(sequence, big.NewInt(EulerTotient(n)))
		}
		var result *big.Int
		if len(sequence) > 0 {
			result = sequence[len(sequence)-1]
		}
		return &NumericSequence{
			Name:     "Euler totient function phi(n) (A000010)",
			Number:   maxNumber,
			Sequence: sequence,
			Result:   result,
		}, nil
	}

	return nil, fmt.Errorf("maxNumber too large for current implementation")
}

// GetEulerTotientA000010AtPosition returns the n-th term of A000010.
func GetEulerTotientA000010AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() <= 0 {
		return nil, fmt.Errorf("position must be positive")
	}

	if n.IsInt64() {
		val := big.NewInt(EulerTotient(n.Int64()))
		return &NumericSequence{
			Name:     "Euler totient function phi(n) (A000010)",
			Number:   n,
			Sequence: []*big.Int{val},
			Result:   val,
		}, nil
	}

	return nil, fmt.Errorf("position too large for current implementation")
}

// EulerTotient calculates phi(n) for int64.
func EulerTotient(n int64) int64 {
	if n <= 0 {
		return 0
	}
	result := n
	temp := n
	for i := int64(2); i*i <= temp; i++ {
		if temp%i == 0 {
			for temp%i == 0 {
				temp /= i
			}
			result -= result / i
		}
	}
	if temp > 1 {
		result -= result / temp
	}
	return result
}
