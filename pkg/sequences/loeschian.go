package sequences

import (
	"fmt"
	"math/big"
)

// Loeschian numbers (OEIS A003136)
// URL: https://oeis.org/A003136
// Numbers of the form x^2 + xy + y^2; norms of vectors in A2 lattice.

// GetLoeschianSequence returns the Loeschian numbers sequence (OEIS A003136).
func GetLoeschianSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetLoeschianAtPosition(maxNumber)
	}
	return GenerateLoeschianSequence(maxNumber)
}

// GenerateLoeschianSequence generates the A003136 sequence up to maxNumber.
// It returns all Loeschian numbers <= maxNumber.
func GenerateLoeschianSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := maxNumber.Int64()
	isLoeschian := make([]bool, limit+1)

	// x^2 + xy + y^2 <= limit
	// Since x^2 <= limit, x <= sqrt(limit)
	// Similarly for y.
	for x := int64(0); x*x <= limit; x++ {
		for y := int64(0); y <= x; y++ {
			val := x*x + x*y + y*y
			if val <= limit {
				isLoeschian[val] = true
			}
		}
	}

	var sequence []*big.Int
	for i, loeschian := range isLoeschian {
		if loeschian {
			sequence = append(sequence, big.NewInt(int64(i)))
		}
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	} else {
		result = big.NewInt(0)
	}

	return &NumericSequence{
		Name:     "Loeschian numbers (A003136)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetLoeschianAtPosition returns the n-th term of A003136 (n >= 0).
func GetLoeschianAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	pos := n.Int64()
	// We don't have a direct formula for the n-th term, so we estimate and generate.
	// For small n, we can just generate enough.
	// The density of Loeschian numbers is known, but generating is safer for small n.

	limit := int64(100)
	if pos > 20 {
		limit = pos * pos // Overestimate
	}

	for {
		seq, _ := GenerateLoeschianSequence(big.NewInt(limit))
		if int64(len(seq.Sequence)) > pos {
			result := seq.Sequence[pos]
			return &NumericSequence{
				Name:     "Loeschian numbers (A003136)",
				Number:   n,
				Sequence: []*big.Int{result},
				Result:   result,
			}, nil
		}
		limit *= 2
	}
}

// IsLoeschian checks if a number is a Loeschian number.
// A number is Loeschian if and only if in its prime factorization,
// every prime p ≡ 2 (mod 3) appears with an even exponent.
func IsLoeschian(n *big.Int) bool {
	if n.Sign() < 0 {
		return false
	}
	if n.Sign() == 0 {
		return true
	}

	d := new(big.Int).Set(n)
	two := big.NewInt(2)
	three := big.NewInt(3)

	// Handle 2 separately (2 ≡ 2 mod 3)
	if new(big.Int).Mod(d, two).Sign() == 0 {
		exponent := 0
		for new(big.Int).Mod(d, two).Sign() == 0 {
			exponent++
			d.Div(d, two)
		}
		if exponent%2 != 0 {
			return false
		}
	}

	// Check odd primes
	i := big.NewInt(3)
	for new(big.Int).Mul(i, i).Cmp(d) <= 0 {
		if new(big.Int).Mod(d, i).Sign() == 0 {
			exponent := 0
			for new(big.Int).Mod(d, i).Sign() == 0 {
				exponent++
				d.Div(d, i)
			}
			// If p ≡ 2 (mod 3), exponent must be even
			if new(big.Int).Mod(i, three).Cmp(two) == 0 {
				if exponent%2 != 0 {
					return false
				}
			}
		}
		i.Add(i, two)
	}

	if d.Cmp(big.NewInt(1)) > 0 {
		// Remaining d is prime
		if new(big.Int).Mod(d, three).Cmp(two) == 0 {
			return false
		}
	}

	return true
}
