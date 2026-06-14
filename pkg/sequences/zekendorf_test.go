package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetFibonacciSequence(t *testing.T) {
	tests := []struct {
		name         string
		maxNumber    *big.Int
		isPositional bool
		wantSeq      []*big.Int
	}{
		{
			name:         "Fibonacci up to 10",
			maxNumber:    big.NewInt(10),
			isPositional: false,
			wantSeq: []*big.Int{
				big.NewInt(1),
				big.NewInt(1),
				big.NewInt(2),
				big.NewInt(3),
				big.NewInt(5),
				big.NewInt(8),
			},
		},
		{
			name:         "Fibonacci at position 5",
			maxNumber:    big.NewInt(5),
			isPositional: true,
			wantSeq:      []*big.Int{big.NewInt(5)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFibonacciSequence(tt.maxNumber, tt.isPositional)
			if err != nil {
				t.Fatalf("GetFibonacciSequence() error = %v", err)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetFibonacciSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
		})
	}
}

func TestGetZekendorfRepresentationSequence(t *testing.T) {
	tests := []struct {
		name      string
		maxNumber *big.Int
		wantSeq   []*big.Int
	}{
		{
			name:      "Zekendorf of 10",
			maxNumber: big.NewInt(10),
			wantSeq: []*big.Int{
				big.NewInt(8),
				big.NewInt(2),
			},
		},
		{
			name:      "Zekendorf of 100",
			maxNumber: big.NewInt(100),
			wantSeq: []*big.Int{
				big.NewInt(89),
				big.NewInt(8),
				big.NewInt(3),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetZekendorfRepresentationSequence(tt.maxNumber, false)
			if err != nil {
				t.Fatalf("GetZekendorfRepresentationSequence() error = %v", err)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetZekendorfRepresentationSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
		})
	}
}
