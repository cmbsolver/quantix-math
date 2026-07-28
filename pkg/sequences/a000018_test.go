package sequences

import (
	"math/big"
	"testing"
)

func TestCalculateA000018(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 4},
		{5, 8},
		{6, 13},
		{7, 25},
		{8, 44},
		{9, 83},
		{10, 152},
	}

	for _, tt := range tests {
		result := CalculateA000018(tt.n)
		if result != tt.expected {
			t.Errorf("CalculateA000018(%d) = %d; expected %d", tt.n, result, tt.expected)
		}
	}
}

func TestGetA000018Sequence(t *testing.T) {
	maxNumber := big.NewInt(5)
	seq, err := GetA000018Sequence(maxNumber, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedValues := []int64{1, 1, 2, 2, 4}
	if len(seq.Sequence) != len(expectedValues) {
		t.Fatalf("Expected sequence length %d; got %d", len(expectedValues), len(seq.Sequence))
	}

	for i, val := range expectedValues {
		if seq.Sequence[i].Int64() != val {
			t.Errorf("Sequence[%d] = %d; expected %d", i, seq.Sequence[i].Int64(), val)
		}
	}
}
