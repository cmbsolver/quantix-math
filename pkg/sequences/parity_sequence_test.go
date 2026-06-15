package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetParitySequence(t *testing.T) {
	tests := []struct {
		name       string
		maxNumber  *big.Int
		positional bool
		wantSeq    []*big.Int
		wantErr    bool
	}{
		{
			name:       "First 10 terms",
			maxNumber:  big.NewInt(9),
			positional: false,
			wantSeq: []*big.Int{
				big.NewInt(0), big.NewInt(1), big.NewInt(0), big.NewInt(1), big.NewInt(0),
				big.NewInt(1), big.NewInt(0), big.NewInt(1), big.NewInt(0), big.NewInt(1),
			},
			wantErr: false,
		},
		{
			name:       "Position 0",
			maxNumber:  big.NewInt(0),
			positional: true,
			wantSeq:    []*big.Int{big.NewInt(0)},
			wantErr:    false,
		},
		{
			name:       "Position 1",
			maxNumber:  big.NewInt(1),
			positional: true,
			wantSeq:    []*big.Int{big.NewInt(1)},
			wantErr:    false,
		},
		{
			name:       "Position 10",
			maxNumber:  big.NewInt(10),
			positional: true,
			wantSeq:    []*big.Int{big.NewInt(0)},
			wantErr:    false,
		},
		{
			name:       "Negative number",
			maxNumber:  big.NewInt(-1),
			positional: false,
			wantSeq:    nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetParitySequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetParitySequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
					t.Errorf("GetParitySequence() = %v, want %v", got.Sequence, tt.wantSeq)
				}
			}
		})
	}
}
