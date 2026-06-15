package sequences

import (
	"math/big"
)

// PrimesA000040 represents the prime numbers sequence (A000040).
// A number p is prime if (and only if) it is greater than 1 and has no positive divisors except 1 and p.
// URL: https://oeis.org/A000040
type PrimesA000040 struct{}

// GetPrimesA000040Sequence generates the prime numbers sequence up to maxNumber or at a specific position.
// If isPositional is true, it returns the n-th prime (where n is maxNumber).
// If isPositional is false, it returns all primes up to maxNumber.
func GetPrimesA000040Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	numericSequence := &NumericSequence{
		Name:   "The prime numbers (A000040)",
		Number: new(big.Int).Set(maxNumber),
	}

	if isPositional {
		// Calculate the n-th prime
		n := maxNumber.Uint64()
		if n == 0 {
			return numericSequence, nil
		}

		p := big.NewInt(2)
		var count uint64 = 1
		for count < n {
			p.Add(p, big.NewInt(1))
			if IsPrime(p) {
				count++
			}
		}
		numericSequence.Sequence = append(numericSequence.Sequence, new(big.Int).Set(p))
	} else {
		// Calculate all primes up to maxNumber
		for i := big.NewInt(2); i.Cmp(maxNumber) <= 0; i.Add(i, big.NewInt(1)) {
			if IsPrime(i) {
				numericSequence.Sequence = append(numericSequence.Sequence, new(big.Int).Set(i))
			}
		}
	}

	return numericSequence, nil
}
