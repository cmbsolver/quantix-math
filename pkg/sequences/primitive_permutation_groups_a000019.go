package sequences

import (
	"fmt"
	"math/big"
)

// PrimitivePermutationGroupsSequence represents the number of primitive permutation groups of degree n (A000019).
// URL: https://oeis.org/A000019
// a(n) is the number of primitive permutation groups of degree n.
type PrimitivePermutationGroupsSequence struct{}

// a000019Data contains the values of A000019 for n = 1 to 92.
// Values taken from OEIS A000019.
var a000019Data = []int64{
	1, 1, 2, 2, 5, 4, 7, 7, 11, 9, 8, 6, 9, 4, 6, 22, 10, 4, 8, 4, 9, 4, 7, 5, 28, 7, 15, 14, 8, 4, 12, 7, 4, 2, 6, 22, 11, 4, 2, 8, 10, 4, 10, 4, 9, 2, 6, 4, 40, 9, 2, 3, 8, 4, 8, 9, 5, 2, 6, 9, 14, 4, 8, 74, 13, 7, 10, 7, 2, 2, 10, 4, 16, 4, 2, 2, 4, 6, 10, 4, 155, 10, 6, 6, 6, 2, 2, 2, 10, 4, 10, 2,
}

// CalculatePrimitivePermutationGroups calculates the number of primitive permutation groups of degree n.
// This implementation uses a lookup table for known values up to 92.
func CalculatePrimitivePermutationGroups(n int64) int64 {
	if n <= 0 {
		return 0
	}

	index := int(n - 1)
	if index < len(a000019Data) {
		return a000019Data[index]
	}

	return -1 // Indicates not implemented for this n
}

// GetPrimitivePermutationGroupsA000019Sequence returns the number of primitive permutation groups of degree n sequence (OEIS A000019).
func GetPrimitivePermutationGroupsA000019Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	if isPositional {
		return GetPrimitivePermutationGroupsA000019AtPosition(maxNumber)
	}
	return GeneratePrimitivePermutationGroupsA000019Sequence(maxNumber)
}

// GeneratePrimitivePermutationGroupsA000019Sequence generates the A000019 sequence up to maxNumber.
func GeneratePrimitivePermutationGroupsA000019Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	limit := int(maxNumber.Int64())
	sequence := make([]*big.Int, 0, limit)

	for i := 1; i <= limit; i++ {
		val := CalculatePrimitivePermutationGroups(int64(i))
		if val == -1 {
			return nil, fmt.Errorf("sequence A000019 is not implemented for n=%d", i)
		}
		sequence = append(sequence, big.NewInt(val))
	}

	result := big.NewInt(0)
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Number of primitive permutation groups of degree n (A000019)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetPrimitivePermutationGroupsA000019AtPosition returns the n-th term of A000019.
func GetPrimitivePermutationGroupsA000019AtPosition(n *big.Int) (*NumericSequence, error) {
	pos := n.Int64()
	val := CalculatePrimitivePermutationGroups(pos)
	if val == -1 {
		return nil, fmt.Errorf("sequence A000019 is not implemented for n=%d", pos)
	}

	result := big.NewInt(val)

	return &NumericSequence{
		Name:     "Number of primitive permutation groups of degree n (A000019)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// IsPrimitivePermutationGroupA000019 checks if a number exists in the A000019 sequence.
func IsPrimitivePermutationGroupA000019(n *big.Int) (bool, string, error) {
	if n.Sign() < 0 {
		return false, "", nil
	}

	for i := 1; i <= len(a000019Data); i++ {
		val := CalculatePrimitivePermutationGroups(int64(i))
		if big.NewInt(val).Cmp(n) == 0 {
			return true, fmt.Sprintf("%d", i), nil
		}
	}

	return false, "", nil
}
