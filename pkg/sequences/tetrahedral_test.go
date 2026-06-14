package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetTetrahedralSequence(t *testing.T) {
	tests := []struct {
		name         string
		n            *big.Int
		isPositional bool
		wantName     string
		wantSequence []*big.Int
		wantErr      bool
	}{
		{
			name:         "first terms up to 35",
			n:            big.NewInt(35),
			isPositional: false,
			wantName:     "Tetrahedral numbers (A000292)",
			wantSequence: []*big.Int{
				big.NewInt(0),
				big.NewInt(1),
				big.NewInt(4),
				big.NewInt(10),
				big.NewInt(20),
				big.NewInt(35),
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
			name:         "positional 2",
			n:            big.NewInt(2),
			isPositional: true,
			wantSequence: []*big.Int{big.NewInt(4)},
			wantErr:      false,
		},
		{
			name:         "positional 3",
			n:            big.NewInt(3),
			isPositional: true,
			wantSequence: []*big.Int{big.NewInt(10)},
			wantErr:      false,
		},
		{
			name:         "positional 4",
			n:            big.NewInt(4),
			isPositional: true,
			wantSequence: []*big.Int{big.NewInt(20)},
			wantErr:      false,
		},
		{
			name:         "positional 10",
			n:            big.NewInt(10),
			isPositional: true,
			wantSequence: []*big.Int{big.NewInt(220)},
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
			got, err := GetTetrahedralSequence(tt.n, tt.isPositional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTetrahedralSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSequence) {
				t.Errorf("GetTetrahedralSequence() Sequence = %v, want %v", got.Sequence, tt.wantSequence)
			}
		})
	}
}
