package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGenerateRootedUnlabeledTreesSequence(t *testing.T) {
	tests := []struct {
		maxNumber *big.Int
		want      []int64
	}{
		{big.NewInt(0), []int64{0}},
		{big.NewInt(1), []int64{0, 1}},
		{big.NewInt(10), []int64{0, 1, 1, 2, 4, 9, 20, 48, 115, 286, 719}},
	}
	for _, tt := range tests {
		t.Run(tt.maxNumber.String(), func(t *testing.T) {
			got, err := GenerateRootedUnlabeledTreesSequence(tt.maxNumber)
			if err != nil {
				t.Errorf("GenerateRootedUnlabeledTreesSequence() error = %v", err)
				return
			}
			var gotInts []int64
			for _, val := range got.Sequence {
				gotInts = append(gotInts, val.Int64())
			}
			if !reflect.DeepEqual(gotInts, tt.want) {
				t.Errorf("GenerateRootedUnlabeledTreesSequence() = %v, want %v", gotInts, tt.want)
			}
		})
	}
}

func TestGetRootedUnlabeledTreesAtPosition(t *testing.T) {
	tests := []struct {
		n    *big.Int
		want int64
	}{
		{big.NewInt(0), 0},
		{big.NewInt(1), 1},
		{big.NewInt(5), 9},
		{big.NewInt(10), 719},
	}
	for _, tt := range tests {
		t.Run(tt.n.String(), func(t *testing.T) {
			got, err := GetRootedUnlabeledTreesAtPosition(tt.n)
			if err != nil {
				t.Errorf("GetRootedUnlabeledTreesAtPosition() error = %v", err)
				return
			}
			if got.Result.Int64() != tt.want {
				t.Errorf("GetRootedUnlabeledTreesAtPosition() = %v, want %v", got.Result.Int64(), tt.want)
			}
		})
	}
}
