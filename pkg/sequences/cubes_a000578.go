package sequences

import (
	"fmt"
	"math/big"
)

// Cubes: a(n) = n^3 (A000578)
// URL: https://oeis.org/A000578
// Description: a(n) = n^3.

// GetCubesA000578Sequence returns the cubes sequence (OEIS A000578).
func GetCubesA000578Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetCubeAtPosition(maxNumber)
	}
	return GenerateCubesSequence(maxNumber)
}

// GenerateCubesSequence generates the A000578 sequence up to maxNumber.
func GenerateCubesSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return &NumericSequence{
			Name:     "Cubes (A000578)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}

	var sequence []*big.Int
	n := big.NewInt(0)
	for {
		cube := new(big.Int).Exp(n, big.NewInt(3), nil)
		if cube.Cmp(maxNumber) > 0 {
			break
		}
		sequence = append(sequence, cube)
		n = new(big.Int).Add(n, big.NewInt(1))
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Cubes (A000578)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetCubeAtPosition returns the n-th term of A000578.
func GetCubeAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	val := new(big.Int).Exp(n, big.NewInt(3), nil)

	return &NumericSequence{
		Name:     "Cubes (A000578)",
		Number:   n,
		Sequence: []*big.Int{val},
		Result:   val,
	}, nil
}
