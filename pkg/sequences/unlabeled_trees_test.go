package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateUnlabeledTreesSequence(t *testing.T) {
	// Values from OEIS A000055: 1, 1, 1, 1, 2, 3, 6, 11, 23, 47, 106, 235
	expected := []int64{1, 1, 1, 1, 2, 3, 6, 11, 23, 47, 106, 235}
	maxN := int64(len(expected) - 1)

	seq, err := GenerateUnlabeledTreesSequence(big.NewInt(maxN))
	if err != nil {
		t.Fatalf("Failed to generate sequence: %v", err)
	}

	if len(seq.Sequence) != len(expected) {
		t.Errorf("Expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, val := range expected {
		if seq.Sequence[i].Int64() != val {
			t.Errorf("At index %d: expected %d, got %s", i, val, seq.Sequence[i].String())
		}
	}
}

func TestGetUnlabeledTreesAtPosition(t *testing.T) {
	// A000055(11) = 235
	n := big.NewInt(11)
	expected := int64(235)

	seq, err := GetUnlabeledTreesAtPosition(n)
	if err != nil {
		t.Fatalf("Failed to get value at position: %v", err)
	}

	if seq.Result.Int64() != expected {
		t.Errorf("Expected %d, got %s", expected, seq.Result.String())
	}
}
