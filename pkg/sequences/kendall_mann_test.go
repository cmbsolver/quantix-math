package sequences

import (
	"math/big"
	"testing"
)

func TestGetKendallMannSequence(t *testing.T) {
	tests := []struct {
		n        int64
		expected *big.Int
	}{
		{1, big.NewInt(1)},
		{2, big.NewInt(1)},
		{3, big.NewInt(2)},
		{4, big.NewInt(6)},
		{5, big.NewInt(22)},
		{6, big.NewInt(101)},
		{7, big.NewInt(573)},
		{8, big.NewInt(3836)},
		{9, big.NewInt(29228)},
		{10, big.NewInt(250749)},
	}

	for _, tt := range tests {
		seq, err := GetKendallMannSequence(big.NewInt(tt.n), true)
		if err != nil {
			t.Errorf("GetKendallMannSequence(%d, true) error: %v", tt.n, err)
			continue
		}
		if seq.Result.Cmp(tt.expected) != 0 {
			t.Errorf("GetKendallMannSequence(%d, true) = %v; want %v", tt.n, seq.Result, tt.expected)
		}
	}
}

func TestGenerateKendallMannSequence(t *testing.T) {
	n := big.NewInt(5)
	seq, err := GetKendallMannSequence(n, false)
	if err != nil {
		t.Fatalf("GetKendallMannSequence(5, false) error: %v", err)
	}

	expected := []*big.Int{
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(6),
		big.NewInt(22),
	}

	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, val := range seq.Sequence {
		if val.Cmp(expected[i]) != 0 {
			t.Errorf("Sequence[%d] = %v; want %v", i, val, expected[i])
		}
	}
}
