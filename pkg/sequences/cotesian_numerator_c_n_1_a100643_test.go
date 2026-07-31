package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateCotesianNumeratorCN1A100643Sequence(t *testing.T) {
	seq, err := GenerateCotesianNumeratorCN1A100643Sequence(big.NewInt(14))
	if err != nil {
		t.Fatalf("GenerateCotesianNumeratorCN1A100643Sequence returned error: %v", err)
	}

	want := []int64{1, 2, 3, 16, 25, 9, 3577, 2944, 15741, 26575, 4495513, 12504, 56280729661, 44436679}

	if len(seq.Sequence) != len(want) {
		t.Fatalf("sequence length = %d, want %d", len(seq.Sequence), len(want))
	}

	for i, expected := range want {
		if seq.Sequence[i].Cmp(big.NewInt(expected)) != 0 {
			t.Fatalf("term %d = %s, want %d", i, seq.Sequence[i].String(), expected)
		}
	}
}

func TestGetCotesianNumeratorCN1A100643AtPosition(t *testing.T) {
	seq, err := GetCotesianNumeratorCN1A100643AtPosition(big.NewInt(10))
	if err != nil {
		t.Fatalf("GetCotesianNumeratorCN1A100643AtPosition returned error: %v", err)
	}

	if len(seq.Sequence) != 1 {
		t.Fatalf("sequence length = %d, want 1", len(seq.Sequence))
	}

	if seq.Sequence[0].Cmp(big.NewInt(4495513)) != 0 {
		t.Fatalf("term at position 10 = %s, want 4495513", seq.Sequence[0].String())
	}
}

func TestCheckExistenceCotesianNumeratorCN1A100643(t *testing.T) {
	exists, index, err := checkExistence(big.NewInt(4495513), "cotesian_numerator_c_n_1_a100643")
	if err != nil {
		t.Fatalf("checkExistence returned error: %v", err)
	}

	if !exists {
		t.Fatalf("checkExistence did not find 4495513 in A100643")
	}

	if index != "10" {
		t.Fatalf("index = %s, want 10", index)
	}
}
