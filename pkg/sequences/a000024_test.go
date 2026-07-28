package sequences

import (
	"math/big"
	"testing"
)

func TestCalculateA000024(t *testing.T) {
	tests := []struct {
		n    int
		want uint64
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 7},
		{5, 10},
		{6, 20},
		{7, 36},
		{8, 65},
		{9, 118},
		{10, 221},
	}

	for _, tt := range tests {
		if got := CalculateA000024(tt.n); got != tt.want {
			t.Errorf("CalculateA000024(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestGetA000024Sequence(t *testing.T) {
	maxNumber := big.NewInt(5)
	seq, err := GetA000024Sequence(maxNumber, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := []int64{1, 1, 2, 2, 7}
	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, v := range expected {
		if seq.Sequence[i].Int64() != v {
			t.Errorf("Sequence[%d] = %d, want %d", i, seq.Sequence[i].Int64(), v)
		}
	}

	if seq.Result.Int64() != 7 {
		t.Errorf("Result = %d, want 7", seq.Result.Int64())
	}
}

func TestGetA000024AtPosition(t *testing.T) {
	n := big.NewInt(4)
	seq, err := GetA000024Sequence(n, true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if seq.Result.Int64() != 7 {
		t.Errorf("Result at position 4 = %d, want 7", seq.Result.Int64())
	}
}
