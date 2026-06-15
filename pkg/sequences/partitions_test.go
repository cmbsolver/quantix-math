package sequences

import (
	"math/big"
	"testing"
)

func TestGeneratePartitionsSequence(t *testing.T) {
	tests := []struct {
		n        int64
		expected []int64
	}{
		{0, []int64{1}},
		{5, []int64{1, 1, 2, 3, 5, 7}},
		{10, []int64{1, 1, 2, 3, 5, 7, 11, 15, 22, 30, 42}},
	}

	for _, tt := range tests {
		seq, err := GeneratePartitionsSequence(big.NewInt(tt.n))
		if err != nil {
			t.Errorf("GeneratePartitionsSequence(%d) error: %v", tt.n, err)
			continue
		}
		if int64(len(seq.Sequence)) != tt.n+1 {
			t.Errorf("GeneratePartitionsSequence(%d) length = %d, want %d", tt.n, len(seq.Sequence), tt.n+1)
		}
		for i, val := range seq.Sequence {
			if val.Int64() != tt.expected[i] {
				t.Errorf("GeneratePartitionsSequence(%d) at index %d = %d, want %d", tt.n, i, val.Int64(), tt.expected[i])
			}
		}
	}
}

func TestGetPartitionsAtPosition(t *testing.T) {
	tests := []struct {
		n        int64
		expected int64
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 7},
		{10, 42},
		{20, 627},
	}

	for _, tt := range tests {
		seq, err := GetPartitionsAtPosition(big.NewInt(tt.n))
		if err != nil {
			t.Errorf("GetPartitionsAtPosition(%d) error: %v", tt.n, err)
			continue
		}
		if seq.Result.Int64() != tt.expected {
			t.Errorf("GetPartitionsAtPosition(%d) = %d, want %d", tt.n, seq.Result.Int64(), tt.expected)
		}
	}
}
