package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetPellSequence(t *testing.T) {
	tests := []struct {
		maxNumber  *big.Int
		positional bool
		want       []*big.Int
	}{
		{big.NewInt(10), false, []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(5)}},
		{big.NewInt(30), false, []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(5), big.NewInt(12), big.NewInt(29)}},
		{big.NewInt(0), true, []*big.Int{big.NewInt(0)}},
		{big.NewInt(1), true, []*big.Int{big.NewInt(1)}},
		{big.NewInt(2), true, []*big.Int{big.NewInt(2)}},
		{big.NewInt(3), true, []*big.Int{big.NewInt(5)}},
		{big.NewInt(4), true, []*big.Int{big.NewInt(12)}},
		{big.NewInt(5), true, []*big.Int{big.NewInt(29)}},
		{big.NewInt(6), true, []*big.Int{big.NewInt(70)}},
	}
	for _, tt := range tests {
		t.Run(tt.maxNumber.String(), func(t *testing.T) {
			got, err := GetPellSequence(tt.maxNumber, tt.positional)
			if err != nil {
				t.Errorf("GetPellSequence() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got.Sequence, tt.want) {
				t.Errorf("GetPellSequence() = %v, want %v", got.Sequence, tt.want)
			}
		})
	}
}
