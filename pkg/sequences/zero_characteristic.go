package sequences

import (
	"fmt"
	"math/big"
)

// Zero Characteristic Sequence (OEIS A000007)
// URL: https://oeis.org/A000007
// The characteristic function of {0}: a(n) = 1 if n = 0, else 0.

// GetZeroCharacteristicSequence returns the characteristic function of {0} (OEIS A000007).
func GetZeroCharacteristicSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetZeroCharacteristicAtPosition(maxNumber)
	}
	return GenerateZeroCharacteristicSequence(maxNumber)
}

// GenerateZeroCharacteristicSequence generates the A000007 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateZeroCharacteristicSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	// a(0) = 1
	sequence[0] = big.NewInt(1)

	// a(n) = 0 for n > 0
	for i := 1; i <= n; i++ {
		sequence[i] = big.NewInt(0)
	}

	result := big.NewInt(0)
	if n < len(sequence) {
		result = sequence[n]
	}

	return &NumericSequence{
		Name:     "Zero Characteristic (A000007)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetZeroCharacteristicAtPosition returns the n-th term of A000007.
func GetZeroCharacteristicAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	var result *big.Int
	if n.Cmp(big.NewInt(0)) == 0 {
		result = big.NewInt(1)
	} else {
		result = big.NewInt(0)
	}

	return &NumericSequence{
		Name:     "Zero Characteristic (A000007)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
