package sequences

import (
	"math"
	"strconv"
)

// IsPrime64 checks if a number is prime.
func IsPrime64(number int64) bool {
	if number < 2 {
		return false
	}
	if number == 2 {
		return true
	}
	if number%2 == 0 {
		return false
	}

	sqrt := int64(math.Sqrt(float64(number)))
	for i := int64(3); i <= sqrt; i += 2 {
		if number%i == 0 {
			return false
		}
	}

	return true
}

// IsEmirp64 checks if a number is an emirp (reversible prime).
func IsEmirp64(number int64) bool {
	if !IsPrime64(number) {
		return false
	}

	// Reverse the digits of the number
	reversedStr := reverseString(strconv.FormatInt(number, 10))
	reversedNumber, err := strconv.ParseInt(reversedStr, 10, 64)
	if err != nil {
		return false
	}

	return IsPrime64(reversedNumber)
}

// GetPrimeSequence64 generates the prime sequence.
func GetPrimeSequence64(maxNumber int64, isPositional bool) (*NumericSequence64, error) {
	numericSequence := &NumericSequence64{Name: "Prime", Number: maxNumber}
	numberToCalculate := maxNumber
	if isPositional {
		numberToCalculate = math.MaxInt64
	}
	counter := int64(0)

	for i := int64(0); i <= numberToCalculate; i++ {
		if IsPrime64(i) {
			if !isPositional {
				numericSequence.Sequence = append(numericSequence.Sequence, i)
			} else {
				if counter == maxNumber {
					numericSequence.Sequence = append(numericSequence.Sequence, i)
					break
				}
			}
			counter++
		}
	}

	return numericSequence, nil
}

// GetFibonacciPrimeSequence64 generates the Fibonacci prime sequence.
func GetFibonacciPrimeSequence64(maxNumber int64, isPositional bool) (*NumericSequence64, error) {
	numericSequence := &NumericSequence64{Name: "Fibonacci Prime", Number: maxNumber}
	numberToCalculate := maxNumber
	if isPositional {
		numberToCalculate = math.MaxInt64
	}

	a, b, c := int64(0), int64(1), int64(0)
	counter := int64(0)

	for c <= numberToCalculate {
		c = a + b
		a = b
		b = c

		if c <= numberToCalculate && IsPrime64(c) {
			if !isPositional {
				numericSequence.Sequence = append(numericSequence.Sequence, c)
			} else {
				if counter == maxNumber {
					numericSequence.Sequence = append(numericSequence.Sequence, c)
					break
				}
			}
			counter++
		}
	}

	return numericSequence, nil
}
