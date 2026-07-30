package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateA000039Sequence_OEISData(t *testing.T) {
	got, err := GenerateA000039Sequence(big.NewInt(10))
	if err != nil {
		t.Fatalf("GenerateA000039Sequence returned error: %v", err)
	}

	want := []int64{1, -2, -3, -5, -6, -10, -11, -17, -21, -27}
	if len(got.Sequence) != len(want) {
		t.Fatalf("sequence length = %d, want %d", len(got.Sequence), len(want))
	}

	for i, expected := range want {
		if got.Sequence[i].Cmp(big.NewInt(expected)) != 0 {
			t.Fatalf("term %d = %s, want %d", i, got.Sequence[i].String(), expected)
		}
	}
}

func TestGetA000039AtPosition_OEISData(t *testing.T) {
	got, err := GetA000039AtPosition(big.NewInt(9))
	if err != nil {
		t.Fatalf("GetA000039AtPosition returned error: %v", err)
	}

	if len(got.Sequence) != 1 {
		t.Fatalf("sequence length = %d, want 1", len(got.Sequence))
	}

	if got.Sequence[0].Cmp(big.NewInt(-27)) != 0 {
		t.Fatalf("a(9) = %s, want -27", got.Sequence[0].String())
	}
}
