package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    *big.Int
		want bool
	}{
		{big.NewInt(2), true},
		{big.NewInt(3), true},
		{big.NewInt(4), false},
		{big.NewInt(17), true},
		{big.NewInt(100), false},
	}
	for _, tt := range tests {
		t.Run(tt.n.String(), func(t *testing.T) {
			if got := IsPrime(tt.n); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPrimeSequence(t *testing.T) {
	tests := []struct {
		name       string
		maxNumber  *big.Int
		positional bool
		wantSeq    []*big.Int
	}{
		{
			name:       "Primes up to 10",
			maxNumber:  big.NewInt(10),
			positional: false,
			wantSeq: []*big.Int{
				big.NewInt(2),
				big.NewInt(3),
				big.NewInt(5),
				big.NewInt(7),
			},
		},
		{
			name:       "4th prime (index 3, but code uses counter starting at 0)",
			maxNumber:  big.NewInt(3),
			positional: true,
			wantSeq:    []*big.Int{big.NewInt(7)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPrimeSequence(tt.maxNumber, tt.positional)
			if err != nil {
				t.Fatalf("GetPrimeSequence() error = %v", err)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetPrimeSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
		})
	}
}

func TestIsSemiPrime(t *testing.T) {
	tests := []struct {
		n    *big.Int
		want bool
	}{
		{big.NewInt(4), true},  // 2*2
		{big.NewInt(6), true},  // 2*3
		{big.NewInt(9), true},  // 3*3
		{big.NewInt(10), true}, // 2*5
		{big.NewInt(8), false}, // 2*2*2
		{big.NewInt(7), false}, // prime
	}
	for _, tt := range tests {
		t.Run(tt.n.String(), func(t *testing.T) {
			if got := IsSemiPrime(tt.n); got != tt.want {
				t.Errorf("IsSemiPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
