package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetEulerZigzagSequence(t *testing.T) {
	tests := []struct {
		maxNumber  *big.Int
		positional bool
		expected   []*big.Int
	}{
		{
			maxNumber:  big.NewInt(200),
			positional: false,
			expected: []*big.Int{
				big.NewInt(1),
				big.NewInt(1),
				big.NewInt(1),
				big.NewInt(2),
				big.NewInt(5),
				big.NewInt(16),
				big.NewInt(61),
			},
		},
		{
			maxNumber:  big.NewInt(7),
			positional: true,
			expected: []*big.Int{
				big.NewInt(272),
			},
		},
	}

	for _, tt := range tests {
		seq, err := GetEulerZigzagSequence(tt.maxNumber, tt.positional)
		if err != nil {
			t.Errorf("GetEulerZigzagSequence(%v, %v) error = %v", tt.maxNumber, tt.positional, err)
			continue
		}
		if !reflect.DeepEqual(seq.Sequence, tt.expected) {
			t.Errorf("GetEulerZigzagSequence(%v, %v) sequence = %v, expected %v", tt.maxNumber, tt.positional, seq.Sequence, tt.expected)
		}
	}
}
