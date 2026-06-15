package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetFactorialSequence(t *testing.T) {
	tests := []struct {
		maxNumber    *big.Int
		isPositional bool
		wantName     string
		wantSequence []*big.Int
		wantResult   *big.Int
		wantErr      bool
	}{
		{
			maxNumber:    big.NewInt(120),
			isPositional: false,
			wantName:     "Factorial numbers (A000142)",
			wantSequence: []*big.Int{
				big.NewInt(1),   // 0!
				big.NewInt(1),   // 1!
				big.NewInt(2),   // 2!
				big.NewInt(6),   // 3!
				big.NewInt(24),  // 4!
				big.NewInt(120), // 5!
			},
			wantResult: big.NewInt(120),
			wantErr:    false,
		},
		{
			maxNumber:    big.NewInt(5),
			isPositional: true,
			wantName:     "Factorial numbers (A000142)",
			wantSequence: []*big.Int{
				big.NewInt(120), // 5!
			},
			wantResult: big.NewInt(120),
			wantErr:    false,
		},
		{
			maxNumber:    big.NewInt(0),
			isPositional: true,
			wantName:     "Factorial numbers (A000142)",
			wantSequence: []*big.Int{
				big.NewInt(1), // 0!
			},
			wantResult: big.NewInt(1),
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.wantName, func(t *testing.T) {
			got, err := GetFactorialSequence(tt.maxNumber, tt.isPositional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFactorialSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetFactorialSequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSequence) {
				t.Errorf("GetFactorialSequence() Sequence = %v, want %v", got.Sequence, tt.wantSequence)
			}
			if tt.wantResult != nil && got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetFactorialSequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}

func TestCalculateFactorial(t *testing.T) {
	tests := []struct {
		n    int64
		want *big.Int
	}{
		{0, big.NewInt(1)},
		{1, big.NewInt(1)},
		{2, big.NewInt(2)},
		{3, big.NewInt(6)},
		{4, big.NewInt(24)},
		{5, big.NewInt(120)},
		{6, big.NewInt(720)},
	}
	for _, tt := range tests {
		if got := CalculateFactorial(tt.n); got.Cmp(tt.want) != 0 {
			t.Errorf("CalculateFactorial(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
