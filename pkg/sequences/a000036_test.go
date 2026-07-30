package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateA000036Sequence_OEISExample(t *testing.T) {
	got, err := GenerateA000036Sequence(big.NewInt(10))
	if err != nil {
		t.Fatalf("GenerateA000036Sequence returned error: %v", err)
	}

	want := []int64{2, 3, 5, 6, 6, -6, 7, 8, 10, 13}
	if len(got.Sequence) != len(want) {
		t.Fatalf("sequence length = %d, want %d", len(got.Sequence), len(want))
	}

	for i, expected := range want {
		if got.Sequence[i].Cmp(big.NewInt(expected)) != 0 {
			t.Fatalf("term %d = %s, want %d", i+1, got.Sequence[i].String(), expected)
		}
	}
}

func TestGetA000036AtPosition_OEISExample(t *testing.T) {
	got, err := GetA000036AtPosition(big.NewInt(14))
	if err != nil {
		t.Fatalf("GetA000036AtPosition returned error: %v", err)
	}

	if len(got.Sequence) != 1 {
		t.Fatalf("sequence length = %d, want 1", len(got.Sequence))
	}

	if got.Sequence[0].Cmp(big.NewInt(-17)) != 0 {
		t.Fatalf("a(14) = %s, want -17", got.Sequence[0].String())
	}
}
