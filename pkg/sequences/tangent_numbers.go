package sequences

import (
	"fmt"
	"math/big"
)

// Tangent numbers (A000182)
// URL: https://oeis.org/A000182
// Description: Tangent numbers: e.g.f. tan(x) = sum(n >= 1) T_n * x^(2*n-1) / (2*n-1)!
// The sequence starts: 1, 2, 16, 272, 7936, 353792, 22368256, 1903757312, ...

// GetTangentNumbersSequence returns the sequence of Tangent numbers.
func GetTangentNumbersSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetTangentNumberAtPosition(maxNumber)
	}
	return GenerateTangentNumbersSequence(maxNumber)
}

// GenerateTangentNumbersSequence generates Tangent numbers up to maxNumber.
func GenerateTangentNumbersSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() <= 0 {
		return &NumericSequence{
			Name:     "Tangent numbers (A000182)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
			Result:   nil,
		}, nil
	}

	var sequence []*big.Int
	// T(n, k) = (k-1)*T(n, k-1) + (k+1)*T(n-1, k+1) is not the standard way,
	// let's use the Seidel triangle or the iterative method.
	// We'll use the algorithm described in OEIS for Tangent numbers.

	for n := 1; ; n++ {
		tn := calculateTangentNumber(n)
		if tn.Cmp(maxNumber) > 0 {
			break
		}
		sequence = append(sequence, tn)
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	}

	return &NumericSequence{
		Name:     "Tangent numbers (A000182)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetTangentNumberAtPosition returns the n-th Tangent number (1-indexed).
func GetTangentNumberAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() <= 0 {
		return nil, fmt.Errorf("position must be a positive integer")
	}

	if !n.IsInt64() {
		return nil, fmt.Errorf("position %s is too large", n.String())
	}

	pos := int(n.Int64())
	result := calculateTangentNumber(pos)

	return &NumericSequence{
		Name:     "Tangent numbers (A000182)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateTangentNumber computes the n-th Tangent number using an iterative algorithm.
// We use the property that T_n are related to Zig-Zag numbers or using a triangle.
func calculateTangentNumber(n int) *big.Int {
	if n <= 0 {
		return big.NewInt(0)
	}

	// We can use the Seidel-like triangle for tangent numbers (also known as the tangent triangle)
	// For n=1: 1
	// For n=2: 2
	// For n=3: 16
	// The algorithm from Richard Brent's paper (provided in search results):
	// T_1 = 1
	// For k=2..n:
	//   T_k = (k-1)! (initialization)
	//   ...
	// Actually, a simpler way is to use the zig-zag triangle (Andre permutations)
	// and pick the values at odd positions.
	// Euler numbers E_n (A000111) are the total number of zig-zag permutations.
	// Tangent numbers T_n are the number of zig-zag permutations of 2n-1 elements.

	size := 2*n - 1
	if size <= 0 {
		return big.NewInt(0)
	}

	// Andre triangle (also known as Euler triangle or Zigzag triangle)
	// Row 0: 1
	// Row 1: 0 1
	// Row 2: 1 1 0
	// Row 3: 0 1 2 2
	// Row 4: 5 5 4 2 0

	row := make([]*big.Int, size+1)
	row[0] = big.NewInt(1)

	for i := 1; i <= size; i++ {
		newRow := make([]*big.Int, i+1)
		if i%2 == 1 {
			// Odd row: fill from right to left
			newRow[i] = big.NewInt(0)
			for j := i - 1; j >= 0; j-- {
				newRow[j] = new(big.Int).Add(newRow[j+1], row[j])
			}
		} else {
			// Even row: fill from left to right
			newRow[0] = big.NewInt(0)
			for j := 1; j <= i; j++ {
				newRow[j] = new(big.Int).Add(newRow[j-1], row[j-1])
			}
		}
		row = newRow
	}

	// The Andre numbers are at the ends of the rows.
	// For n=1, size=1, row is [1, 0] or [0, 1]. A(1) = 1.
	// For n=2, size=3, row is [0, 1, 2, 2]. A(3) = 2.
	// For n=3, size=5, row is [0, 5, 10, 14, 16, 16]. A(5) = 16.

	if size%2 == 1 {
		return row[0]
	}
	return row[size]
}
