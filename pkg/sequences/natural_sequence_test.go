package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetNaturalSequence(t *testing.T) {
	tests := []struct {
		name       string
		maxNumber  *big.Int
		positional bool
		wantSeq    []*big.Int
	}{
		{
			name:       "Natural sequence up to 5",
			maxNumber:  big.NewInt(5),
			positional: false,
			wantSeq: []*big.Int{
				big.NewInt(0),
				big.NewInt(1),
				big.NewInt(2),
				big.NewInt(3),
				big.NewInt(4),
				big.NewInt(5),
			},
		},
		{
			name:       "Natural number at position 10",
			maxNumber:  big.NewInt(10),
			positional: true,
			wantSeq:    []*big.Int{big.NewInt(10)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNaturalSequence(tt.maxNumber, tt.positional)
			if err != nil {
				t.Fatalf("GetNaturalSequence() error = %v", err)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetNaturalSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
		})
	}
}
