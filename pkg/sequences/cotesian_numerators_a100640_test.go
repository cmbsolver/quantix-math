package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateCotesianNumeratorsA100640Sequence(t *testing.T) {
	seq, err := GenerateCotesianNumeratorsA100640Sequence(big.NewInt(59))
	if err != nil {
		t.Fatalf("GenerateCotesianNumeratorsA100640Sequence returned error: %v", err)
	}

	want := []int64{
		0,
		1, 1,
		1, 2, 1,
		1, 3, 3, 1,
		7, 16, 2, 16, 7,
		19, 25, 25, 25, 25, 19,
		41, 9, 9, 34, 9, 9, 41,
		751, 3577, 49, 2989, 2989, 49, 3577, 751,
		989, 2944, -464, 5248, -454, 5248, -464, 2944, 989,
		2857, 15741, 27, 1209, 2889, 2889, 1209, 27, 15741, 2857,
		16067, 26575, -16175, 5675,
	}

	if len(seq.Sequence) != len(want) {
		t.Fatalf("sequence length = %d, want %d", len(seq.Sequence), len(want))
	}

	for i, expected := range want {
		if seq.Sequence[i].Cmp(big.NewInt(expected)) != 0 {
			t.Fatalf("term %d = %s, want %d", i, seq.Sequence[i].String(), expected)
		}
	}
}

func TestGetCotesianNumeratorsA100640AtPosition(t *testing.T) {
	seq, err := GetCotesianNumeratorsA100640AtPosition(big.NewInt(57))
	if err != nil {
		t.Fatalf("GetCotesianNumeratorsA100640AtPosition returned error: %v", err)
	}

	if len(seq.Sequence) != 1 {
		t.Fatalf("sequence length = %d, want 1", len(seq.Sequence))
	}

	if seq.Sequence[0].Cmp(big.NewInt(-16175)) != 0 {
		t.Fatalf("term at position 57 = %s, want -16175", seq.Sequence[0].String())
	}
}

func TestCheckExistenceCotesianNumeratorsA100640(t *testing.T) {
	exists, index, err := checkExistence(big.NewInt(-16175), "cotesian_numerators_a100640")
	if err != nil {
		t.Fatalf("checkExistence returned error: %v", err)
	}

	if !exists {
		t.Fatalf("checkExistence did not find -16175 in A100640")
	}

	if index != "57" {
		t.Fatalf("index = %s, want 57", index)
	}
}
