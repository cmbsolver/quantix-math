package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetNtoNSequence(t *testing.T) {
	tests := []struct {
		name         string
		n            *big.Int
		isPositional bool
		wantSequence []*big.Int
		wantErr      bool
	}{
		{
			name:         "first few terms",
			n:            big.NewInt(300),
			isPositional: false,
			wantSequence: []*big.Int{
				big.NewInt(1),   // 0^0
				big.NewInt(1),   // 1^1
				big.NewInt(4),   // 2^2
				big.NewInt(27),  // 3^3
				big.NewInt(256), // 4^4
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
			name:         "positional 2",
			n:            big.NewInt(2),
			isPositional: true,
			wantSequence: []*big.Int{big.NewInt(4)},
			wantErr:      false,
		},
		{
			name:         "positional 5",
			n:            big.NewInt(5),
			isPositional: true,
			wantSequence: []*big.Int{big.NewInt(3125)},
			wantErr:      false,
		},
		{
			name:         "invalid position",
			n:            big.NewInt(-1),
			isPositional: true,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNtoNSequence(tt.n, tt.isPositional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNtoNSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSequence) {
				t.Errorf("GetNtoNSequence() Sequence = %v, want %v", got.Sequence, tt.wantSequence)
			}
		})
	}
}
