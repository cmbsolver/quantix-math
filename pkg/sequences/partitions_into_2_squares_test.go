package sequences

import (
	"math/big"
	"testing"
)

func TestGeneratePartitionsInto2SquaresSequence(t *testing.T) {
	// A000161 sequence: 1, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 2, ...
	expected := []int64{1, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 2}
	maxNumber := big.NewInt(int64(len(expected) - 1))

	seq, err := GeneratePartitionsInto2SquaresSequence(maxNumber)
	if err != nil {
		t.Fatalf("Error generating sequence: %v", err)
	}

	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, v := range expected {
		if seq.Sequence[i].Int64() != v {
			t.Errorf("At index %d: expected %d, got %d", i, v, seq.Sequence[i].Int64())
		}
	}
}

func TestGetPartitionsInto2SquaresAtPosition(t *testing.T) {
	tests := []struct {
		n        int64
		expected int64
	}{
		{0, 1}, // 0^2 + 0^2 = 0
		{1, 1}, // 0^2 + 1^2 = 1
		{2, 1}, // 1^2 + 1^2 = 2
		{3, 0},
		{4, 1},  // 0^2 + 2^2 = 4
		{25, 2}, // 0^2 + 5^2 = 25, 3^2 + 4^2 = 25
		{50, 2}, // 1^2 + 7^2 = 50, 5^2 + 5^2 = 50
		{65, 2}, // 1^2 + 8^2 = 65, 4^2 + 7^2 = 65
	}

	for _, tt := range tests {
		seq, err := GetPartitionsInto2SquaresAtPosition(big.NewInt(tt.n))
		if err != nil {
			t.Errorf("Error getting value at position %d: %v", tt.n, err)
			continue
		}
		if seq.Result.Int64() != tt.expected {
			t.Errorf("At position %d: expected %d, got %d", tt.n, tt.expected, seq.Result.Int64())
		}
	}
}
