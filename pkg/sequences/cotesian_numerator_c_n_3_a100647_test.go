package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateCotesianNumeratorCN3A100647Sequence(t *testing.T) {
	seq, err := GenerateCotesianNumeratorCN3A100647Sequence(big.NewInt(12))
	if err != nil {
		t.Fatalf("GenerateCotesianNumeratorCN3A100647Sequence returned error: %v", err)
	}

	want := []string{"1", "16", "25", "34", "2989", "5248", "1209", "5675", "560593", "893128", "11148172711", "109420087"}

	if len(seq.Sequence) != len(want) {
		t.Fatalf("sequence length = %d, want %d", len(seq.Sequence), len(want))
	}

	for i, expected := range want {
		if seq.Sequence[i].String() != expected {
			t.Fatalf("term %d = %s, want %s", i, seq.Sequence[i].String(), expected)
		}
	}
}

func TestGetCotesianNumeratorCN3A100647AtPosition(t *testing.T) {
	seq, err := GetCotesianNumeratorCN3A100647AtPosition(big.NewInt(8))
	if err != nil {
		t.Fatalf("GetCotesianNumeratorCN3A100647AtPosition returned error: %v", err)
	}

	if len(seq.Sequence) != 1 {
		t.Fatalf("sequence length = %d, want 1", len(seq.Sequence))
	}

	if seq.Sequence[0].Cmp(big.NewInt(560593)) != 0 {
		t.Fatalf("term at position 8 = %s, want 560593", seq.Sequence[0].String())
	}
}

func TestCheckExistenceCotesianNumeratorCN3A100647(t *testing.T) {
	exists, index, err := checkExistence(big.NewInt(560593), "cotesian_numerator_c_n_3_a100647")
	if err != nil {
		t.Fatalf("checkExistence returned error: %v", err)
	}

	if !exists {
		t.Fatalf("checkExistence did not find 560593 in A100647")
	}

	if index != "8" {
		t.Fatalf("index = %s, want 8", index)
	}
}
