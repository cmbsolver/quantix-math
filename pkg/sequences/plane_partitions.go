package sequences

import (
	"fmt"
	"math/big"
)

// Plane partitions (OEIS A000219)
// URL: https://oeis.org/A000219
// Number of plane partitions (or planar partitions) of n.
// Two-dimensional partitions of n in which no row or column is longer than the one before it.

// GetPlanePartitionsSequence returns the A000219 sequence.
func GetPlanePartitionsSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetPlanePartitionsAtPosition(maxNumber)
	}
	return GeneratePlanePartitionsSequence(maxNumber)
}

// GeneratePlanePartitionsSequence generates the A000219 sequence up to maxNumber.
func GeneratePlanePartitionsSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	// Using the recurrence: a(n) = (1/n) * sum_{k=1}^n a(n-k) * sigma_2(k)
	// where sigma_2(k) is the sum of the squares of the divisors of k.

	dp := make([]*big.Int, n+1)
	dp[0] = big.NewInt(1)

	sigma2 := make([]*big.Int, n+1)
	for i := 1; i <= n; i++ {
		sigma2[i] = big.NewInt(0)
		for d := 1; d <= i; d++ {
			if i%d == 0 {
				d2 := big.NewInt(int64(d))
				d2.Mul(d2, d2)
				sigma2[i].Add(sigma2[i], d2)
			}
		}
	}

	for i := 1; i <= n; i++ {
		sum := big.NewInt(0)
		for k := 1; k <= i; k++ {
			term := new(big.Int).Mul(dp[i-k], sigma2[k])
			sum.Add(sum, term)
		}
		dp[i] = new(big.Int).Div(sum, big.NewInt(int64(i)))
	}

	for i := 0; i <= n; i++ {
		sequence[i] = dp[i]
	}

	return &NumericSequence{
		Name:     "Plane partitions (A000219)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetPlanePartitionsAtPosition returns the n-th term of A000219.
func GetPlanePartitionsAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	target := int(n.Int64())

	// Reusing the generation logic as it's efficient enough for typical n
	seq, err := GeneratePlanePartitionsSequence(n)
	if err != nil {
		return nil, err
	}

	result := seq.Sequence[target]

	return &NumericSequence{
		Name:     "Plane partitions (A000219)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
