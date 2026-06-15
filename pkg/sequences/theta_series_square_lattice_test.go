package sequences

import (
	"math/big"
	"testing"
)

func TestThetaSeriesSquareLattice(t *testing.T) {
	// A004018: 1, 4, 4, 0, 4, 8, 0, 0, 4, 4, 8, 0, 0, 8, 0, 0, 4, 8, 4, 0, 8, 0, 0, 0, 0, 12, 8, 0, 0, 8, 0, 0, 4, 0, 8, 0, 4, 8, 0, 0, 8, 8, 0, 0, 0, 8, 0, 0, 0, 4, 12, 0, 8, 8, 0, 0, 0, 0, 8, 0, 0, 8, 0, 0, 4, 16, 0, 0, 8, 0, 0, 0, 4, 8, 8, 0, 0, 0, 0, 0, 8, 4, 8, 0, 0, 16, 0, 0, 0, 8, 8, 0, 0, 0, 0, 0, 0, 8, 4
	expected := []int64{1, 4, 4, 0, 4, 8, 0, 0, 4, 4, 8, 0, 0, 8, 0, 0, 4, 8, 4, 0, 8, 0, 0, 0, 0, 12, 8, 0, 0, 8, 0, 0, 4, 0, 8, 0, 4, 8, 0, 0, 8, 8, 0, 0, 0, 8, 0, 0, 0, 4, 12, 0, 8, 8, 0, 0, 0, 0, 8, 0, 0, 8, 0, 0, 4, 16, 0, 0, 8, 0, 0, 0, 4, 8, 8, 0, 0, 0, 0, 0, 8, 4, 8, 0, 0, 16, 0, 0, 0, 8, 8, 0, 0, 0, 0, 0, 0, 8, 4}

	maxN := int64(len(expected) - 1)
	seq, err := GenerateThetaSeriesSquareLatticeSequence(big.NewInt(maxN))
	if err != nil {
		t.Fatalf("Failed to generate sequence: %v", err)
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

func TestThetaSeriesSquareLatticeAtPosition(t *testing.T) {
	testCases := []struct {
		n        int64
		expected int64
	}{
		{0, 1},
		{1, 4},
		{2, 4},
		{3, 0},
		{4, 4},
		{5, 8},
		{25, 12},
		{65, 16},
	}

	for _, tc := range testCases {
		seq, err := GetThetaSeriesSquareLatticeAtPosition(big.NewInt(tc.n))
		if err != nil {
			t.Errorf("Failed to get value at position %d: %v", tc.n, err)
			continue
		}
		if seq.Result.Int64() != tc.expected {
			t.Errorf("At position %d: expected %d, got %d", tc.n, tc.expected, seq.Result.Int64())
		}
	}
}
