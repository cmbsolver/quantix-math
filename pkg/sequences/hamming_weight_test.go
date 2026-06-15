package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateHammingWeightSequence(t *testing.T) {
	maxNumber := big.NewInt(15)
	expected := []int64{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4}

	seq, err := GenerateHammingWeightSequence(maxNumber)
	if err != nil {
		t.Fatalf("Failed to generate Hamming Weight sequence: %v", err)
	}

	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, val := range seq.Sequence {
		if val.Int64() != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], val.Int64())
		}
	}
}

func TestGetHammingWeightAtPosition(t *testing.T) {
	tests := []struct {
		position int64
		expected int64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{7, 3},
		{15, 4},
		{1023, 10},
	}

	for _, tt := range tests {
		pos := big.NewInt(tt.position)
		seq, err := GetHammingWeightAtPosition(pos)
		if err != nil {
			t.Errorf("Failed to get Hamming Weight at position %d: %v", tt.position, err)
			continue
		}
		if seq.Result.Int64() != tt.expected {
			t.Errorf("At position %d: expected %d, got %d", tt.position, tt.expected, seq.Result.Int64())
		}
	}
}
