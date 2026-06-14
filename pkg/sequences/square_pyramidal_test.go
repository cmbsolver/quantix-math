package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetSquarePyramidalSequence(t *testing.T) {
	tests := []struct {
		name         string
		n            *big.Int
		isPositional bool
		wantName     string
		wantSequence []*big.Int
		wantErr      bool
	}{
		{
			name:         "first 6 terms",
			n:            big.NewInt(55),
			isPositional: false,
			wantName:     "Square pyramidal numbers (A000330)",
			wantSequence: []*big.Int{
				big.NewInt(0),
				big.NewInt(1),
				big.NewInt(5),
				big.NewInt(14),
				big.NewInt(30),
				big.NewInt(55),
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
			name:         "positional 4",
			n:            big.NewInt(4),
			isPositional: true,
			wantSequence: []*big.Int{big.NewInt(30)},
			wantErr:      false,
		},
		{
			name:         "positional 10",
			n:            big.NewInt(10),
			isPositional: true,
			wantSequence: []*big.Int{big.NewInt(385)},
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
			got, err := GetSquarePyramidalSequence(tt.n, tt.isPositional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSquarePyramidalSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSequence) {
				t.Errorf("GetSquarePyramidalSequence() Sequence = %v, want %v", got.Sequence, tt.wantSequence)
			}
		})
	}
}
