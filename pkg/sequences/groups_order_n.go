package sequences

import (
	"fmt"
	"math/big"
)

// Number of groups of order n (OEIS A000001)
// URL: https://oeis.org/A000001
// a(n) is the number of nonisomorphic groups of order n.

// a000001Data contains the values of A000001 for n = 0 to 100.
// Values taken from OEIS A000001.
var a000001Data = []int64{
	0, 1, 1, 1, 2, 1, 2, 1, 5, 2, 2, 1, 5, 1, 2, 1, 14, 1, 5, 1, 5, 2, 2, 1, 15, 2, 2, 5, 4, 1, 4, 1, 51, 1, 2, 1, 14, 1, 2, 2, 14, 1, 6, 1, 4, 2, 2, 1, 52, 2, 5, 1, 5, 1, 15, 2, 13, 2, 2, 1, 13, 1, 2, 4, 267, 1, 4, 1, 5, 1, 4, 1, 50, 1, 2, 3, 4, 1, 6, 1, 52, 15, 2, 1, 15, 1, 2, 1, 12, 1, 10, 1, 4, 2, 2, 1, 231, 1, 5, 2, 16,
}

// GetGroupsOrderNSequence returns the Number of groups of order n sequence (OEIS A000001).
// Since calculating a(n) for large n is a hard problem, this implementation uses a lookup table
// for known values up to 100 and returns an error for larger values.
func GetGroupsOrderNSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetGroupsOrderNAtPosition(maxNumber)
	}
	return GenerateGroupsOrderNSequence(maxNumber)
}

// GenerateGroupsOrderNSequence generates the A000001 sequence up to maxNumber (a(0), a(1), ..., a(maxNumber)).
func GenerateGroupsOrderNSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := int(maxNumber.Int64())
	if limit >= len(a000001Data) {
		return nil, fmt.Errorf("sequence A000001 is only implemented up to n=%d", len(a000001Data)-1)
	}

	sequence := make([]*big.Int, limit+1)
	for i := 0; i <= limit; i++ {
		sequence[i] = big.NewInt(a000001Data[i])
	}

	return &NumericSequence{
		Name:     "Number of Groups (A000001)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[len(sequence)-1],
	}, nil
}

// GetGroupsOrderNAtPosition returns the n-th term of A000001 (n >= 0).
func GetGroupsOrderNAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	pos := int(n.Int64())
	if pos >= len(a000001Data) {
		return nil, fmt.Errorf("sequence A000001 is only implemented up to n=%d", len(a000001Data)-1)
	}

	result := big.NewInt(a000001Data[pos])

	return &NumericSequence{
		Name:     "Number of Groups (A000001)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
