package sequences

import (
	"fmt"
	"math/big"
)

// Bicolorable Primitive Necklaces (OEIS A000048)
// URL: https://oeis.org/A000048
// a(n) is the number of n-bead necklaces with beads of 2 colors and primitive period n,
// when turning over is not allowed but the two colors can be interchanged.
//
// Formula: a(n) = (1/(2*n)) * Sum_{odd d|n} mu(d) * 2^(n/d)

// GetBicolorableNecklacesSequence returns the A000048 sequence.
func GetBicolorableNecklacesSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetBicolorableNecklacesAtPosition(maxNumber)
	}
	return GenerateBicolorableNecklacesSequence(maxNumber)
}

// GenerateBicolorableNecklacesSequence generates the A000048 sequence up to maxNumber.
func GenerateBicolorableNecklacesSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number must be at least 0")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	for i := 0; i <= n; i++ {
		val := calculateA000048(int64(i))
		sequence[i] = val
	}

	return &NumericSequence{
		Name:     "Bicolorable Primitive Necklaces (A000048)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetBicolorableNecklacesAtPosition returns the n-th term of A000048.
func GetBicolorableNecklacesAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	result := calculateA000048(n.Int64())

	return &NumericSequence{
		Name:     "Bicolorable Primitive Necklaces (A000048)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// calculateA000048 implements the formula: a(n) = (1/(2*n)) * Sum_{odd d|n} mu(d) * 2^(n/d)
func calculateA000048(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	if n == 1 {
		return big.NewInt(1)
	}

	total := new(big.Int)
	divs := getOddDivisors(n)

	for _, d := range divs {
		m := getMobius(d)
		if m == 0 {
			continue
		}

		// 2^(n/d)
		term := new(big.Int).Lsh(big.NewInt(1), uint(n/d))

		if m == -1 {
			total.Sub(total, term)
		} else {
			total.Add(total, term)
		}
	}

	// Result = total / (2 * n)
	divisor := big.NewInt(2 * n)
	result := new(big.Int).Div(total, divisor)

	return result
}

func getOddDivisors(n int64) []int64 {
	var divs []int64
	for i := int64(1); i*i <= n; i++ {
		if n%i == 0 {
			if i%2 != 0 {
				divs = append(divs, i)
			}
			j := n / i
			if j != i && j%2 != 0 {
				divs = append(divs, j)
			}
		}
	}
	return divs
}

func getMobius(n int64) int {
	if n == 1 {
		return 1
	}
	pCount := 0
	temp := n
	for i := int64(2); i*i <= temp; i++ {
		if temp%i == 0 {
			pCount++
			temp /= i
			if temp%i == 0 {
				return 0
			}
		}
	}
	if temp > 1 {
		pCount++
	}
	if pCount%2 == 0 {
		return 1
	}
	return -1
}
