package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetPartitionsDistinctSequence(t *testing.T) {
	tests := []struct {
		name       string
		maxNumber  *big.Int
		positional bool
		wantSeq    []*big.Int
		wantErr    bool
	}{
		{
			name:       "First 11 terms of A000009",
			maxNumber:  big.NewInt(10),
			positional: false,
			wantSeq: []*big.Int{
				big.NewInt(1),  // a(0)
				big.NewInt(1),  // a(1)
				big.NewInt(1),  // a(2)
				big.NewInt(2),  // a(3)
				big.NewInt(2),  // a(4)
				big.NewInt(3),  // a(5)
				big.NewInt(4),  // a(6)
				big.NewInt(5),  // a(7)
				big.NewInt(6),  // a(8)
				big.NewInt(8),  // a(9)
				big.NewInt(10), // a(10)
			},
			wantErr: false,
		},
		{
			name:       "Term at position 10",
			maxNumber:  big.NewInt(10),
			positional: true,
			wantSeq:    []*big.Int{big.NewInt(10)},
			wantErr:    false,
		},
		{
			name:       "Term at position 0",
			maxNumber:  big.NewInt(0),
			positional: true,
			wantSeq:    []*big.Int{big.NewInt(1)},
			wantErr:    false,
		},
		{
			name:       "Negative input",
			maxNumber:  big.NewInt(-1),
			positional: false,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPartitionsDistinctSequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPartitionsDistinctSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetPartitionsDistinctSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
		})
	}
}
