package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateA000038Sequence_OEISData(t *testing.T) {
	got, err := GenerateA000038Sequence(big.NewInt(10))
	if err != nil {
		t.Fatalf("GenerateA000038Sequence returned error: %v", err)
	}

	want := []int64{2, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if len(got.Sequence) != len(want) {
		t.Fatalf("sequence length = %d, want %d", len(got.Sequence), len(want))
	}

	for i, expected := range want {
		if got.Sequence[i].Cmp(big.NewInt(expected)) != 0 {
			t.Fatalf("term %d = %s, want %d", i, got.Sequence[i].String(), expected)
		}
	}
}

func TestGetA000038AtPosition_OEISData(t *testing.T) {
	gotZero, err := GetA000038AtPosition(big.NewInt(0))
	if err != nil {
		t.Fatalf("GetA000038AtPosition returned error: %v", err)
	}
	if gotZero.Result == nil || gotZero.Result.Cmp(big.NewInt(2)) != 0 {
		t.Fatalf("a(0) = %v, want 2", gotZero.Result)
	}

	gotOne, err := GetA000038AtPosition(big.NewInt(1))
	if err != nil {
		t.Fatalf("GetA000038AtPosition returned error: %v", err)
	}
	if gotOne.Result == nil || gotOne.Result.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("a(1) = %v, want 0", gotOne.Result)
	}
}
