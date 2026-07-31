package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateCotesianDenominatorCN3A100648Sequence(t *testing.T) {
	seq, err := GenerateCotesianDenominatorCN3A100648Sequence(big.NewInt(12))
	if err != nil {
		t.Fatalf("GenerateCotesianDenominatorCN3A100648Sequence returned error: %v", err)
	}

	want := []string{"8", "45", "144", "105", "17280", "14175", "5600", "12474", "1935360", "1576575", "28740096000", "156370500"}

	if len(seq.Sequence) != len(want) {
		t.Fatalf("sequence length = %d, want %d", len(seq.Sequence), len(want))
	}

	for i, expected := range want {
		if seq.Sequence[i].String() != expected {
			t.Fatalf("term %d = %s, want %s", i, seq.Sequence[i].String(), expected)
		}
	}
}

func TestGetCotesianDenominatorCN3A100648AtPosition(t *testing.T) {
	seq, err := GetCotesianDenominatorCN3A100648AtPosition(big.NewInt(8))
	if err != nil {
		t.Fatalf("GetCotesianDenominatorCN3A100648AtPosition returned error: %v", err)
	}

	if len(seq.Sequence) != 1 {
		t.Fatalf("sequence length = %d, want 1", len(seq.Sequence))
	}

	if seq.Sequence[0].Cmp(big.NewInt(1935360)) != 0 {
		t.Fatalf("term at position 8 = %s, want 1935360", seq.Sequence[0].String())
	}
}

func TestCheckExistenceCotesianDenominatorCN3A100648(t *testing.T) {
	exists, index, err := checkExistence(big.NewInt(1935360), "cotesian_denominator_c_n_3_a100648")
	if err != nil {
		t.Fatalf("checkExistence returned error: %v", err)
	}

	if !exists {
		t.Fatalf("checkExistence did not find 1935360 in A100648")
	}

	if index != "8" {
		t.Fatalf("index = %s, want 8", index)
	}
}
