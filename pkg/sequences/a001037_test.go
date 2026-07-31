package sequences

import (
	"math/big"
	"testing"
)

func TestCalculateA001037(t *testing.T) {
	// A001037: 1, 2, 1, 2, 3, 6, 9, 18
	expected := []int64{1, 2, 1, 2, 3, 6, 9, 18}

	for i, exp := range expected {
		result := calculateA001037(int64(i))
		if result.Cmp(big.NewInt(exp)) != 0 {
			t.Errorf("calculateA001037(%d) = %v; want %v", i, result, exp)
		}
	}
}

func TestGenerateA001037Sequence(t *testing.T) {
	max := big.NewInt(20)
	seq, err := GetA001037Sequence(max, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []int64{1, 2, 1, 2, 3, 6, 9, 18}
	if len(seq.Sequence) != len(expected) {
		t.Fatalf("expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, val := range seq.Sequence {
		if val.Cmp(big.NewInt(expected[i])) != 0 {
			t.Errorf("sequence at %d = %v; want %v", i, val, expected[i])
		}
	}
}

func TestGetA001037AtPosition(t *testing.T) {
	pos := big.NewInt(5)
	seq, err := GetA001037Sequence(pos, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := int64(6)
	if seq.Result.Cmp(big.NewInt(expected)) != 0 {
		t.Errorf("expected result %d, got %v", expected, seq.Result)
	}
}
