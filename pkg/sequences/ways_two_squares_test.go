package sequences

import (
	"math/big"
	"testing"
)

func TestGetWaysTwoSquaresSequence(t *testing.T) {
	// A002654: 1, 1, 0, 1, 2, 0, 0, 1, 1, 2, 0, 0, 2, 0, 0, 1, 2, 1, 0, 2
	expected := []int64{1, 1, 0, 1, 2, 0, 0, 1, 1, 2, 0, 0, 2, 0, 0, 1, 2, 1, 0, 2}

	maxNumber := big.NewInt(20)
	seq, err := GetWaysTwoSquaresSequence(maxNumber, false)
	if err != nil {
		t.Fatalf("Error generating sequence: %v", err)
	}

	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, val := range seq.Sequence {
		if val.Int64() != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], val.Int64())
		}
	}
}

func TestGetWaysTwoSquaresAtPosition(t *testing.T) {
	tests := []struct {
		n        int64
		expected int64
	}{
		{1, 1},
		{2, 1},
		{4, 1},
		{5, 2},
		{10, 2},
		{13, 2},
		{17, 2},
		{18, 1},
		{25, 3},
	}

	for _, tt := range tests {
		n := big.NewInt(tt.n)
		seq, err := GetWaysTwoSquaresSequence(n, true)
		if err != nil {
			t.Errorf("Error for n=%d: %v", tt.n, err)
			continue
		}
		if seq.Result.Int64() != tt.expected {
			t.Errorf("For n=%d: expected %d, got %d", tt.n, tt.expected, seq.Result.Int64())
		}
	}
}
