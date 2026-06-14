package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetCollatzSequence(t *testing.T) {
	tests := []struct {
		name       string
		n          int64
		isPosition bool
		wantSeq    []*big.Int
	}{
		{
			name:       "Collatz sequence for 6",
			n:          6,
			isPosition: false,
			wantSeq: []*big.Int{
				big.NewInt(6),
				big.NewInt(3),
				big.NewInt(10),
				big.NewInt(5),
				big.NewInt(16),
				big.NewInt(8),
				big.NewInt(4),
				big.NewInt(2),
				big.NewInt(1),
			},
		},
		{
			name:       "Find number with sequence length 4",
			n:          4,
			isPosition: true,
			wantSeq: []*big.Int{
				big.NewInt(8),
				big.NewInt(4),
				big.NewInt(2),
				big.NewInt(1),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCollatzSequence(tt.n, tt.isPosition)
			if err != nil {
				t.Fatalf("GetCollatzSequence() error = %v", err)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetCollatzSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
		})
	}
}
