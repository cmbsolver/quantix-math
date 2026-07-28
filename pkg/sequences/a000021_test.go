package sequences

import (
	"math/big"
	"testing"
)

func TestCalculateA000021(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 6},
		{5, 9},
		{6, 17},
		{7, 30},
		{8, 54},
		{9, 98},
		{10, 183},
	}

	for _, tt := range tests {
		if got := CalculateA000021(tt.n); got != tt.want {
			t.Errorf("CalculateA000021(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

func TestGetA000021Sequence(t *testing.T) {
	maxNumber := big.NewInt(5)
	seq, err := GetA000021Sequence(maxNumber, false)
	if err != nil {
		t.Fatalf("GetA000021Sequence failed: %v", err)
	}

	expected := []int64{1, 1, 2, 2, 6}
	if len(seq.Sequence) != len(expected) {
		t.Fatalf("expected sequence length %d; got %d", len(expected), len(seq.Sequence))
	}

	for i, val := range seq.Sequence {
		if val.Int64() != expected[i] {
			t.Errorf("sequence[%d] = %d; want %d", i, val.Int64(), expected[i])
		}
	}
}
