package sequences

import (
	"fmt"
	"math/big"
)

// Number of Abelian groups of order n (OEIS A000688)
// URL: https://oeis.org/A000688
// a(n) = product P(e_i) where n = product p_i^e_i and P is the partition function (A000041).

// GetAbelianGroupsOrderNSequence returns the Number of Abelian groups of order n sequence (OEIS A000688).
func GetAbelianGroupsOrderNSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetAbelianGroupsOrderNAtPosition(maxNumber)
	}
	return GenerateAbelianGroupsOrderNSequence(maxNumber)
}

// GenerateAbelianGroupsOrderNSequence generates the A000688 sequence up to maxNumber (a(1), ..., a(maxNumber)).
func GenerateAbelianGroupsOrderNSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("max number must be at least 1")
	}

	limit := int(maxNumber.Int64())
	sequence := make([]*big.Int, limit)

	for i := 1; i <= limit; i++ {
		sequence[i-1] = CalculateAbelianGroupsOrderN(int64(i))
	}

	return &NumericSequence{
		Name:     "Number of Abelian groups (A000688)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[len(sequence)-1],
	}, nil
}

// GetAbelianGroupsOrderNAtPosition returns the n-th term of A000688 (n >= 1).
func GetAbelianGroupsOrderNAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("position must be at least 1")
	}

	result := CalculateAbelianGroupsOrderN(n.Int64())

	return &NumericSequence{
		Name:     "Number of Abelian groups (A000688)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

// CalculateAbelianGroupsOrderN calculates a(n) for OEIS A000688.
func CalculateAbelianGroupsOrderN(n int64) *big.Int {
	if n == 1 {
		return big.NewInt(1)
	}

	factors := factorize(n)
	res := big.NewInt(1)
	for _, exponent := range factors {
		res.Mul(res, partition(exponent))
	}
	return res
}

// factorize returns the exponents of the prime factorization of n.
func factorize(n int64) []int {
	exponents := []int{}
	d := int64(2)
	temp := n
	for d*d <= temp {
		if temp%d == 0 {
			count := 0
			for temp%d == 0 {
				count++
				temp /= d
			}
			exponents = append(exponents, count)
		}
		d++
	}
	if temp > 1 {
		exponents = append(exponents, 1)
	}
	return exponents
}

var partitionCache = make(map[int]*big.Int)

// partition calculates the number of partitions of n (A000041).
func partition(n int) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	if val, ok := partitionCache[n]; ok {
		return val
	}

	res := big.NewInt(0)
	for k := 1; ; k++ {
		pent1 := k * (3*k - 1) / 2
		pent2 := -k * (-3*k - 1) / 2

		stop := true
		if pent1 <= n {
			term := partition(n - pent1)
			if k%2 == 1 {
				res.Add(res, term)
			} else {
				res.Sub(res, term)
			}
			stop = false
		}
		if pent2 <= n {
			term := partition(n - pent2)
			if k%2 == 1 {
				res.Add(res, term)
			} else {
				res.Sub(res, term)
			}
			stop = false
		}
		if stop {
			break
		}
	}

	partitionCache[n] = new(big.Int).Set(res)
	return partitionCache[n]
}
