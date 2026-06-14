package sequences

import (
	"fmt"
	"math/big"
)

// Fubini numbers (OEIS A000670): number of preferential arrangements of n labeled elements.
// URL: https://oeis.org/A000670
// a(n) = sum_{k=1}^n binomial(n, k) * a(n-k) for n >= 1, a(0) = 1.

var fubiniCache = make(map[int64]*big.Int)
var binomialCache = make(map[string]*big.Int)

// GetFubiniSequence returns the Fubini numbers sequence (OEIS A000670).
func GetFubiniSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetFubiniAtPosition(maxNumber)
	}
	return GenerateFubiniSequence(maxNumber)
}

// GenerateFubiniSequence generates the A000670 sequence up to maxNumber (a(0), ..., a(maxNumber)).
func GenerateFubiniSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := maxNumber.Int64()
	sequence := make([]*big.Int, limit+1)

	for i := int64(0); i <= limit; i++ {
		sequence[i] = CalculateFubini(i)
	}

	return &NumericSequence{
		Name:     "Fubini numbers (A000670)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[len(sequence)-1],
	}, nil
}

// GetFubiniAtPosition returns the n-th term of A000670 (n >= 0).
func GetFubiniAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	result := CalculateFubini(n.Int64())

	return &NumericSequence{
		Name:     "Fubini numbers (A000670)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// CalculateFubini calculates a(n) for OEIS A000670.
func CalculateFubini(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	if val, ok := fubiniCache[n]; ok {
		return val
	}

	res := big.NewInt(0)
	for k := int64(1); k <= n; k++ {
		term := new(big.Int).Mul(binomial(n, k), CalculateFubini(n-k))
		res.Add(res, term)
	}

	fubiniCache[n] = new(big.Int).Set(res)
	return fubiniCache[n]
}

// binomial calculates the binomial coefficient (n choose k).
func binomial(n, k int64) *big.Int {
	if k < 0 || k > n {
		return big.NewInt(0)
	}
	if k == 0 || k == n {
		return big.NewInt(1)
	}
	if k > n/2 {
		k = n - k
	}

	key := fmt.Sprintf("%d,%d", n, k)
	if val, ok := binomialCache[key]; ok {
		return val
	}

	res := big.NewInt(1)
	for i := int64(1); i <= k; i++ {
		res.Mul(res, big.NewInt(n-i+1))
		res.Div(res, big.NewInt(i))
	}

	binomialCache[key] = new(big.Int).Set(res)
	return binomialCache[key]
}
