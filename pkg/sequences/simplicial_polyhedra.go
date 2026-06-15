package sequences

import (
	"fmt"
	"math/big"
)

// Simplicial Polyhedra (OEIS A000109)
// URL: https://oeis.org/A000109
// Number of simplicial polyhedra with n vertices; simple planar graphs with n vertices and 3n-6 edges;
// maximal simple planar graphs with n vertices; planar triangulations with n vertices;
// triangulations of the sphere with n vertices; 3-connected cubic planar graphs on 2n-4 vertices.
//
// These values are known for n = 3 to 23.
var a000109Values = []string{
	"1", "1", "1", "2", "5", "14", "50", "233", "1249", "7595", "49566", "339722", "2406841", "17490241", "129664753", "977526957", "7475907149", "57896349553", "453382272049", "3585853662949", "28615703421545",
}

// GetSimplicialPolyhedraSequence returns the number of simplicial polyhedra with n vertices (OEIS A000109).
func GetSimplicialPolyhedraSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetSimplicialPolyhedraAtPosition(maxNumber)
	}
	return GenerateSimplicialPolyhedraSequence(maxNumber)
}

// GenerateSimplicialPolyhedraSequence generates the A000109 sequence up to maxNumber.
// It returns a(3), a(4), ..., a(maxNumber).
func GenerateSimplicialPolyhedraSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(3)) < 0 {
		return nil, fmt.Errorf("max number must be at least 3")
	}

	n := int(maxNumber.Int64())
	if n > 23 {
		return nil, fmt.Errorf("max number for A000109 is currently limited to 23")
	}

	size := n - 3 + 1
	sequence := make([]*big.Int, size)

	for i := 0; i < size; i++ {
		val := new(big.Int)
		val.SetString(a000109Values[i], 10)
		sequence[i] = val
	}

	return &NumericSequence{
		Name:     "Simplicial polyhedra (A000109)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[size-1],
	}, nil
}

// GetSimplicialPolyhedraAtPosition returns the n-th term of A000109 (n >= 3).
func GetSimplicialPolyhedraAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(3)) < 0 {
		return nil, fmt.Errorf("position must be at least 3")
	}

	pos := int(n.Int64())
	if pos > 23 {
		return nil, fmt.Errorf("position for A000109 is currently limited to 23")
	}

	result := new(big.Int)
	result.SetString(a000109Values[pos-3], 10)

	return &NumericSequence{
		Name:     "Simplicial polyhedra (A000109)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
