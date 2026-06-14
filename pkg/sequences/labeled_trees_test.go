package sequences

import (
	"math/big"
	"testing"
)

func TestGetLabeledTreesSequence(t *testing.T) {
	// OEIS A000272: 1, 1, 1, 3, 16, 125, 1296, 16807, 262144, 4782969, 100000000
	expected := []int64{1, 1, 1, 3, 16, 125, 1296, 16807}

	maxN := big.NewInt(int64(len(expected) - 1))
	seq, err := GetLabeledTreesSequence(maxN, false)
	if err != nil {
		t.Fatalf("Error generating sequence: %v", err)
	}

	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, val := range expected {
		if seq.Sequence[i].Cmp(big.NewInt(val)) != 0 {
			t.Errorf("At position %d: expected %d, got %s", i, val, seq.Sequence[i].String())
		}
	}
}

func TestGetLabeledTreesAtPosition(t *testing.T) {
	tests := []struct {
		n        int64
		expected string
	}{
		{0, "1"},
		{1, "1"},
		{2, "1"},
		{3, "3"},
		{4, "16"},
		{5, "125"},
		{10, "100000000"}, // 10^(10-2) = 10^8
	}

	for _, tt := range tests {
		seq, err := GetLabeledTreesSequence(big.NewInt(tt.n), true)
		if err != nil {
			t.Errorf("Error for n=%d: %v", tt.n, err)
			continue
		}
		if seq.Result.String() != tt.expected {
			t.Errorf("For n=%d: expected %s, got %s", tt.n, tt.expected, seq.Result.String())
		}
	}
}
