package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGenerateSqrt3ConvergentsA002531Sequence(t *testing.T) {
	tests := []struct {
		maxNumber *big.Int
		want      []*big.Int
	}{
		{
			maxNumber: big.NewInt(0),
			want:      []*big.Int{big.NewInt(1)},
		},
		{
			maxNumber: big.NewInt(1),
			want:      []*big.Int{big.NewInt(1), big.NewInt(1)},
		},
		{
			maxNumber: big.NewInt(10),
			want: []*big.Int{
				big.NewInt(1),   // a(0)
				big.NewInt(1),   // a(1)
				big.NewInt(2),   // a(2)
				big.NewInt(5),   // a(3)
				big.NewInt(7),   // a(4)
				big.NewInt(19),  // a(5)
				big.NewInt(26),  // a(6)
				big.NewInt(71),  // a(7)
				big.NewInt(97),  // a(8)
				big.NewInt(265), // a(9)
				big.NewInt(362), // a(10)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.maxNumber.String(), func(t *testing.T) {
			got, err := GenerateSqrt3ConvergentsA002531Sequence(tt.maxNumber)
			if err != nil {
				t.Errorf("GenerateSqrt3ConvergentsA002531Sequence() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got.Sequence, tt.want) {
				t.Errorf("GenerateSqrt3ConvergentsA002531Sequence() = %v, want %v", got.Sequence, tt.want)
			}
		})
	}
}

func TestGetSqrt3ConvergentsA002531AtPosition(t *testing.T) {
	tests := []struct {
		n    *big.Int
		want *big.Int
	}{
		{big.NewInt(0), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(2), big.NewInt(2)},
		{big.NewInt(5), big.NewInt(19)},
		{big.NewInt(10), big.NewInt(362)},
	}
	for _, tt := range tests {
		t.Run(tt.n.String(), func(t *testing.T) {
			got, err := GetSqrt3ConvergentsA002531AtPosition(tt.n)
			if err != nil {
				t.Errorf("GetSqrt3ConvergentsA002531AtPosition() error = %v", err)
				return
			}
			if got.Result.Cmp(tt.want) != 0 {
				t.Errorf("GetSqrt3ConvergentsA002531AtPosition() = %v, want %v", got.Result, tt.want)
			}
		})
	}
}
