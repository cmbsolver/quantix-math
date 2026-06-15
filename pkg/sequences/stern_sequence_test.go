package sequences

import (
	"math/big"
	"testing"
)

func TestGetSternSequence(t *testing.T) {
	// a(0) to a(16)
	// 0, 1, 1, 2, 1, 3, 2, 3, 1, 4, 3, 5, 2, 5, 3, 4, 1
	expected := []int64{0, 1, 1, 2, 1, 3, 2, 3, 1, 4, 3, 5, 2, 5, 3, 4, 1}

	maxN := big.NewInt(16)
	seq, err := GetSternSequence(maxN, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(seq.Sequence) != 17 {
		t.Errorf("Expected length 17, got %d", len(seq.Sequence))
	}

	for i, val := range expected {
		if seq.Sequence[i].Int64() != val {
			t.Errorf("At index %d, expected %d, got %d", i, val, seq.Sequence[i].Int64())
		}
	}

	if seq.Result.Int64() != expected[16] {
		t.Errorf("Expected result %d, got %d", expected[16], seq.Result.Int64())
	}
}

func TestGetSternAtPosition(t *testing.T) {
	// a(91) = 19 from OEIS example
	pos := big.NewInt(91)
	expected := int64(19)

	seq, err := GetSternSequence(pos, true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if seq.Result.Int64() != expected {
		t.Errorf("At position 91, expected %d, got %d", expected, seq.Result.Int64())
	}
}

func TestIsSternNumber(t *testing.T) {
	tests := []struct {
		n      int64
		exists bool
	}{
		{0, true},
		{1, true},
		{5, true},
		{100, true},
		{-1, false},
	}

	for _, tt := range tests {
		exists, _ := IsSternNumber(big.NewInt(tt.n))
		if exists != tt.exists {
			t.Errorf("For n=%d, expected exists=%v, got %v", tt.n, tt.exists, exists)
		}
	}
}
