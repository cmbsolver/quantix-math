package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateCotesianDenominatorCN1A100644Sequence(t *testing.T) {
	seq, err := GenerateCotesianDenominatorCN1A100644Sequence(big.NewInt(14))
	if err != nil {
		t.Fatalf("GenerateCotesianDenominatorCN1A100644Sequence returned error: %v", err)
	}

	want := []int64{2, 3, 8, 45, 96, 35, 17280, 14175, 89600, 149688, 29030400, 79625, 402361344000, 312741000}

	if len(seq.Sequence) != len(want) {
		t.Fatalf("sequence length = %d, want %d", len(seq.Sequence), len(want))
	}

	for i, expected := range want {
		if seq.Sequence[i].Cmp(big.NewInt(expected)) != 0 {
			t.Fatalf("term %d = %s, want %d", i, seq.Sequence[i].String(), expected)
		}
	}
}

func TestGetCotesianDenominatorCN1A100644AtPosition(t *testing.T) {
	seq, err := GetCotesianDenominatorCN1A100644AtPosition(big.NewInt(10))
	if err != nil {
		t.Fatalf("GetCotesianDenominatorCN1A100644AtPosition returned error: %v", err)
	}

	if len(seq.Sequence) != 1 {
		t.Fatalf("sequence length = %d, want 1", len(seq.Sequence))
	}

	if seq.Sequence[0].Cmp(big.NewInt(29030400)) != 0 {
		t.Fatalf("term at position 10 = %s, want 29030400", seq.Sequence[0].String())
	}
}

func TestCheckExistenceCotesianDenominatorCN1A100644(t *testing.T) {
	exists, index, err := checkExistence(big.NewInt(29030400), "cotesian_denominator_c_n_1_a100644")
	if err != nil {
		t.Fatalf("checkExistence returned error: %v", err)
	}

	if !exists {
		t.Fatalf("checkExistence did not find 29030400 in A100644")
	}

	if index != "10" {
		t.Fatalf("index = %s, want 10", index)
	}
}
