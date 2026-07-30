package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestMersennePrimeExponentsOEISExample(t *testing.T) {
	seq, err := GenerateMersennePrimeExponentsSequence(big.NewInt(31))
	if err != nil {
		t.Fatalf("GenerateMersennePrimeExponentsSequence() error = %v", err)
	}

	wantExponents := []*big.Int{
		big.NewInt(2),
		big.NewInt(3),
		big.NewInt(5),
		big.NewInt(7),
		big.NewInt(13),
		big.NewInt(17),
		big.NewInt(19),
		big.NewInt(31),
	}

	if !reflect.DeepEqual(seq.Sequence, wantExponents) {
		t.Fatalf("exponents = %v, want %v", seq.Sequence, wantExponents)
	}

	wantMersennePrimes := []string{"3", "7", "31", "127", "8191", "131071", "524287", "2147483647"}
	for i, p := range seq.Sequence {
		m := new(big.Int).Exp(big.NewInt(2), p, nil)
		m.Sub(m, big.NewInt(1))
		if m.String() != wantMersennePrimes[i] {
			t.Fatalf("2^%v - 1 = %v, want %v", p, m, wantMersennePrimes[i])
		}
	}
}

func TestGenerateMersennePrimeExponentsSequence(t *testing.T) {
	tests := []struct {
		maxNumber *big.Int
		expected  []*big.Int
	}{
		{
			maxNumber: big.NewInt(10),
			expected: []*big.Int{
				big.NewInt(2),
				big.NewInt(3),
				big.NewInt(5),
				big.NewInt(7),
			},
		},
		{
			maxNumber: big.NewInt(31),
			expected: []*big.Int{
				big.NewInt(2),
				big.NewInt(3),
				big.NewInt(5),
				big.NewInt(7),
				big.NewInt(13),
				big.NewInt(17),
				big.NewInt(19),
				big.NewInt(31),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.maxNumber.String(), func(t *testing.T) {
			got, err := GenerateMersennePrimeExponentsSequence(tt.maxNumber)
			if err != nil {
				t.Errorf("GenerateMersennePrimeExponentsSequence() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got.Sequence, tt.expected) {
				t.Errorf("GenerateMersennePrimeExponentsSequence() = %v, want %v", got.Sequence, tt.expected)
			}
		})
	}
}

func TestGetMersennePrimeExponentsAtPosition(t *testing.T) {
	tests := []struct {
		n        *big.Int
		expected *big.Int
	}{
		{big.NewInt(1), big.NewInt(2)},
		{big.NewInt(2), big.NewInt(3)},
		{big.NewInt(3), big.NewInt(5)},
		{big.NewInt(4), big.NewInt(7)},
		{big.NewInt(5), big.NewInt(13)},
		{big.NewInt(6), big.NewInt(17)},
		{big.NewInt(7), big.NewInt(19)},
		{big.NewInt(8), big.NewInt(31)},
	}

	for _, tt := range tests {
		t.Run(tt.n.String(), func(t *testing.T) {
			got, err := GetMersennePrimeExponentsAtPosition(tt.n)
			if err != nil {
				t.Errorf("GetMersennePrimeExponentsAtPosition() error = %v", err)
				return
			}
			if got.Result.Cmp(tt.expected) != 0 {
				t.Errorf("GetMersennePrimeExponentsAtPosition() = %v, want %v", got.Result, tt.expected)
			}
		})
	}
}

func TestIsMersennePrime(t *testing.T) {
	tests := []struct {
		p        *big.Int
		expected bool
	}{
		{big.NewInt(2), true},
		{big.NewInt(3), true},
		{big.NewInt(5), true},
		{big.NewInt(7), true},
		{big.NewInt(11), false}, // 2^11 - 1 = 2047 = 23 * 89
		{big.NewInt(13), true},
		{big.NewInt(17), true},
		{big.NewInt(19), true},
		{big.NewInt(23), false},
		{big.NewInt(31), true},
	}

	for _, tt := range tests {
		t.Run(tt.p.String(), func(t *testing.T) {
			if got := IsMersennePrime(tt.p); got != tt.expected {
				t.Errorf("IsMersennePrime() = %v, want %v", got, tt.expected)
			}
		})
	}
}
