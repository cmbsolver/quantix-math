package sequences

import (
	"fmt"
	"math/big"
	"math/bits"
)

// Hamming Weight Sequence (OEIS A000120)
// URL: https://oeis.org/A000120
// Ones-counting sequence: number of 1's in binary expansion of n (or the binary weight of n).

// GetHammingWeightSequence returns the number of 1's in binary expansion of n (OEIS A000120).
func GetHammingWeightSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetHammingWeightAtPosition(maxNumber)
	}
	return GenerateHammingWeightSequence(maxNumber)
}

// GenerateHammingWeightSequence generates the A000120 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateHammingWeightSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number must be at least 0 for this sequence")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		sequence[i] = big.NewInt(int64(bits.OnesCount64(uint64(i))))
	}

	return &NumericSequence{
		Name:     "Hamming weight (A000120)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetHammingWeightAtPosition returns the n-th term of A000120 (n >= 0).
func GetHammingWeightAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	val := n.Uint64()
	result := big.NewInt(int64(bits.OnesCount64(val)))

	return &NumericSequence{
		Name:     "Hamming weight (A000120)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
