package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateCotesianNumeratorCN2A100645Sequence(t *testing.T) {
	seq, err := GenerateCotesianNumeratorCN2A100645Sequence(big.NewInt(14))
	if err != nil {
		t.Fatalf("GenerateCotesianNumeratorCN2A100645Sequence returned error: %v", err)
	}

	want := []int64{1, 3, 2, 25, 9, 49, -464, 27, -16175, -3237113, -105387, -1737125143, -770720657, -25881785}

	if len(seq.Sequence) != len(want) {
		t.Fatalf("sequence length = %d, want %d", len(seq.Sequence), len(want))
	}

	for i, expected := range want {
		if seq.Sequence[i].Cmp(big.NewInt(expected)) != 0 {
			t.Fatalf("term %d = %s, want %d", i, seq.Sequence[i].String(), expected)
		}
	}
}

func TestGetCotesianNumeratorCN2A100645AtPosition(t *testing.T) {
	seq, err := GetCotesianNumeratorCN2A100645AtPosition(big.NewInt(8))
	if err != nil {
		t.Fatalf("GetCotesianNumeratorCN2A100645AtPosition returned error: %v", err)
	}

	if len(seq.Sequence) != 1 {
		t.Fatalf("sequence length = %d, want 1", len(seq.Sequence))
	}

	if seq.Sequence[0].Cmp(big.NewInt(-16175)) != 0 {
		t.Fatalf("term at position 8 = %s, want -16175", seq.Sequence[0].String())
	}
}

func TestCheckExistenceCotesianNumeratorCN2A100645(t *testing.T) {
	exists, index, err := checkExistence(big.NewInt(-16175), "cotesian_numerator_c_n_2_a100645")
	if err != nil {
		t.Fatalf("checkExistence returned error: %v", err)
	}

	if !exists {
		t.Fatalf("checkExistence did not find -16175 in A100645")
	}

	if index != "8" {
		t.Fatalf("index = %s, want 8", index)
	}
}
