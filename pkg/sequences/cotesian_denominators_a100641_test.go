package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateCotesianDenominatorsA100641Sequence(t *testing.T) {
	seq, err := GenerateCotesianDenominatorsA100641Sequence(big.NewInt(52))
	if err != nil {
		t.Fatalf("GenerateCotesianDenominatorsA100641Sequence returned error: %v", err)
	}

	want := []int64{
		1,
		2, 2,
		6, 3, 6,
		8, 8, 8, 8,
		90, 45, 15, 45, 90,
		288, 96, 144, 144, 96, 288,
		840, 35, 280, 105, 280, 35, 840,
		17280, 17280, 640, 17280, 17280, 640, 17280, 17280,
		28350, 14175, 14175, 14175, 2835, 14175, 14175, 14175, 28350,
		89600, 89600, 2240, 5600, 44800, 44800, 5600,
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

func TestGetCotesianDenominatorsA100641AtPosition(t *testing.T) {
	seq, err := GetCotesianDenominatorsA100641AtPosition(big.NewInt(50))
	if err != nil {
		t.Fatalf("GetCotesianDenominatorsA100641AtPosition returned error: %v", err)
	}

	if len(seq.Sequence) != 1 {
		t.Fatalf("sequence length = %d, want 1", len(seq.Sequence))
	}

	if seq.Sequence[0].Cmp(big.NewInt(44800)) != 0 {
		t.Fatalf("term at position 50 = %s, want 44800", seq.Sequence[0].String())
	}
}

func TestCheckExistenceCotesianDenominatorsA100641(t *testing.T) {
	exists, index, err := checkExistence(big.NewInt(28350), "cotesian_denominators_a100641")
	if err != nil {
		t.Fatalf("checkExistence returned error: %v", err)
	}

	if !exists {
		t.Fatalf("checkExistence did not find 28350 in A100641")
	}

	if index != "36" {
		t.Fatalf("index = %s, want 36", index)
	}
}
