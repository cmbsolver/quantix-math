package sequences

import (
	"math/big"
	"testing"
)

func TestGetLabeledRootedTreesSequence(t *testing.T) {
	tests := []struct {
		n        int64
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{3, "9"},
		{4, "64"},
		{5, "625"},
		{6, "7776"},
		{7, "117649"},
		{8, "2097152"},
		{9, "43046721"},
		{10, "1000000000"},
	}

	for _, tt := range tests {
		n := big.NewInt(tt.n)
		seq, err := GetLabeledRootedTreesSequence(n, true)
		if err != nil {
			t.Errorf("GetLabeledRootedTreesSequence(%d, true) error = %v", tt.n, err)
			continue
		}
		if seq.Result.String() != tt.expected {
			t.Errorf("GetLabeledRootedTreesSequence(%d, true) = %v, expected %v", tt.n, seq.Result.String(), tt.expected)
		}
	}
}

func TestGenerateLabeledRootedTreesSequence(t *testing.T) {
	n := big.NewInt(5)
	seq, err := GetLabeledRootedTreesSequence(n, false)
	if err != nil {
		t.Fatalf("GetLabeledRootedTreesSequence(5, false) error = %v", err)
	}

	expected := []string{"1", "2", "9", "64", "625"}
	if len(seq.Sequence) != len(expected) {
		t.Fatalf("expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, v := range seq.Sequence {
		if v.String() != expected[i] {
			t.Errorf("at index %d, expected %s, got %s", i, expected[i], v.String())
		}
	}
}
