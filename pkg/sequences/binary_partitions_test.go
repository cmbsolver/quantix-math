package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateBinaryPartitionsSequence(t *testing.T) {
	// A000123: 1, 2, 4, 6, 10, 14, 20, 26, 36, 46, 60, 74, 94, 114, 140, 166, 202, 238, 284, 330, 390
	expected := []int64{1, 2, 4, 6, 10, 14, 20, 26, 36, 46, 60, 74, 94, 114, 140, 166, 202, 238, 284, 330, 390}
	maxN := big.NewInt(20)

	seq, err := GenerateBinaryPartitionsSequence(maxN)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, val := range expected {
		if seq.Sequence[i].Int64() != val {
			t.Errorf("At index %d: expected %d, got %d", i, val, seq.Sequence[i].Int64())
		}
	}
}

func TestGetBinaryPartitionsAtPosition(t *testing.T) {
	tests := []struct {
		n        int64
		expected int64
	}{
		{0, 1},
		{1, 2},
		{2, 4},
		{3, 6},
		{4, 10},
		{5, 14},
		{10, 60},
		{20, 390},
	}

	for _, tt := range tests {
		seq, err := GetBinaryPartitionsAtPosition(big.NewInt(tt.n))
		if err != nil {
			t.Errorf("Unexpected error for n=%d: %v", tt.n, err)
			continue
		}
		if seq.Result.Int64() != tt.expected {
			t.Errorf("For n=%d: expected %d, got %d", tt.n, tt.expected, seq.Result.Int64())
		}
	}
}
