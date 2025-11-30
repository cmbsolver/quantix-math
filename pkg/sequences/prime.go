package sequences

import (
	"math/big"
	"strings"
)

// IsPrime checks if a number is prime.
func IsPrime(number *big.Int) bool {
	numberArray := strings.Split(number.String(), "")
	if len(numberArray) >= 2 {
		lastChar := numberArray[len(numberArray)-1]
		if lastChar == "0" || lastChar == "2" || lastChar == "4" || lastChar == "5" || lastChar == "6" || lastChar == "8" {
			return false
		}
	}

	if number.Cmp(big.NewInt(2)) < 0 {
		return false
	}
	if number.Cmp(big.NewInt(2)) == 0 || number.Cmp(big.NewInt(3)) == 0 {
		return true
	}

	if number.Cmp(big.NewInt(3)) > 0 && number.Cmp(big.NewInt(158981)) <= 0 {
		return IsNumberInPrimeList(number.Int64())
	}

	if new(big.Int).Mod(number, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 ||
		new(big.Int).Mod(number, big.NewInt(3)).Cmp(big.NewInt(0)) == 0 {
		return false
	}

	// Start checking with 6k ± 1
	sqrt := new(big.Int).Sqrt(number)
	k := big.NewInt(5)
	for k.Cmp(sqrt) <= 0 {
		if new(big.Int).Mod(number, k).Cmp(big.NewInt(0)) == 0 ||
			new(big.Int).Mod(number, new(big.Int).Add(k, big.NewInt(2))).Cmp(big.NewInt(0)) == 0 {
			return false
		}
		k.Add(k, big.NewInt(6))
	}

	return true
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
