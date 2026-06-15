package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetPrimesA000040Sequence(t *testing.T) {
	tests := []struct {
		name       string
		maxNumber  *big.Int
		positional bool
		wantSeq    []*big.Int
	}{
		{
			name:       "Primes up to 20",
			maxNumber:  big.NewInt(20),
			positional: false,
			wantSeq: []*big.Int{
				big.NewInt(2), big.NewInt(3), big.NewInt(5), big.NewInt(7),
				big.NewInt(11), big.NewInt(13), big.NewInt(17), big.NewInt(19),
			},
		},
		{
			name:       "10th prime",
			maxNumber:  big.NewInt(10),
			positional: true,
			wantSeq: []*big.Int{
				big.NewInt(29),
			},
		},
		{
			name:       "1st prime",
			maxNumber:  big.NewInt(1),
			positional: true,
			wantSeq: []*big.Int{
				big.NewInt(2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPrimesA000040Sequence(tt.maxNumber, tt.positional)
			if err != nil {
				t.Fatalf("GetPrimesA000040Sequence() error = %v", err)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetPrimesA000040Sequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
		})
	}
}
