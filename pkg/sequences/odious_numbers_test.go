package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGenerateOdiousNumbersSequence(t *testing.T) {
	tests := []struct {
		maxNumber *big.Int
		want      []int64
	}{
		{big.NewInt(10), []int64{1, 2, 4, 7, 8}},
		{big.NewInt(1), []int64{1}},
		{big.NewInt(0), nil},
		{big.NewInt(16), []int64{1, 2, 4, 7, 8, 11, 13, 14, 16}},
	}
	for _, tt := range tests {
		t.Run(tt.maxNumber.String(), func(t *testing.T) {
			got, err := GenerateOdiousNumbersSequence(tt.maxNumber)
			if err != nil {
				t.Errorf("GenerateOdiousNumbersSequence() error = %v", err)
				return
			}
			var gotValues []int64
			for _, v := range got.Sequence {
				gotValues = append(gotValues, v.Int64())
			}
			if !reflect.DeepEqual(gotValues, tt.want) {
				t.Errorf("GenerateOdiousNumbersSequence() = %v, want %v", gotValues, tt.want)
			}
		})
	}
}

func TestGetOdiousNumberAtPosition(t *testing.T) {
	tests := []struct {
		n    *big.Int
		want int64
	}{
		{big.NewInt(1), 1},
		{big.NewInt(2), 2},
		{big.NewInt(3), 4},
		{big.NewInt(4), 7},
		{big.NewInt(5), 8},
		{big.NewInt(6), 11},
		{big.NewInt(10), 19},
	}
	for _, tt := range tests {
		t.Run(tt.n.String(), func(t *testing.T) {
			got, err := GetOdiousNumberAtPosition(tt.n)
			if err != nil {
				t.Errorf("GetOdiousNumberAtPosition() error = %v", err)
				return
			}
			if got.Result.Int64() != tt.want {
				t.Errorf("GetOdiousNumberAtPosition() = %v, want %v", got.Result.Int64(), tt.want)
			}
		})
	}
}

func TestIsOdious(t *testing.T) {
	tests := []struct {
		n    *big.Int
		want bool
	}{
		{big.NewInt(1), true},   // 1
		{big.NewInt(2), true},   // 10
		{big.NewInt(3), false},  // 11
		{big.NewInt(4), true},   // 100
		{big.NewInt(7), true},   // 111
		{big.NewInt(8), true},   // 1000
		{big.NewInt(15), false}, // 1111
		{big.NewInt(0), false},
	}
	for _, tt := range tests {
		t.Run(tt.n.String(), func(t *testing.T) {
			if got := IsOdious(tt.n); got != tt.want {
				t.Errorf("IsOdious() = %v, want %v", got, tt.want)
			}
		})
	}
}
