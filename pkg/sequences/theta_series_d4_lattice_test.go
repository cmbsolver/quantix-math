package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateThetaSeriesD4LatticeSequence(t *testing.T) {
	// A004011: 1, 24, 24, 96, 24, 144, 96, 192, 24, 312, 144, 288, 96, 336, 192, 576
	expected := []int64{1, 24, 24, 96, 24, 144, 96, 192, 24, 312, 144, 288, 96, 336, 192, 576}

	maxN := big.NewInt(int64(len(expected) - 1))
	seq, err := GenerateThetaSeriesD4LatticeSequence(maxN)
	if err != nil {
		t.Fatalf("Error generating sequence: %v", err)
	}

	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, exp := range expected {
		if seq.Sequence[i].Cmp(big.NewInt(exp)) != 0 {
			t.Errorf("At index %d: expected %d, got %v", i, exp, seq.Sequence[i])
		}
	}
}

func TestGetThetaSeriesD4LatticeAtPosition(t *testing.T) {
	tests := []struct {
		n        int64
		expected int64
	}{
		{0, 1},
		{1, 24},
		{2, 24},
		{3, 96},
		{4, 24},
		{15, 576},
	}

	for _, tt := range tests {
		n := big.NewInt(tt.n)
		seq, err := GetThetaSeriesD4LatticeAtPosition(n)
		if err != nil {
			t.Errorf("Error getting term at position %d: %v", tt.n, err)
			continue
		}
		if seq.Result.Cmp(big.NewInt(tt.expected)) != 0 {
			t.Errorf("At position %d: expected %d, got %v", tt.n, tt.expected, seq.Result)
		}
	}
}
