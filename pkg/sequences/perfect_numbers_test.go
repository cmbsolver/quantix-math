package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetPerfectNumbersSequence(t *testing.T) {
	tests := []struct {
		name         string
		n            *big.Int
		isPositional bool
		wantName     string
		wantSequence []*big.Int
		wantErr      bool
	}{
		{
			name:         "first 4 terms",
			n:            big.NewInt(10000),
			isPositional: false,
			wantName:     "Perfect numbers (A000396)",
			wantSequence: []*big.Int{
				big.NewInt(6),
				big.NewInt(28),
				big.NewInt(496),
				big.NewInt(8128),
			},
			wantErr: false,
		},
		{
			name:         "positional 1st",
			n:            big.NewInt(1),
			isPositional: true,
			wantName:     "Perfect numbers (A000396)",
			wantSequence: []*big.Int{big.NewInt(6)},
			wantErr:      false,
		},
		{
			name:         "positional 5th",
			n:            big.NewInt(5),
			isPositional: true,
			wantName:     "Perfect numbers (A000396)",
			wantSequence: []*big.Int{big.NewInt(33550336)},
			wantErr:      false,
		},
		{
			name:         "invalid position 0",
			n:            big.NewInt(0),
			isPositional: true,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPerfectNumbersSequence(tt.n, tt.isPositional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPerfectNumbersSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetPerfectNumbersSequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSequence) {
				t.Errorf("GetPerfectNumbersSequence() Sequence = %v, want %v", got.Sequence, tt.wantSequence)
			}
		})
	}
}
