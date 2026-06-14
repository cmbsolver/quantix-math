package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetTotientSequence(t *testing.T) {
	tests := []struct {
		name      string
		maxNumber *big.Int
		wantSeq   []*big.Int
		wantLen   int64
	}{
		{
			name:      "Totient of 10",
			maxNumber: big.NewInt(10),
			wantSeq: []*big.Int{
				big.NewInt(1),
				big.NewInt(3),
				big.NewInt(7),
				big.NewInt(9),
			},
			wantLen: 4,
		},
		{
			name:      "Totient of 1",
			maxNumber: big.NewInt(1),
			wantSeq: []*big.Int{
				big.NewInt(1),
			},
			wantLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTotientSequence(tt.maxNumber)
			if err != nil {
				t.Fatalf("GetTotientSequence() error = %v", err)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetTotientSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
			if got.Result.Int64() != tt.wantLen {
				t.Errorf("GetTotientSequence() Result = %v, want %v", got.Result, tt.wantLen)
			}
		})
	}
}

func TestGetTotientPrimeSequence(t *testing.T) {
	tests := []struct {
		name      string
		maxNumber *big.Int
		wantSeq   []*big.Int
		wantLen   int64
	}{
		{
			name:      "Totient primes of 10",
			maxNumber: big.NewInt(10),
			wantSeq: []*big.Int{
				big.NewInt(3),
				big.NewInt(7),
			},
			wantLen: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTotientPrimeSequence(tt.maxNumber)
			if err != nil {
				t.Fatalf("GetTotientPrimeSequence() error = %v", err)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetTotientPrimeSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
			if got.Result.Int64() != tt.wantLen {
				t.Errorf("GetTotientPrimeSequence() Result = %v, want %v", got.Result, tt.wantLen)
			}
		})
	}
}
