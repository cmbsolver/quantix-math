package sequences

import (
	"math/big"
	"testing"
)

func TestCalculateMotzkinNumbersA001006(t *testing.T) {
	// A001006: 1, 1, 2, 4, 9, 21, 51, 127
	expected := []int64{1, 1, 2, 4, 9, 21, 51, 127}

	for i, exp := range expected {
		result := calculateMotzkinNumbersA001006(i)
		if result.Cmp(big.NewInt(exp)) != 0 {
			t.Errorf("calculateMotzkinNumbersA001006(%d) = %v; want %v", i, result, exp)
		}
	}
}

func TestGenerateMotzkinNumbersA001006Sequence(t *testing.T) {
	max := big.NewInt(100)
	seq, err := GetMotzkinNumbersA001006Sequence(max, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := []int64{1, 1, 2, 4, 9, 21, 51}
	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, val := range seq.Sequence {
		if val.Cmp(big.NewInt(expected[i])) != 0 {
			t.Errorf("Sequence at %d = %v; want %v", i, val, expected[i])
		}
	}
}

func TestGetMotzkinNumbersA001006AtPosition(t *testing.T) {
	pos := big.NewInt(5)
	seq, err := GetMotzkinNumbersA001006Sequence(pos, true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := int64(21)
	if seq.Result.Cmp(big.NewInt(expected)) != 0 {
		t.Errorf("Expected result %d, got %v", expected, seq.Result)
	}
}
