package sequences

import (
	"math/big"
	"testing"
)

func TestA000022(t *testing.T) {
	expected := []int64{0, 1, 0, 1, 1, 2, 2, 6, 9, 20, 37, 86, 181, 422, 943, 2223}
	limit := len(expected) - 1
	
	seq, err := GetA000022Sequence(big.NewInt(int64(limit)), false)
	if err != nil {
		t.Fatalf("Error generating A000022: %v", err)
	}
	
	if len(seq.Sequence) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(seq.Sequence))
	}
	
	for i, val := range seq.Sequence {
		if val.Int64() != expected[i] {
			t.Errorf("A000022(%d) expected %d, got %d", i, expected[i], val.Int64())
		}
	}
}

func TestA000022AtPosition(t *testing.T) {
	tests := []struct {
		n        int64
		expected int64
	}{
		{0, 0},
		{1, 1},
		{2, 0},
		{3, 1},
		{4, 1},
		{5, 2},
		{7, 6},
		{10, 37},
	}
	
	for _, tt := range tests {
		seq, err := GetA000022Sequence(big.NewInt(tt.n), true)
		if err != nil {
			t.Errorf("Error generating A000022 at position %d: %v", tt.n, err)
			continue
		}
		if seq.Result.Int64() != tt.expected {
			t.Errorf("A000022(%d) expected %d, got %d", tt.n, tt.expected, seq.Result.Int64())
		}
	}
}
