package sequences

import (
	"math/big"
	"testing"
)

func TestCalculateA000262(t *testing.T) {
	tests := []struct {
		n    int64
		want *big.Int
	}{
		{0, big.NewInt(1)},
		{1, big.NewInt(1)},
		{2, big.NewInt(3)},
		{3, big.NewInt(13)},
		{4, big.NewInt(73)},
		{5, big.NewInt(501)},
		{6, big.NewInt(4051)},
		{7, big.NewInt(37633)},
	}

	for _, tt := range tests {
		got := CalculateA000262(tt.n)
		if got.Cmp(tt.want) != 0 {
			t.Errorf("CalculateA000262(%d) = %v; want %v", tt.n, got, tt.want)
		}
	}
}

func TestGenerateSetsOfListsSequence(t *testing.T) {
	maxNumber := big.NewInt(4)
	seq, err := GenerateSetsOfListsSequence(maxNumber)
	if err != nil {
		t.Fatalf("GenerateSetsOfListsSequence failed: %v", err)
	}

	expected := []*big.Int{
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(3),
		big.NewInt(13),
		big.NewInt(73),
	}

	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, v := range seq.Sequence {
		if v.Cmp(expected[i]) != 0 {
			t.Errorf("At index %d: expected %v, got %v", i, expected[i], v)
		}
	}
}
