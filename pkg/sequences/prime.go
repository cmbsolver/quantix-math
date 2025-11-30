package sequences

import (
	"math/big"
)

// IsPrime checks if a number is prime.
func IsPrime(n *big.Int) bool {
	return n.ProbablyPrime(20)
}

// IsEmirp checks if a number is an emirp (reversible prime).
func IsEmirp(number *big.Int) bool {
	if !IsPrime(number) {
		return false
	}

	// Reverse the digits of the number
	reversedStr := reverseString(number.String())
	reversedNumber := new(big.Int)
	reversedNumber, ok := reversedNumber.SetString(reversedStr, 10)
	if !ok {
		return false
	}

	return IsPrime(reversedNumber)
}

// reverseString reverses a string.
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// GetPrimeSequence generates the prime sequence.
func GetPrimeSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	numericSequence := &NumericSequence{Name: "Prime", Number: new(big.Int).Set(maxNumber)}
	numberToCalculate := new(big.Int).Set(maxNumber)
	if isPositional {
		numberToCalculate = new(big.Int).SetUint64(^uint64(0)) // Max uint64 value
	}
	counter := big.NewInt(0)

	for i := big.NewInt(0); i.Cmp(numberToCalculate) <= 0; i.Add(i, big.NewInt(1)) {
		if IsPrime(i) {
			if !isPositional {
				numericSequence.Sequence = append(numericSequence.Sequence, new(big.Int).Set(i))
			} else {
				if counter.Cmp(maxNumber) == 0 {
					numericSequence.Sequence = append(numericSequence.Sequence, new(big.Int).Set(i))
					break
				}
			}
			counter.Add(counter, big.NewInt(1))
		}
	}

	return numericSequence, nil
}

// GetFibonacciPrimeSequence generates the Fibonacci prime sequence.
func GetFibonacciPrimeSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	numericSequence := &NumericSequence{Name: "Fibonacci Prime", Number: new(big.Int).Set(maxNumber)}
	numberToCalculate := new(big.Int).Set(maxNumber)
	if isPositional {
		numberToCalculate = new(big.Int).SetUint64(^uint64(0)) // Max uint64 value
	}

	a, b, c := big.NewInt(0), big.NewInt(1), big.NewInt(0)
	counter := big.NewInt(0)

	for c.Cmp(numberToCalculate) <= 0 {
		c.Add(a, b)
		a.Set(b)
		b.Set(c)

		if c.Cmp(numberToCalculate) <= 0 && IsPrime(c) {
			if !isPositional {
				numericSequence.Sequence = append(numericSequence.Sequence, new(big.Int).Set(c))
			} else {
				if counter.Cmp(maxNumber) == 0 {
					numericSequence.Sequence = append(numericSequence.Sequence, new(big.Int).Set(c))
					break
				}
			}
			counter.Add(counter, big.NewInt(1))
		}
	}

	return numericSequence, nil
}

// IsSemiPrime checks if a number is a semi-prime (product of exactly two primes).
// It takes a *big.Int as a parameter.
func IsSemiPrime(n *big.Int) bool {
	if n.Cmp(big.NewInt(4)) < 0 {
		return false
	}

	factors := 0
	tempN := new(big.Int).Set(n)
	i := big.NewInt(2)
	two := big.NewInt(2)
	one := big.NewInt(1)
	zero := big.NewInt(0)

	// Optimization: check small factors first or standard trial division
	// Since we don't have a complex factorization library, we will do trial division
	// up to sqrt(n).

	limit := new(big.Int).Sqrt(tempN)

	for i.Cmp(limit) <= 0 {
		remainder := new(big.Int)
		quotient := new(big.Int)

		quotient.DivMod(tempN, i, remainder)

		for remainder.Cmp(zero) == 0 {
			factors++
			tempN.Set(quotient)
			if factors > 2 {
				return false
			}
			// Re-calculate limit/remainder for the new tempN
			limit.Sqrt(tempN)
			quotient.DivMod(tempN, i, remainder)
		}

		if i.Cmp(two) == 0 {
			i.Add(i, one)
		} else {
			i.Add(i, two)
		}
	}

	if tempN.Cmp(one) > 0 {
		factors++
	}

	return factors == 2
}

// IsCircularPrime checks if a number is a circular prime.
func IsCircularPrime(n *big.Int) bool {
	if !IsPrime(n) {
		return false
	}

	s := n.String()
	length := len(s)

	if length == 1 {
		return true
	}

	for _, char := range s {
		if (char-'0')%2 == 0 || char == '5' {
			return false
		}
	}

	for i := 1; i < length; i++ {
		s = s[1:] + s[0:1]
		rotatedN := new(big.Int)
		rotatedN.SetString(s, 10)

		if !IsPrime(rotatedN) {
			return false
		}
	}

	return true
}
