package sequences

import (
	"fmt"
	"math/big"
)

// Mersenne prime exponents: primes p such that 2^p - 1 is prime.
// URL: https://oeis.org/A000043

// GetMersennePrimeExponentsSequence returns the Mersenne prime exponents sequence (OEIS A000043).
// If isPositional is true, it returns the n-th term.
// Otherwise, it returns all terms up to maxNumber.
func GetMersennePrimeExponentsSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetMersennePrimeExponentsAtPosition(maxNumber)
	}
	return GenerateMersennePrimeExponentsSequence(maxNumber)
}

// GenerateMersennePrimeExponentsSequence generates the A000043 sequence up to maxNumber.
// It returns a(1), a(2), ..., a(n) such that a(n) <= maxNumber.
func GenerateMersennePrimeExponentsSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	var sequence []*big.Int
	// We iterate through primes p and check if 2^p - 1 is prime using Lucas-Lehmer test.
	// Since we need to calculate it, we'll implement a basic prime sieve/check for p,
	// and then the Lucas-Lehmer test for 2^p - 1.

	p := big.NewInt(2)
	for p.Cmp(maxNumber) <= 0 {
		if IsPrime(p) {
			if IsMersennePrime(p) {
				sequence = append(sequence, new(big.Int).Set(p))
			}
		}
		p = new(big.Int).Add(p, big.NewInt(1))
	}

	var result *big.Int
	if len(sequence) > 0 {
		result = sequence[len(sequence)-1]
	} else {
		result = big.NewInt(0)
	}

	return &NumericSequence{
		Name:     "Mersenne prime exponents (A000043)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetMersennePrimeExponentsAtPosition returns the n-th term of A000043 (n >= 1).
func GetMersennePrimeExponentsAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() <= 0 {
		return nil, fmt.Errorf("position must be positive")
	}

	target := int(n.Int64())
	var sequence []*big.Int
	count := 0
	p := big.NewInt(2)

	for count < target {
		if IsPrime(p) {
			if IsMersennePrime(p) {
				count++
				if count == target {
					sequence = append(sequence, new(big.Int).Set(p))
					break
				}
			}
		}
		p = new(big.Int).Add(p, big.NewInt(1))
		// Safety break for very large N in this environment
		if p.Cmp(big.NewInt(10000)) > 0 {
			break
		}
	}

	if len(sequence) == 0 {
		return nil, fmt.Errorf("could not find %d-th Mersenne prime exponent in reasonable time", target)
	}

	return &NumericSequence{
		Name:     "Mersenne prime exponents (A000043)",
		Number:   n,
		Sequence: sequence,
		Result:   sequence[0],
	}, nil
}

// IsMersennePrime checks if 2^p - 1 is prime using the Lucas-Lehmer test.
// p must be a prime.
func IsMersennePrime(p *big.Int) bool {
	pInt := p.Int64()
	if pInt == 2 {
		return true // 2^2 - 1 = 3 is prime
	}

	// Lucas-Lehmer test:
	// s_0 = 4
	// s_i = (s_{i-1}^2 - 2) mod (2^p - 1)
	// 2^p - 1 is prime iff s_{p-2} = 0

	m := new(big.Int).Exp(big.NewInt(2), p, nil)
	m.Sub(m, big.NewInt(1))

	s := big.NewInt(4)
	two := big.NewInt(2)

	for range pInt - 2 {
		s.Mul(s, s)
		s.Sub(s, two)
		s.Mod(s, m)
	}

	return s.Sign() == 0
}
