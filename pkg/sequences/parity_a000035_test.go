package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateParityA000035Sequence_OEISExample(t *testing.T) {
	got, err := GenerateParityA000035Sequence(big.NewInt(10))
	if err != nil {
		t.Fatalf("GenerateParityA000035Sequence returned error: %v", err)
	}

	want := []int64{0, 1, 0, 1, 0, 1, 0, 1, 0, 1}
	if len(got.Sequence) != len(want) {
		t.Fatalf("sequence length = %d, want %d", len(got.Sequence), len(want))
	}

	for i, expected := range want {
		if got.Sequence[i].Cmp(big.NewInt(expected)) != 0 {
			t.Fatalf("term %d = %s, want %d", i, got.Sequence[i].String(), expected)
		}
	}
}

func TestGetParityA000035AtPosition_OEISExample(t *testing.T) {
	got, err := GetParityA000035AtPosition(big.NewInt(11))
	if err != nil {
		t.Fatalf("GetParityA000035AtPosition returned error: %v", err)
	}

	if len(got.Sequence) != 1 {
		t.Fatalf("sequence length = %d, want 1", len(got.Sequence))
	}

	if got.Sequence[0].Cmp(big.NewInt(1)) != 0 {
		t.Fatalf("a(11) = %s, want 1", got.Sequence[0].String())
	}
}
