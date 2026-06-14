package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetEulerNumbersSequence(t *testing.T) {
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
			n:            big.NewInt(100000),
			isPositional: false,
			wantName:     "Euler numbers (A000364)",
			wantSequence: []*big.Int{
				big.NewInt(1),
				big.NewInt(1),
				big.NewInt(5),
				big.NewInt(61),
				big.NewInt(1385),
				big.NewInt(50521),
			},
			wantErr: false,
		},
		{
			name:         "positional 0",
			n:            big.NewInt(0),
			isPositional: true,
			wantSequence: []*big.Int{big.NewInt(1)},
			wantErr:      false,
		},
		{
			name:         "positional 3",
			n:            big.NewInt(3),
			isPositional: true,
			wantSequence: []*big.Int{big.NewInt(61)},
			wantErr:      false,
		},
		{
			name:         "positional 7",
			n:            big.NewInt(7),
			isPositional: true,
			wantSequence: []*big.Int{big.NewInt(199360981)},
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
			got, err := GetEulerNumbersSequence(tt.n, tt.isPositional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEulerNumbersSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSequence) {
				t.Errorf("GetEulerNumbersSequence() Sequence = %v, want %v", got.Sequence, tt.wantSequence)
			}
		})
	}
}
