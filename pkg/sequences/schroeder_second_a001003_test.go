package sequences

import (
	"math/big"
	"testing"
)

func TestCalculateSchroederSecondA001003(t *testing.T) {
	// A001003: 1, 1, 3, 11, 45, 197, 903, 4279
	expected := []int64{1, 1, 3, 11, 45, 197, 903, 4279}

	for i, exp := range expected {
		result := calculateSchroederSecondA001003(i)
		if result.Cmp(big.NewInt(exp)) != 0 {
			t.Errorf("calculateSchroederSecondA001003(%d) = %v; want %v", i, result, exp)
		}
	}
}

func TestGenerateSchroederSecondA001003Sequence(t *testing.T) {
	max := big.NewInt(100)
	seq, err := GetSchroederSecondA001003Sequence(max, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := []int64{1, 1, 3, 11, 45}
	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, val := range seq.Sequence {
		if val.Cmp(big.NewInt(expected[i])) != 0 {
			t.Errorf("Sequence at %d = %v; want %v", i, val, expected[i])
		}
	}
}

func TestGetSchroederSecondA001003AtPosition(t *testing.T) {
	pos := big.NewInt(5)
	seq, err := GetSchroederSecondA001003Sequence(pos, true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := int64(197)
	if seq.Result.Cmp(big.NewInt(expected)) != 0 {
		t.Errorf("Expected result %d, got %v", expected, seq.Result)
	}
}
