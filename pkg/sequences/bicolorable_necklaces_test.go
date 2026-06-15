package sequences

import (
	"math/big"
	"testing"
)

func TestCalculateA000048(t *testing.T) {
	tests := []struct {
		n    int64
		want *big.Int
	}{
		{0, big.NewInt(1)},
		{1, big.NewInt(1)},
		{2, big.NewInt(1)},
		{3, big.NewInt(1)},
		{4, big.NewInt(2)},
		{5, big.NewInt(3)},
		{6, big.NewInt(5)},
		{7, big.NewInt(9)},
		{8, big.NewInt(16)},
		{9, big.NewInt(28)},
		{10, big.NewInt(51)},
		{11, big.NewInt(93)},
		{12, big.NewInt(170)},
	}

	for _, tt := range tests {
		got := calculateA000048(tt.n)
		if got.Cmp(tt.want) != 0 {
			t.Errorf("calculateA000048(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

func TestGetBicolorableNecklacesSequence(t *testing.T) {
	maxNumber := big.NewInt(5)
	seq, err := GetBicolorableNecklacesSequence(maxNumber, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := []*big.Int{
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(3),
	}

	if len(seq.Sequence) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, val := range seq.Sequence {
		if val.Cmp(expected[i]) != 0 {
			t.Errorf("At index %d: expected %v, got %v", i, expected[i], val)
		}
	}
}
