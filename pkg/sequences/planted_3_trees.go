package sequences

import (
	"fmt"
	"math/big"
)

// Planted 3-trees of height n (OEIS A002658)
// URL: https://oeis.org/A002658
// a(0) = a(1) = 1; for n > 0, a(n+1) = a(n)*(a(0) + ... + a(n-1)) + a(n)*(a(n) + 1)/2.
// Also called planted 3-trees or planted unary-binary trees.

// GetPlanted3TreesSequence returns the number of planted 3-trees of height n (OEIS A002658).
func GetPlanted3TreesSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetPlanted3TreesAtPosition(maxNumber)
	}
	return GeneratePlanted3TreesSequence(maxNumber)
}

// GeneratePlanted3TreesSequence generates the A002658 sequence up to maxNumber (a(0), a(1), ..., a(maxNumber)).
func GeneratePlanted3TreesSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	// Pre-calculate sums to keep it O(n)
	for i := 0; i <= n; i++ {
		var val *big.Int
		if i == 0 || i == 1 {
			val = big.NewInt(1)
		} else {
			// a(i) = a(i-1)*(a(0) + ... + a(i-2)) + a(i-1)*(a(i-1) + 1)/2
			// This matches a(n+1) = a(n)*(a(0) + ... + a(n-1)) + a(n)*(a(n)+1)/2
			// where n = i-1.
			prev := sequence[i-1]

			currentSum := big.NewInt(0)
			for j := 0; j <= i-2; j++ {
				currentSum.Add(currentSum, sequence[j])
			}

			// term1 = prev * currentSum
			term1 := new(big.Int).Mul(prev, currentSum)

			// term2 = prev * (prev + 1) / 2
			term2 := new(big.Int).Add(prev, big.NewInt(1))
			term2.Mul(term2, prev)
			term2.Div(term2, big.NewInt(2))

			val = new(big.Int).Add(term1, term2)
		}
		sequence[i] = val
	}

	return &NumericSequence{
		Name:     "Planted 3-trees of height n (A002658)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetPlanted3TreesAtPosition returns the n-th term of A002658 (n >= 0).
func GetPlanted3TreesAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	seq, err := GeneratePlanted3TreesSequence(n)
	if err != nil {
		return nil, err
	}

	return &NumericSequence{
		Name:     "Planted 3-trees of height n (A002658)",
		Number:   n,
		Sequence: []*big.Int{seq.Result},
		Result:   seq.Result,
	}, nil
}

// IsPlanted3Tree checks if a number is a member of the A002658 sequence.
func IsPlanted3Tree(n *big.Int) (bool, string) {
	if n.Sign() < 0 {
		return false, ""
	}

	// Since the sequence grows very rapidly, we can just generate terms until we exceed n.
	// 1, 1, 2, 7, 56, 2212, 2595782, ...

	if n.Cmp(big.NewInt(1)) == 0 {
		return true, "0, 1"
	}

	// Re-implement IsPlanted3Tree more cleanly
	fullSeq, _ := GeneratePlanted3TreesSequence(big.NewInt(20)) // 20 terms is plenty for big.Int
	for i, val := range fullSeq.Sequence {
		if val.Cmp(n) == 0 {
			if i == 0 || i == 1 {
				return true, "0, 1"
			}
			return true, fmt.Sprintf("%d", i)
		}
		if val.Cmp(n) > 0 {
			break
		}
	}
	return false, ""
}
