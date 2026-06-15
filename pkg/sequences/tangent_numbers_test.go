package sequences

import (
	"math/big"
	"testing"
)

func TestGetTangentNumbersSequence(t *testing.T) {
	tests := []struct {
		maxNumber *big.Int
		expected  []*big.Int
	}{
		{big.NewInt(1), []*big.Int{big.NewInt(1)}},
		{big.NewInt(2), []*big.Int{big.NewInt(1), big.NewInt(2)}},
		{big.NewInt(16), []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(16)}},
		{big.NewInt(300), []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(16), big.NewInt(272)}},
		{big.NewInt(8000), []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(16), big.NewInt(272), big.NewInt(7936)}},
	}

	for _, tt := range tests {
		seq, err := GetTangentNumbersSequence(tt.maxNumber, false)
		if err != nil {
			t.Errorf("GetTangentNumbersSequence(%v) error: %v", tt.maxNumber, err)
			continue
		}
		if len(seq.Sequence) != len(tt.expected) {
			t.Errorf("GetTangentNumbersSequence(%v) len = %d, expected %d", tt.maxNumber, len(seq.Sequence), len(tt.expected))
			continue
		}
		for i, v := range seq.Sequence {
			if v.Cmp(tt.expected[i]) != 0 {
				t.Errorf("GetTangentNumbersSequence(%v) [%d] = %v, expected %v", tt.maxNumber, i, v, tt.expected[i])
			}
		}
	}
}

func TestGetTangentNumberAtPosition(t *testing.T) {
	tests := []struct {
		pos      *big.Int
		expected *big.Int
	}{
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(2), big.NewInt(2)},
		{big.NewInt(3), big.NewInt(16)},
		{big.NewInt(4), big.NewInt(272)},
		{big.NewInt(5), big.NewInt(7936)},
	}

	for _, tt := range tests {
		seq, err := GetTangentNumbersSequence(tt.pos, true)
		if err != nil {
			t.Errorf("GetTangentNumberAtPosition(%v) error: %v", tt.pos, err)
			continue
		}
		if len(seq.Sequence) != 1 {
			t.Errorf("GetTangentNumberAtPosition(%v) len = %d, expected 1", tt.pos, len(seq.Sequence))
			continue
		}
		if seq.Sequence[0].Cmp(tt.expected) != 0 {
			t.Errorf("GetTangentNumberAtPosition(%v) = %v, expected %v", tt.pos, seq.Sequence[0], tt.expected)
		}
	}
}
