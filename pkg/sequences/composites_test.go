package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateCompositesSequence(t *testing.T) {
	tests := []struct {
		maxNumber int64
		expected  []int64
	}{
		{3, []int64{}},
		{4, []int64{4}},
		{10, []int64{4, 6, 8, 9, 10}},
		{20, []int64{4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20}},
	}

	for _, tt := range tests {
		max := big.NewInt(tt.maxNumber)
		seq, err := GenerateCompositesSequence(max)
		if err != nil {
			t.Errorf("GenerateCompositesSequence(%d) error: %v", tt.maxNumber, err)
			continue
		}

		if len(seq.Sequence) != len(tt.expected) {
			t.Errorf("GenerateCompositesSequence(%d) length = %d, want %d", tt.maxNumber, len(seq.Sequence), len(tt.expected))
			continue
		}

		for i, v := range seq.Sequence {
			if v.Int64() != tt.expected[i] {
				t.Errorf("GenerateCompositesSequence(%d) at index %d = %d, want %d", tt.maxNumber, i, v.Int64(), tt.expected[i])
			}
		}
	}
}

func TestGetCompositesAtPosition(t *testing.T) {
	tests := []struct {
		pos      int64
		expected int64
	}{
		{1, 4},
		{2, 6},
		{3, 8},
		{4, 9},
		{5, 10},
		{10, 18},
		{20, 32},
	}

	for _, tt := range tests {
		pos := big.NewInt(tt.pos)
		seq, err := GetCompositesAtPosition(pos)
		if err != nil {
			t.Errorf("GetCompositesAtPosition(%d) error: %v", tt.pos, err)
			continue
		}

		if seq.Result.Int64() != tt.expected {
			t.Errorf("GetCompositesAtPosition(%d) = %d, want %d", tt.pos, seq.Result.Int64(), tt.expected)
		}
	}
}
