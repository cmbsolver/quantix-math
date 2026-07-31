package sequences

import (
	"math/big"
	"testing"
)

func TestCalculateA010025(t *testing.T) {
	tests := []struct {
		n    int
		want uint64
	}{
		{0, 1},
		{1, 13},
		{2, 85},
		{3, 377},
		{4, 1239},
		{5, 3291},
		{6, 7503},
		{7, 15275},
		{8, 28517},
		{9, 49729},
	}

	for _, tt := range tests {
		if got := CalculateA010025(tt.n); got != tt.want {
			t.Errorf("CalculateA010025(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestGetA010025Sequence(t *testing.T) {
	maxNumber := big.NewInt(6)
	seq, err := GetA010025Sequence(maxNumber, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := []int64{1, 13, 85, 377, 1239, 3291}
	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, v := range expected {
		if seq.Sequence[i].Int64() != v {
			t.Errorf("Sequence[%d] = %d, want %d", i, seq.Sequence[i].Int64(), v)
		}
	}

	if seq.Result.Int64() != 3291 {
		t.Errorf("Result = %d, want 3291", seq.Result.Int64())
	}
}

func TestGetA010025AtPosition(t *testing.T) {
	n := big.NewInt(4)
	seq, err := GetA010025Sequence(n, true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if seq.Result.Int64() != 1239 {
		t.Errorf("Result at position 4 = %d, want 1239", seq.Result.Int64())
	}
}

func TestCheckExistenceA010025(t *testing.T) {
	exists, index, err := checkExistence(big.NewInt(1239), "crystal_ball_squashed_d5_lattice_a010025")
	if err != nil {
		t.Fatalf("checkExistence returned error: %v", err)
	}
	if !exists {
		t.Fatalf("checkExistence did not find 1239 in A010025")
	}
	if index != "4" {
		t.Fatalf("checkExistence index = %s, want 4", index)
	}
}
