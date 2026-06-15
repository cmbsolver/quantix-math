package sequences

import (
	"fmt"
	"math/big"
)

// Connected Planar Graphs (OEIS A003094)
// URL: https://oeis.org/A003094
// a(n) is the number of unlabeled connected planar simple graphs with n nodes.

// A003094_values contains precomputed values for OEIS A003094.
// n=0 to 13.
var A003094_values = []string{
	"1",          // a(0)
	"1",          // a(1)
	"1",          // a(2)
	"2",          // a(3)
	"6",          // a(4)
	"20",         // a(5)
	"99",         // a(6)
	"646",        // a(7)
	"5974",       // a(8)
	"71885",      // a(9)
	"1052805",    // a(10)
	"17449299",   // a(11)
	"313372298",  // a(12)
	"5942258308", // a(13)
}

// GetConnectedPlanarGraphsSequence returns the number of unlabeled connected planar simple graphs on n nodes (OEIS A003094).
func GetConnectedPlanarGraphsSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetConnectedPlanarGraphsAtPosition(maxNumber)
	}
	return GenerateConnectedPlanarGraphsSequence(maxNumber)
}

// GenerateConnectedPlanarGraphsSequence generates the A003094 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateConnectedPlanarGraphsSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number must be at least 0")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		val, err := calculateA003094(int64(i))
		if err != nil {
			return nil, err
		}
		sequence[i] = val
	}

	return &NumericSequence{
		Name:     "Connected Planar Graphs (A003094)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetConnectedPlanarGraphsAtPosition returns the n-th term of A003094 (n >= 0).
func GetConnectedPlanarGraphsAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	result, err := calculateA003094(n.Int64())
	if err != nil {
		return nil, err
	}

	return &NumericSequence{
		Name:     "Connected Planar Graphs (A003094)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateA003094 returns the n-th term of OEIS A003094.
// For small n, it uses precomputed values.
func calculateA003094(n int64) (*big.Int, error) {
	if n < int64(len(A003094_values)) {
		res := new(big.Int)
		res.SetString(A003094_values[n], 10)
		return res, nil
	}

	// For larger n, the calculation is extremely complex (Inverse Euler Transform of A005470).
	// Since A005470 is also hard to calculate, we limit to precomputed values for now.
	return nil, fmt.Errorf("calculation for n=%d is not implemented beyond precomputed table", n)
}
