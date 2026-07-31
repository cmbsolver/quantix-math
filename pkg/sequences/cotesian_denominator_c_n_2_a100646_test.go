package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateCotesianDenominatorCN2A100646Sequence(t *testing.T) {
	seq, err := GenerateCotesianDenominatorCN2A100646Sequence(big.NewInt(14))
	if err != nil {
		t.Fatalf("GenerateCotesianDenominatorCN2A100646Sequence returned error: %v", err)
	}

	want := []int64{6, 8, 15, 144, 280, 640, 14175, 2240, 199584, 87091200, 875875, 22353408000, 5003856000, 229605376}

	if len(seq.Sequence) != len(want) {
		t.Fatalf("sequence length = %d, want %d", len(seq.Sequence), len(want))
	}

	for i, expected := range want {
		if seq.Sequence[i].Cmp(big.NewInt(expected)) != 0 {
			t.Fatalf("term %d = %s, want %d", i, seq.Sequence[i].String(), expected)
		}
	}
}

func TestGetCotesianDenominatorCN2A100646AtPosition(t *testing.T) {
	seq, err := GetCotesianDenominatorCN2A100646AtPosition(big.NewInt(8))
	if err != nil {
		t.Fatalf("GetCotesianDenominatorCN2A100646AtPosition returned error: %v", err)
	}

	if len(seq.Sequence) != 1 {
		t.Fatalf("sequence length = %d, want 1", len(seq.Sequence))
	}

	if seq.Sequence[0].Cmp(big.NewInt(199584)) != 0 {
		t.Fatalf("term at position 8 = %s, want 199584", seq.Sequence[0].String())
	}
}

func TestCheckExistenceCotesianDenominatorCN2A100646(t *testing.T) {
	exists, index, err := checkExistence(big.NewInt(199584), "cotesian_denominator_c_n_2_a100646")
	if err != nil {
		t.Fatalf("checkExistence returned error: %v", err)
	}

	if !exists {
		t.Fatalf("checkExistence did not find 199584 in A100646")
	}

	if index != "8" {
		t.Fatalf("index = %s, want 8", index)
	}
}
