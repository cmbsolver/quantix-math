package sequences

import (
	"math/big"
	"testing"
)

func TestGetA000028Sequence(t *testing.T) {
	// Example terms from OEIS A000028
	expected := []int64{2, 3, 4, 5, 7, 9, 11, 13, 16, 17, 19, 23, 24, 25, 29, 30, 31, 37, 40, 41, 42, 43, 47, 49, 53, 54, 56, 59, 60, 61, 66, 67, 70, 71, 72, 73, 78, 79, 81, 83, 84, 88, 89, 90, 96}

	limit := big.NewInt(int64(len(expected)))
	seq, err := GetA000028Sequence(limit, false)
	if err != nil {
		t.Fatalf("Error generating sequence: %v", err)
	}

	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected %d terms, got %d", len(expected), len(seq.Sequence))
	}

	for i, val := range seq.Sequence {
		if val.Int64() != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], val.Int64())
		}
	}
}

func TestGetA000028AtPosition(t *testing.T) {
	// Test a specific term: 96 is the 45th term
	pos := big.NewInt(45)
	seq, err := GetA000028Sequence(pos, true)
	if err != nil {
		t.Fatalf("Error getting sequence at position: %v", err)
	}

	if seq.Result.Int64() != 96 {
		t.Errorf("Expected 96 at position 45, got %d", seq.Result.Int64())
	}
}
