package sequences

import (
	"fmt"
	"math/big"
)

// Unlabeled trees with n nodes (OEIS A000055)
// URL: https://oeis.org/A000055
// a(n) is the number of trees with n unlabeled nodes.
// For n > 0, a(n) = a81(n) - sum_{i=1}^{floor(n/2)} a81(i)*a81(n-i) + (if n is even: a81(n/2)*(a81(n/2)+1)/2 else 0)
// where a81(n) is the number of rooted unlabeled trees with n nodes (A000081).

// GetUnlabeledTreesSequence returns the number of unlabeled trees with n nodes (OEIS A000055).
func GetUnlabeledTreesSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetUnlabeledTreesAtPosition(maxNumber)
	}
	return GenerateUnlabeledTreesSequence(maxNumber)
}

// GenerateUnlabeledTreesSequence generates the A000055 sequence up to maxNumber.
func GenerateUnlabeledTreesSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number must be at least 0")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	if n >= 0 {
		sequence[0] = big.NewInt(1)
	}
	if n == 0 {
		return &NumericSequence{
			Name:     "Unlabeled trees (A000055)",
			Number:   maxNumber,
			Sequence: sequence,
			Result:   sequence[0],
		}, nil
	}

	// We need A000081 up to n
	a81Seq, err := GenerateRootedUnlabeledTreesSequence(big.NewInt(int64(n)))
	if err != nil {
		return nil, err
	}
	a81 := a81Seq.Sequence

	for i := 1; i <= n; i++ {
		// t := a81(i)
		t := new(big.Int).Set(a81[i])

		// sum_{j=1}^{floor(i/2)} a81(j)*a81(i-j)
		limit := i / 2
		sum := new(big.Int)
		for j := 1; j <= limit; j++ {
			term := new(big.Int).Mul(a81[j], a81[i-j])
			sum.Add(sum, term)
		}
		t.Sub(t, sum)

		// if i is even: + a81(i/2)*(a81(i/2)+1)/2
		if i%2 == 0 {
			val := a81[i/2]
			term := new(big.Int).Add(val, big.NewInt(1))
			term.Mul(term, val)
			term.Div(term, big.NewInt(2))
			t.Add(t, term)
		}

		sequence[i] = t
	}

	return &NumericSequence{
		Name:     "Unlabeled trees (A000055)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetUnlabeledTreesAtPosition returns the n-th term of A000055.
func GetUnlabeledTreesAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	seq, err := GenerateUnlabeledTreesSequence(n)
	if err != nil {
		return nil, err
	}

	return &NumericSequence{
		Name:     "Unlabeled trees (A000055)",
		Number:   n,
		Sequence: []*big.Int{seq.Result},
		Result:   seq.Result,
	}, nil
}
