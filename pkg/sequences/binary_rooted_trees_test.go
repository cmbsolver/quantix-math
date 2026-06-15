package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGenerateBinaryRootedTreesA002572Sequence(t *testing.T) {
	tests := []struct {
		maxNumber *big.Int
		want      []*big.Int
	}{
		{
			maxNumber: big.NewInt(1),
			want:      []*big.Int{big.NewInt(1)},
		},
		{
			maxNumber: big.NewInt(5),
			want: []*big.Int{
				big.NewInt(1), // a(1)
				big.NewInt(1), // a(2)
				big.NewInt(1), // a(3)
				big.NewInt(2), // a(4)
				big.NewInt(3), // a(5)
			},
		},
		{
			maxNumber: big.NewInt(10),
			want: []*big.Int{
				big.NewInt(1),  // a(1)
				big.NewInt(1),  // a(2)
				big.NewInt(1),  // a(3)
				big.NewInt(2),  // a(4)
				big.NewInt(3),  // a(5)
				big.NewInt(5),  // a(6)
				big.NewInt(9),  // a(7)
				big.NewInt(16), // a(8)
				big.NewInt(28), // a(9)
				big.NewInt(50), // a(10)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.maxNumber.String(), func(t *testing.T) {
			got, err := GenerateBinaryRootedTreesA002572Sequence(tt.maxNumber)
			if err != nil {
				t.Errorf("GenerateBinaryRootedTreesA002572Sequence() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got.Sequence, tt.want) {
				t.Errorf("GenerateBinaryRootedTreesA002572Sequence() = %v, want %v", got.Sequence, tt.want)
			}
		})
	}
}

func TestGetBinaryRootedTreesA002572AtPosition(t *testing.T) {
	tests := []struct {
		n    *big.Int
		want *big.Int
	}{
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(4), big.NewInt(2)},
		{big.NewInt(5), big.NewInt(3)},
		{big.NewInt(10), big.NewInt(50)},
	}
	for _, tt := range tests {
		t.Run(tt.n.String(), func(t *testing.T) {
			got, err := GetBinaryRootedTreesA002572AtPosition(tt.n)
			if err != nil {
				t.Errorf("GetBinaryRootedTreesA002572AtPosition() error = %v", err)
				return
			}
			if got.Result.Cmp(tt.want) != 0 {
				t.Errorf("GetBinaryRootedTreesA002572AtPosition() = %v, want %v", got.Result, tt.want)
			}
		})
	}
}
