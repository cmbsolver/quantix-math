package sequences

import (
	"fmt"
	"math/big"
)

// Kolakoski Sequence (OEIS A000002)
// URL: https://oeis.org/A000002
// a(n) is the length of n-th run; a(1) = 1; sequence consists just of 1's and 2's.

// GetKolakoskiSequence returns the Kolakoski sequence (OEIS A000002).
func GetKolakoskiSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetKolakoskiSequenceAtPosition(maxNumber)
	}
	return GenerateKolakoskiSequence(maxNumber)
}

// GenerateKolakoskiSequence generates the A000002 sequence up to maxNumber (a(1), a(2), ..., a(maxNumber)).
func GenerateKolakoskiSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("max number must be at least 1")
	}

	n := int(maxNumber.Int64())
	res := generateKolakoski(n)

	sequence := make([]*big.Int, len(res))
	for i, v := range res {
		sequence[i] = big.NewInt(int64(v))
	}

	return &NumericSequence{
		Name:     "Kolakoski Sequence (A000002)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[len(sequence)-1],
	}, nil
}

// GetKolakoskiSequenceAtPosition returns the n-th term of A000002 (n >= 1).
func GetKolakoskiSequenceAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("position must be at least 1")
	}

	pos := int(n.Int64())
	res := generateKolakoski(pos)
	result := big.NewInt(int64(res[pos-1]))

	return &NumericSequence{
		Name:     "Kolakoski Sequence (A000002)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// generateKolakoski generates the first n terms of the Kolakoski sequence.
func generateKolakoski(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{1}
	}
	if n == 2 {
		return []int{1, 2}
	}

	a := make([]int, n)
	a[0] = 1
	a[1] = 2
	a[2] = 2

	// m is the index of the run length we are currently processing.
	// it starts at 2 (the 3rd term, which is 2)
	m := 2
	// lastIdx is the index of the last element we added to the sequence.
	lastIdx := 2

	for lastIdx < n-1 {
		valToAppend := 1
		if m%2 == 0 {
			valToAppend = 1
		} else {
			valToAppend = 2
		}

		runLen := a[m]
		for i := 0; i < runLen && lastIdx < n-1; i++ {
			lastIdx++
			a[lastIdx] = valToAppend
		}
		m++
	}

	return a
}
