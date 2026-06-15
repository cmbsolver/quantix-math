package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGenerateSqrt3ConvergentsDenominatorsSequence(t *testing.T) {
	tests := []struct {
		maxNumber *big.Int
		want      []*big.Int
	}{
		{
			maxNumber: big.NewInt(11),
			want: []*big.Int{
				big.NewInt(0), big.NewInt(1), big.NewInt(1), big.NewInt(3),
				big.NewInt(4), big.NewInt(11), big.NewInt(15), big.NewInt(41),
				big.NewInt(56), big.NewInt(153), big.NewInt(209), big.NewInt(571),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.maxNumber.String(), func(t *testing.T) {
			got, err := GenerateSqrt3ConvergentsDenominatorsSequence(tt.maxNumber)
			if err != nil {
				t.Errorf("GenerateSqrt3ConvergentsDenominatorsSequence() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got.Sequence, tt.want) {
				t.Errorf("GenerateSqrt3ConvergentsDenominatorsSequence() = %v, want %v", got.Sequence, tt.want)
			}
		})
	}
}

func TestGetSqrt3ConvergentsDenominatorsAtPosition(t *testing.T) {
	tests := []struct {
		n    *big.Int
		want *big.Int
	}{
		{big.NewInt(0), big.NewInt(0)},
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(5), big.NewInt(11)},
		{big.NewInt(10), big.NewInt(209)},
		{big.NewInt(11), big.NewInt(571)},
	}
	for _, tt := range tests {
		t.Run(tt.n.String(), func(t *testing.T) {
			got, err := GetSqrt3ConvergentsDenominatorsAtPosition(tt.n)
			if err != nil {
				t.Errorf("GetSqrt3ConvergentsDenominatorsAtPosition() error = %v", err)
				return
			}
			if got.Result.Cmp(tt.want) != 0 {
				t.Errorf("GetSqrt3ConvergentsDenominatorsAtPosition() = %v, want %v", got.Result, tt.want)
			}
		})
	}
}
