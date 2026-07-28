package sequences

import (
	"math/big"
	"testing"
)

func TestCalculateA000023(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{0, "1"},
		{1, "-1"},
		{2, "2"},
		{3, "-2"},
		{4, "8"},
		{5, "8"},
		{6, "112"},
		{7, "656"},
		{8, "5504"},
		{9, "49024"},
		{10, "491264"},
		{22, "152116956851941670912"},
	}

	for _, tt := range tests {
		got := CalculateA000023(tt.n)
		if got.String() != tt.want {
			t.Errorf("CalculateA000023(%d) = %s; want %s", tt.n, got.String(), tt.want)
		}
	}
}

func TestGetA000023Sequence(t *testing.T) {
	maxNumber := big.NewInt(5)
	seq, err := GetA000023Sequence(maxNumber, false)
	if err != nil {
		t.Fatalf("GetA000023Sequence failed: %v", err)
	}

	want := []string{"1", "-1", "2", "-2", "8"}
	if len(seq.Sequence) != len(want) {
		t.Fatalf("got sequence length %d; want %d", len(seq.Sequence), len(want))
	}

	for i, v := range want {
		if seq.Sequence[i].String() != v {
			t.Errorf("seq.Sequence[%d] = %s; want %s", i, seq.Sequence[i].String(), v)
		}
	}
}
