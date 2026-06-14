package sequences

import (
	"math/big"
	"testing"
)

func TestGetUnlabeledDigraphsSequence(t *testing.T) {
	tests := []struct {
		n        int64
		expected *big.Int
	}{
		{0, big.NewInt(1)},
		{1, big.NewInt(1)},
		{2, big.NewInt(3)},
		{3, big.NewInt(16)},
		{4, big.NewInt(218)},
		{5, big.NewInt(9608)},
	}

	for _, tt := range tests {
		result, err := calculateA000273(tt.n)
		if err != nil {
			t.Errorf("calculateA000273(%d) returned error: %v", tt.n, err)
			continue
		}
		if result.Cmp(tt.expected) != 0 {
			t.Errorf("calculateA000273(%d) = %v; want %v", tt.n, result, tt.expected)
		}
	}
}

func TestUnlabeledDigraphsSequencePositional(t *testing.T) {
	maxNumber := big.NewInt(5)
	seq, err := GetUnlabeledDigraphsSequence(maxNumber, false)
	if err != nil {
		t.Fatalf("GenerateUnlabeledDigraphsSequence failed: %v", err)
	}

	expected := []int64{1, 1, 3, 16, 218, 9608}
	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, val := range seq.Sequence {
		if val.Cmp(big.NewInt(expected[i])) != 0 {
			t.Errorf("At index %d: expected %d, got %v", i, expected[i], val)
		}
	}
}
