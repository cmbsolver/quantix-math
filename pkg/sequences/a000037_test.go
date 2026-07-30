package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateA000037Sequence_OEISData(t *testing.T) {
	got, err := GenerateA000037Sequence(big.NewInt(10))
	if err != nil {
		t.Fatalf("GenerateA000037Sequence returned error: %v", err)
	}

	want := []int64{2, 3, 5, 6, 7, 8, 10, 11, 12, 13}
	if len(got.Sequence) != len(want) {
		t.Fatalf("sequence length = %d, want %d", len(got.Sequence), len(want))
	}

	for i, expected := range want {
		if got.Sequence[i].Cmp(big.NewInt(expected)) != 0 {
			t.Fatalf("term %d = %s, want %d", i+1, got.Sequence[i].String(), expected)
		}
	}
}

func TestGetA000037AtPosition_OEISData(t *testing.T) {
	got, err := GetA000037AtPosition(big.NewInt(1))
	if err != nil {
		t.Fatalf("GetA000037AtPosition returned error: %v", err)
	}

	if len(got.Sequence) != 1 {
		t.Fatalf("sequence length = %d, want 1", len(got.Sequence))
	}

	if got.Sequence[0].Cmp(big.NewInt(2)) != 0 {
		t.Fatalf("a(1) = %s, want 2", got.Sequence[0].String())
	}
}
