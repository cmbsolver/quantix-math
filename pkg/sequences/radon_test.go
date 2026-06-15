package sequences

import (
	"math/big"
	"testing"
)

func TestCalculateRadon(t *testing.T) {
	tests := []struct {
		n        int64
		expected int64
	}{
		{1, 1},    // k=0, a=0, b=0 -> 2^0 + 8*0 = 1
		{2, 2},    // k=1, a=0, b=1 -> 2^1 + 8*0 = 2
		{3, 1},    // k=0, a=0, b=0 -> 1
		{4, 4},    // k=2, a=0, b=2 -> 4
		{5, 1},    // k=0 -> 1
		{6, 2},    // k=1 -> 2
		{7, 1},    // k=0 -> 1
		{8, 8},    // k=3, a=0, b=3 -> 8
		{9, 1},    // k=0 -> 1
		{10, 2},   // k=1 -> 2
		{11, 1},   // k=0 -> 1
		{12, 4},   // k=2 -> 4
		{13, 1},   // k=0 -> 1
		{14, 2},   // k=1 -> 2
		{15, 1},   // k=0 -> 1
		{16, 9},   // k=4, a=1, b=0 -> 2^0 + 8*1 = 9
		{32, 10},  // k=5, a=1, b=1 -> 2^1 + 8*1 = 10
		{64, 12},  // k=6, a=1, b=2 -> 2^2 + 8*1 = 12
		{128, 16}, // k=7, a=1, b=3 -> 2^3 + 8*1 = 16
		{256, 17}, // k=8, a=2, b=0 -> 2^0 + 8*2 = 17
	}

	for _, tt := range tests {
		result := CalculateRadon(tt.n)
		if result.Int64() != tt.expected {
			t.Errorf("CalculateRadon(%d) = %d; expected %d", tt.n, result.Int64(), tt.expected)
		}
	}
}

func TestGenerateRadonSequence(t *testing.T) {
	maxNum := big.NewInt(16)
	seq, err := GenerateRadonSequence(maxNum)
	if err != nil {
		t.Fatalf("GenerateRadonSequence failed: %v", err)
	}

	expected := []int64{1, 2, 1, 4, 1, 2, 1, 8, 1, 2, 1, 4, 1, 2, 1, 9}
	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, v := range seq.Sequence {
		if v.Int64() != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], v.Int64())
		}
	}
}
