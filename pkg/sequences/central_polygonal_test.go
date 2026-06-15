package sequences

import (
	"math/big"
	"testing"
)

func TestGetCentralPolygonalNumbersSequence(t *testing.T) {
	tests := []struct {
		maxNumber  *big.Int
		positional bool
		expected   []*big.Int
	}{
		{
			maxNumber:  big.NewInt(11),
			positional: false,
			expected: []*big.Int{
				big.NewInt(1),
				big.NewInt(2),
				big.NewInt(4),
				big.NewInt(7),
				big.NewInt(11),
			},
		},
		{
			maxNumber:  big.NewInt(0),
			positional: true,
			expected:   []*big.Int{big.NewInt(1)},
		},
		{
			maxNumber:  big.NewInt(1),
			positional: true,
			expected:   []*big.Int{big.NewInt(2)},
		},
		{
			maxNumber:  big.NewInt(2),
			positional: true,
			expected:   []*big.Int{big.NewInt(4)},
		},
		{
			maxNumber:  big.NewInt(3),
			positional: true,
			expected:   []*big.Int{big.NewInt(7)},
		},
		{
			maxNumber:  big.NewInt(4),
			positional: true,
			expected:   []*big.Int{big.NewInt(11)},
		},
	}

	for _, tt := range tests {
		seq, err := GetCentralPolygonalNumbersSequence(tt.maxNumber, tt.positional)
		if err != nil {
			t.Errorf("GetCentralPolygonalNumbersSequence(%v, %v) returned error: %v", tt.maxNumber, tt.positional, err)
			continue
		}

		if len(seq.Sequence) != len(tt.expected) {
			t.Errorf("GetCentralPolygonalNumbersSequence(%v, %v) length = %d, expected %d", tt.maxNumber, tt.positional, len(seq.Sequence), len(tt.expected))
			continue
		}

		for i, v := range seq.Sequence {
			if v.Cmp(tt.expected[i]) != 0 {
				t.Errorf("GetCentralPolygonalNumbersSequence(%v, %v) at index %d = %v, expected %v", tt.maxNumber, tt.positional, i, v, tt.expected[i])
			}
		}
	}
}
