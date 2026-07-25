package sequences

import (
	"math/big"
	"testing"
)

func TestCalculateGroupsOrderN(t *testing.T) {
	tests := []struct {
		n    int64
		want int64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 1},
		{6, 2},
		{7, 1},
		{8, 5},
		{9, 2},
		{10, 2},
		{11, 1},
		{12, 5},
		{13, 1},
		{14, 2},
		{15, 1},
		{16, 14},
		{32, 51},
		{81, 15},
	}

	for _, tt := range tests {
		got := CalculateGroupsOrderN(tt.n)
		if got != tt.want {
			t.Errorf("CalculateGroupsOrderN(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

func TestGetGroupsOrderNSequence(t *testing.T) {
	maxNumber := big.NewInt(10)
	seq, err := GetGroupsOrderNSequence(maxNumber, false)
	if err != nil {
		t.Fatalf("GetGroupsOrderNSequence failed: %v", err)
	}

	want := []int64{0, 1, 1, 1, 2, 1, 2, 1, 5, 2, 2}
	if len(seq.Sequence) != len(want) {
		t.Fatalf("Got sequence length %d; want %d", len(seq.Sequence), len(want))
	}

	for i, val := range seq.Sequence {
		if val.Int64() != want[i] {
			t.Errorf("Sequence[%d] = %d; want %d", i, val.Int64(), want[i])
		}
	}
}

func TestGetGroupsOrderNAtPosition(t *testing.T) {
	n := big.NewInt(16)
	seq, err := GetGroupsOrderNSequence(n, true)
	if err != nil {
		t.Fatalf("GetGroupsOrderNSequence failed: %v", err)
	}

	want := int64(14)
	if seq.Result.Int64() != want {
		t.Errorf("GetGroupsOrderNSequence(16, true) = %d; want %d", seq.Result.Int64(), want)
	}
}
