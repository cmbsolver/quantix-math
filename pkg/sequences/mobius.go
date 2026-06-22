package sequences

import (
	"math"
	"slices"
)

// MobiusCalculator precomputes Mobius values up to a maximum bound.
type MobiusCalculator struct {
	mu   []int8
	maxN int
}

// NewMobiusCalculator creates and initializes a MobiusCalculator with precomputed values up to maxN.
func NewMobiusCalculator(maxN int) *MobiusCalculator {
	mu := make([]int8, maxN+1)
	for i := range mu {
		mu[i] = 1
	}

	primes := make([]bool, maxN+1)
	for i := range primes {
		primes[i] = true
	}

	for i := 2; i <= maxN; i++ {
		if primes[i] {
			for j := i; j <= maxN; j += i {
				if j > i {
					primes[j] = false
				}
				mu[j] *= -1
				if (j/i)%i == 0 {
					mu[j] = 0
				}
			}
		}
	}

	return &MobiusCalculator{
		mu:   mu,
		maxN: maxN,
	}
}

// GetMu returns the Mobius value for n.
func (mc *MobiusCalculator) GetMu(n int) int8 {
	if n < 1 || n > mc.maxN {
		return 0 // Or handle as error/calculate on the fly for small out of bounds
	}
	return mc.mu[n]
}

// DirectEvaluation returns Mobius values for a list of integers.
func (mc *MobiusCalculator) DirectEvaluation(nums []int) []int8 {
	results := make([]int8, len(nums))
	for i, n := range nums {
		results[i] = mc.GetMu(n)
	}
	return results
}

// SequentialIndexMasking multiplies each element at 1-based index i by mu(i).
func (mc *MobiusCalculator) SequentialIndexMasking(input []float64) []float64 {
	output := make([]float64, len(input))
	for i, val := range input {
		index := i + 1
		if index <= mc.maxN {
			output[i] = val * float64(mc.mu[index])
		} else {
			output[i] = 0 // Or keep original? The requirement says multiplied by mu(i)
		}
	}
	return output
}

// DivisorSummation calculates the sum of mu(d) for all divisors d of n.
func (mc *MobiusCalculator) DivisorSummation(n int) int {
	if n < 1 {
		return 0
	}
	divisors := mc.GetDivisors(n)
	sum := 0
	for _, d := range divisors {
		sum += int(mc.GetMu(d))
	}
	return sum
}

// MobiusInversionStep calculates sum_{d|n} mu(n/d) * g(d).
// g is represented as a map or we can use a callback.
func (mc *MobiusCalculator) MobiusInversionStep(n int, g func(int) float64) float64 {
	if n < 1 {
		return 0
	}
	divisors := mc.GetDivisors(n)
	sum := 0.0
	for _, d := range divisors {
		sum += float64(mc.GetMu(n/d)) * g(d)
	}
	return sum
}

// GetDivisors returns all divisors of n.
func (mc *MobiusCalculator) GetDivisors(n int) []int {
	var divisors []int
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i*i != n {
				divisors = append(divisors, n/i)
			}
		}
	}
	slices.Sort(divisors)
	return divisors
}
