package sequences

import (
	"math/big"
	"testing"
)

func TestSchroederFourth(t *testing.T) {
	// A000311: 0, 1, 1, 4, 26, 236, 2752, 39208, 660032
	expected := []int64{0, 1, 1, 4, 26, 236, 2752, 39208, 660032}

	for i, exp := range expected {
		result := calculateSchroederFourth(i)
		if result.Cmp(big.NewInt(exp)) != 0 {
			t.Errorf("calculateSchroederFourth(%d) = %v; want %v", i, result, exp)
		}
	}
}

func TestGetSchroederFourthSequence(t *testing.T) {
	max := big.NewInt(100)
	seq, err := GetSchroederFourthSequence(max, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedSeq := []int64{0, 1, 1, 4, 26}
	if len(seq.Sequence) != len(expectedSeq) {
		t.Errorf("Expected sequence length %d, got %d", len(expectedSeq), len(seq.Sequence))
	}

	for i, val := range seq.Sequence {
		if val.Cmp(big.NewInt(expectedSeq[i])) != 0 {
			t.Errorf("Sequence at %d = %v; want %v", i, val, expectedSeq[i])
		}
	}
}

func TestGetSchroederFourthAtPosition(t *testing.T) {
	pos := big.NewInt(5)
	seq, err := GetSchroederFourthSequence(pos, true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := int64(236)
	if seq.Result.Cmp(big.NewInt(expected)) != 0 {
		t.Errorf("Expected result %d, got %v", expected, seq.Result)
	}
}
