package sequences

import (
	"fmt"
	"math/big"
)

// Number of classes of primitive positive definite binary quadratic forms of discriminant D = -4n.
// URL: https://oeis.org/A000003
// Description: a(n) is the number of reduced primitive positive definite binary quadratic forms ax^2 + bxy + cy^2 of discriminant b^2 - 4ac = -4n.
// For D = -4n, we have b^2 - 4ac = -4n, which means b must be even, say b = 2B.
// Then 4B^2 - 4ac = -4n  => B^2 - ac = -n => ac - B^2 = n.
// The form is ax^2 + 2Bxy + cy^2.
// Reduction conditions for positive definite forms:
// 1. |2B| <= a <= c
// 2. If |2B| = a or a = c, then 2B >= 0.
// Primitive means gcd(a, 2B, c) = 1.

// GetBinaryQuadraticFormsA000003Sequence returns the A000003 sequence.
func GetBinaryQuadraticFormsA000003Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetBinaryQuadraticFormsA000003AtPosition(maxNumber)
	}
	return GenerateBinaryQuadraticFormsA000003Sequence(maxNumber)
}

// GenerateBinaryQuadraticFormsA000003Sequence generates the A000003 sequence up to maxNumber.
func GenerateBinaryQuadraticFormsA000003Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() <= 0 {
		return &NumericSequence{
			Name:     "Binary quadratic forms (A000003)",
			Number:   maxNumber,
			Sequence: []*big.Int{},
		}, nil
	}

	limit := maxNumber.Int64()
	if maxNumber.IsInt64() && limit > 0 {
		var sequence []*big.Int
		for n := int64(1); n <= limit; n++ {
			sequence = append(sequence, big.NewInt(CountPrimitiveForms(n)))
		}
		var result *big.Int
		if len(sequence) > 0 {
			result = sequence[len(sequence)-1]
		}
		return &NumericSequence{
			Name:     "Binary quadratic forms (A000003)",
			Number:   maxNumber,
			Sequence: sequence,
			Result:   result,
		}, nil
	}

	return nil, fmt.Errorf("maxNumber too large for current implementation")
}

// GetBinaryQuadraticFormsA000003AtPosition returns the n-th term of A000003.
func GetBinaryQuadraticFormsA000003AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() <= 0 {
		return nil, fmt.Errorf("position must be positive")
	}

	if n.IsInt64() {
		val := big.NewInt(CountPrimitiveForms(n.Int64()))
		return &NumericSequence{
			Name:     "Binary quadratic forms (A000003)",
			Number:   n,
			Sequence: []*big.Int{val},
			Result:   val,
		}, nil
	}

	return nil, fmt.Errorf("position too large for current implementation")
}

// CountPrimitiveForms counts reduced primitive positive definite binary quadratic forms of determinant n.
func CountPrimitiveForms(n int64) int64 {
	count := int64(0)
	// Determinant n = ac - B^2, discriminant D = -4n.
	// b = 2B, so ac - (b/2)^2 = n.
	// Reduction conditions: |b| <= a <= c. If |b| = a or a = c, then b >= 0.
	// Since b = 2B, this means |2B| <= a <= c. If |2B| = a or a = c, then 2B >= 0.
	// ac = n + B^2.
	// Since a^2 <= ac = n + B^2 and |2B| <= a, we have B^2 <= a^2/4 <= (n+B^2)/4.
	// 4B^2 <= n + B^2 => 3B^2 <= n => B^2 <= n/3.

	for B := int64(0); B*B <= n/3; B++ {
		target := n + B*B
		// Find all factors a of target such that a >= 2*|B| and a*a <= target.
		// B can be positive, zero, or negative.
		// For B > 0, we check both B and -B.

		bs := []int64{B}
		if B > 0 {
			bs = append(bs, -B)
		}

		for _, b := range bs {
			// b = 2B or -2B is handled by the loop over bs.
			// Actually the loop above uses B as the half of b.
			// The conditions are:
			// 1. |b| <= a <= c
			// 2. If |b| = a or a = c, then b >= 0.
			// b here is what I called 2B in comments.

			curr_b := 2 * b
			abs_b := curr_b
			if abs_b < 0 {
				abs_b = -abs_b
			}

			for a := int64(1); a*a <= target; a++ {
				if target%a == 0 {
					c := target / a

					// Reduction conditions:
					// a <= c is guaranteed by a*a <= target
					if abs_b <= a {
						// Primitive check: gcd(a, b, c) = 1
						if gcd3(a, curr_b, c) == 1 {
							// If |b| = a or a = c, then b >= 0.
							if abs_b == a || a == c {
								if curr_b >= 0 {
									count++
								}
							} else {
								count++
							}
						}
					}
				}
			}
		}
	}
	return count
}

func gcd3(a, b, c int64) int64 {
	return gcdA000003(a, gcdA000003(b, c))
}

func gcdA000003(a, b int64) int64 {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a %= b
		a, b = b, a
	}
	return a
}
