package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetPentagonalSequence(t *testing.T) {
	tests := []struct {
		name         string
		n            *big.Int
		isPositional bool
		wantName     string
		wantSequence []*big.Int
		wantErr      bool
	}{
		{
			name:         "first 10 terms",
			n:            big.NewInt(117),
			isPositional: false,
			wantName:     "Pentagonal numbers (A000326)",
			wantSequence: []*big.Int{
				big.NewInt(0),
				big.NewInt(1),
				big.NewInt(5),
				big.NewInt(12),
				big.NewInt(22),
				big.NewInt(35),
				big.NewInt(51),
				big.NewInt(70),
				big.NewInt(92),
				big.NewInt(117),
			},
			wantErr: false,
		},
		{
			name:         "positional 0",
			n:            big.NewInt(0),
			isPositional: true,
			wantSequence: []*big.Int{big.NewInt(0)},
			wantErr:      false,
		},
		{
			name:         "positional 1",
			n:            big.NewInt(1),
			isPositional: true,
			wantSequence: []*big.Int{big.NewInt(1)},
			wantErr:      false,
		},
		{
			name:         "positional 5",
			n:            big.NewInt(5),
			isPositional: true,
			wantSequence: []*big.Int{big.NewInt(35)},
			wantErr:      false,
		},
		{
			name:         "positional 10",
			n:            big.NewInt(10),
			isPositional: true,
			wantSequence: []*big.Int{big.NewInt(145)},
			wantErr:      false,
		},
		{
			name:         "invalid position -1",
			n:            big.NewInt(-1),
			isPositional: true,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPentagonalSequence(tt.n, tt.isPositional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPentagonalSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSequence) {
				t.Errorf("GetPentagonalSequence() Sequence = %v, want %v", got.Sequence, tt.wantSequence)
			}
		})
	}
}
