package sequences

import (
	"math/big"
	"testing"
)

func TestPlanePartitions(t *testing.T) {
	// Values from OEIS A000219: 1, 1, 3, 6, 13, 24, 48, 86, 160, 282, 500, 859, 1479
	expected := []int64{1, 1, 3, 6, 13, 24, 48, 86, 160, 282, 500, 859, 1479}

	maxNum := big.NewInt(int64(len(expected) - 1))
	seq, err := GeneratePlanePartitionsSequence(maxNum)
	if err != nil {
		t.Fatalf("Failed to generate plane partitions: %v", err)
	}

	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, val := range expected {
		if seq.Sequence[i].Int64() != val {
			t.Errorf("At index %d: expected %d, got %d", i, val, seq.Sequence[i].Int64())
		}
	}
}

func TestPlanePartitionsAtPosition(t *testing.T) {
	n := big.NewInt(10)
	expected := int64(500) // a(10) = 500

	seq, err := GetPlanePartitionsAtPosition(n)
	if err != nil {
		t.Fatalf("Failed to get plane partition at position: %v", err)
	}

	if seq.Result.Int64() != expected {
		t.Errorf("Expected %d, got %d", expected, seq.Result.Int64())
	}
}
