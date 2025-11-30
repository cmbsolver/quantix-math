package sequences

import (
	"math/big"
)

// GCD calculates the greatest common divisor of a and b.
func GCD(a, b *big.Int) *big.Int {
	zero := big.NewInt(0)
	for b.Cmp(zero) != 0 {
		t := new(big.Int).Set(b)
		b = b.Mod(a, b)
		a.Set(t)
	}
	return new(big.Int).Set(a)
}

// GetTotientSequence generates the Totient sequence.
func GetTotientSequence(maxNumber *big.Int) (*NumericSequence, error) {
	retval := &NumericSequence{Name: "Totient", Number: new(big.Int).Set(maxNumber)}
	one := big.NewInt(1)
	i := big.NewInt(1)
	cmp := i.Cmp(maxNumber)

	for cmp <= 0 {
		o := new(big.Int).Set(i)
		n := new(big.Int).Set(maxNumber)
		if GCD(o, n).Cmp(one) == 0 {
			retval.Sequence = append(retval.Sequence, new(big.Int).Set(i))
		}
		i = i.Add(i, one)
		cmp = i.Cmp(maxNumber)
	}

	retval.Result = big.NewInt(int64(len(retval.Sequence)))
	return retval, nil
}

// GetTotientPrimeSequence generates the Totient Prime sequence.
func GetTotientPrimeSequence(maxNumber *big.Int) (*NumericSequence, error) {
	retval := &NumericSequence{Name: "Totient", Number: new(big.Int).Set(maxNumber)}
	one := big.NewInt(1)
	i := big.NewInt(1)
	cmp := i.Cmp(maxNumber)

	for cmp <= 0 {
		o := new(big.Int).Set(i)
		n := new(big.Int).Set(maxNumber)
		if GCD(o, n).Cmp(one) == 0 {
			isPrime := IsPrime(i)
			if isPrime {
				retval.Sequence = append(retval.Sequence, new(big.Int).Set(i))
			}
		}
		i = i.Add(i, one)
		cmp = i.Cmp(maxNumber)
	}

	retval.Result = big.NewInt(int64(len(retval.Sequence)))
	return retval, nil
}
