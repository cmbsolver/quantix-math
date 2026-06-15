package sequences

import (
	"fmt"
	"math/big"
	"math/bits"
)

// Odious numbers: numbers with an odd number of 1's in their binary expansion (OEIS A000069).
// URL: https://oeis.org/A000069
// a(n) is the n-th number with an odd number of 1s in its binary representation.

// GetOdiousNumbersSequence returns the odious numbers up to maxNumber or at a position (OEIS A000069).
func GetOdiousNumbersSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetOdiousNumberAtPosition(maxNumber)
	}
	return GenerateOdiousNumbersSequence(maxNumber)
}

// GenerateOdiousNumbersSequence generates the A000069 sequence up to maxNumber.
// It returns all odious numbers <= maxNumber.
func GenerateOdiousNumbersSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number must be at least 0")
	}

	limit := maxNumber.Uint64()
	var sequence []*big.Int

	for i := uint64(1); i <= limit; i++ {
		if bits.OnesCount64(i)%2 != 0 {
			sequence = append(sequence, new(big.Int).SetUint64(i))
		}
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	} else {
		result = big.NewInt(0)
	}

	return &NumericSequence{
		Name:     "Odious numbers (A000069)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetOdiousNumberAtPosition returns the n-th term of A000069 (n >= 1).
func GetOdiousNumberAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("position must be at least 1")
	}

	pos := n.Uint64()
	var count uint64 = 0
	var i uint64 = 0
	var result uint64

	for count < pos {
		i++
		if bits.OnesCount64(i)%2 != 0 {
			count++
			result = i
		}
		// Safety break for very large n that might take too long
		if i > 10000000 && count < pos {
			return nil, fmt.Errorf("position too large for simple calculation")
		}
	}

	resBig := new(big.Int).SetUint64(result)
	return &NumericSequence{
		Name:     "Odious numbers (A000069)",
		Number:   n,
		Sequence: []*big.Int{resBig},
		Result:   resBig,
	}, nil
}

// IsOdious checks if a number is an odious number.
func IsOdious(n *big.Int) bool {
	if n.Sign() <= 0 {
		return false
	}
	// For big.Int, we can use BitLen and check each bit, or use a better way if available.
	// However, bits.OnesCount64 is fine for numbers within uint64.
	if n.IsUint64() {
		return bits.OnesCount64(n.Uint64())%2 != 0
	}

	// For larger numbers:
	count := 0
	for i := 0; i < n.BitLen(); i++ {
		if n.Bit(i) == 1 {
			count++
		}
	}
	return count%2 != 0
}
