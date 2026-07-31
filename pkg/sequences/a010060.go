package sequences

import (
	"fmt"
	"math/big"
)

// A010060: Thue-Morse sequence.
// URL: https://oeis.org/A010060
// Description: a(n) is the parity of the number of 1's in the binary representation of n.

// GetA010060Sequence returns the A010060 sequence up to maxNumber terms or the n-th term.
func GetA010060Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA010060AtPosition(maxNumber)
	}

	return GenerateA010060Sequence(maxNumber)
}

// GenerateA010060Sequence generates the A010060 sequence up to maxNumber terms.
func GenerateA010060Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	if !maxNumber.IsInt64() {
		return nil, fmt.Errorf("max number too large for current implementation")
	}

	n := maxNumber.Int64()
	sequence := make([]*big.Int, n)

	for i := int64(0); i < n; i++ {
		sequence[i] = big.NewInt(CalculateA010060(i))
	}

	result := big.NewInt(0)
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Thue-Morse sequence (A010060)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetA010060AtPosition returns the n-th term of A010060.
func GetA010060AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	value := big.NewInt(0)
	for i := 0; i < n.BitLen(); i++ {
		value.Xor(value, big.NewInt(int64(n.Bit(i))))
	}

	return &NumericSequence{
		Name:     "Thue-Morse sequence (A010060)",
		Number:   n,
		Sequence: []*big.Int{value},
		Result:   value,
	}, nil
}

// CalculateA010060 calculates the n-th term of A010060 for n >= 0.
func CalculateA010060(n int64) int64 {
	if n < 0 {
		return 0
	}

	parity := int64(0)
	for n > 0 {
		parity ^= (n & 1)
		n >>= 1
	}

	return parity
}
