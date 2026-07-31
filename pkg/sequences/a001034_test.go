package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateA001034Sequence(t *testing.T) {
	// OEIS A001034 example terms from the sequence page.
	expected := []int64{60, 168, 360, 504, 660, 1092, 2448, 2520, 3420, 4080}

	seq, err := GenerateA001034Sequence(big.NewInt(4080))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(seq.Sequence) != len(expected) {
		t.Fatalf("expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, want := range expected {
		if seq.Sequence[i].Cmp(big.NewInt(want)) != 0 {
			t.Errorf("term %d = %v; want %d", i+1, seq.Sequence[i], want)
		}
	}
}

func TestGetA001034AtPosition(t *testing.T) {
	seq, err := GetA001034AtPosition(big.NewInt(36))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if seq.Result.Cmp(big.NewInt(178920)) != 0 {
		t.Fatalf("unexpected value at position 36: got %v want 178920", seq.Result)
	}
}

func TestGetA001034AtPositionInvalid(t *testing.T) {
	if _, err := GetA001034AtPosition(big.NewInt(0)); err == nil {
		t.Fatal("expected error for position 0")
	}
}
