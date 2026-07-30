package sequences

import (
	"fmt"
	"math/big"
)

// Ménage hit polynomials diagonal term: a(n) = A058087(n, n-2) (OEIS A000033).
// URL: https://oeis.org/A000033

// GetA000033Sequence returns the A000033 sequence.
func GetA000033Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000033AtPosition(maxNumber)
	}

	return GenerateA000033Sequence(maxNumber)
}

// GenerateA000033Sequence generates A000033 from a(0) through a(maxNumber).
func GenerateA000033Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := maxNumber.Int64()
	sequence := make([]*big.Int, limit+1)

	for n := int64(0); n <= limit; n++ {
		term, err := calculateA000033(n, sequence)
		if err != nil {
			return nil, err
		}
		sequence[n] = term
	}

	return &NumericSequence{
		Name:     "Ménage hit polynomials (A000033)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[limit],
	}, nil
}

// GetA000033AtPosition returns the n-th term of A000033.
func GetA000033AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	sequence := make([]*big.Int, n.Int64()+1)
	term, err := calculateA000033(n.Int64(), sequence)
	if err != nil {
		return nil, err
	}

	return &NumericSequence{
		Name:     "Ménage hit polynomials (A000033)",
		Number:   n,
		Sequence: []*big.Int{term},
		Result:   term,
	}, nil
}

// calculateA000033 computes a(n) for OEIS A000033 using its recurrence.
func calculateA000033(index int64, sequence []*big.Int) (*big.Int, error) {
	switch index {
	case 0:
		return big.NewInt(0), nil
	case 1:
		return big.NewInt(2), nil
	case 2:
		return big.NewInt(3), nil
	case 3:
		return big.NewInt(4), nil
	}

	n := index + 1

	getTerm := func(index int64) (*big.Int, error) {
		if sequence[index] != nil {
			return sequence[index], nil
		}
		term, err := calculateA000033(index, sequence)
		if err != nil {
			return nil, err
		}
		sequence[index] = term
		return term, nil
	}

	aN1, err := getTerm(index - 1)
	if err != nil {
		return nil, err
	}
	aN2, err := getTerm(index - 2)
	if err != nil {
		return nil, err
	}
	aN3, err := getTerm(index - 3)
	if err != nil {
		return nil, err
	}

	tN := big.NewInt(n)
	nMinus2 := big.NewInt(n - 2)
	nMinus3 := big.NewInt(n - 3)
	nMinus4 := big.NewInt(n - 4)
	twoNMinus3 := big.NewInt(2*n - 3)
	twoNMinus5 := big.NewInt(2*n - 5)
	twoNMinus7 := big.NewInt(2*n - 7)

	leftCoeff := new(big.Int).Mul(nMinus3, nMinus2)
	leftCoeff.Mul(leftCoeff, twoNMinus5)
	leftCoeff.Mul(leftCoeff, twoNMinus7)

	term1Coeff := new(big.Int).Mul(nMinus3, nMinus2)
	term1Coeff.Mul(term1Coeff, tN)
	term1Coeff.Mul(term1Coeff, new(big.Int).Mul(twoNMinus7, twoNMinus7))
	term1 := new(big.Int).Mul(term1Coeff, aN1)

	term2Coeff := new(big.Int).Mul(nMinus4, nMinus3)
	term2Coeff.Mul(term2Coeff, tN)
	term2Coeff.Mul(term2Coeff, new(big.Int).Mul(twoNMinus3, twoNMinus3))
	term2 := new(big.Int).Mul(term2Coeff, aN2)

	term3Coeff := new(big.Int).Mul(nMinus2, tN)
	term3Coeff.Mul(term3Coeff, twoNMinus5)
	term3Coeff.Mul(term3Coeff, twoNMinus3)
	term3 := new(big.Int).Mul(term3Coeff, aN3)

	numerator := new(big.Int).Add(term1, term2)
	numerator.Add(numerator, term3)

	quotient := new(big.Int)
	remainder := new(big.Int)
	quotient.QuoRem(numerator, leftCoeff, remainder)
	if remainder.Sign() != 0 {
		return nil, fmt.Errorf("non-integer division while computing A000033 at n=%d", index)
	}

	return quotient, nil
}
