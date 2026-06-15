package sequences

import (
	"math/big"
	"testing"
)

func TestGetFreePolyominoesSequence(t *testing.T) {
	// A000105: 1, 1, 1, 2, 5, 12, 35, 108, 369, 1285
	expected := []int64{1, 1, 1, 2, 5, 12, 35, 108, 369}
	maxNumber := big.NewInt(int64(len(expected) - 1))

	seq, err := GetFreePolyominoesSequence(maxNumber, false)
	if err != nil {
		t.Fatalf("Failed to get sequence: %v", err)
	}

	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, val := range expected {
		if seq.Sequence[i].Int64() != val {
			t.Errorf("At index %d, expected %d, got %d", i, val, seq.Sequence[i].Int64())
		}
	}
}

func TestGetFreePolyominoesAtPosition(t *testing.T) {
	tests := []struct {
		n        int64
		expected int64
	}{
		{0, 1},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 5},
		{5, 12},
		{6, 35},
		{7, 108},
		{8, 369},
	}

	for _, tt := range tests {
		seq, err := GetFreePolyominoesSequence(big.NewInt(tt.n), true)
		if err != nil {
			t.Errorf("Failed to get position %d: %v", tt.n, err)
			continue
		}
		if seq.Result.Int64() != tt.expected {
			t.Errorf("At position %d, expected %d, got %d", tt.n, tt.expected, seq.Result.Int64())
		}
	}
}
