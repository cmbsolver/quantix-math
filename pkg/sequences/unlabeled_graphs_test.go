package sequences

import (
	"math/big"
	"testing"
)

func TestCalculateA000088(t *testing.T) {
	tests := []struct {
		n    int64
		want *big.Int
	}{
		{0, big.NewInt(1)},
		{1, big.NewInt(1)},
		{2, big.NewInt(2)},
		{3, big.NewInt(4)},
		{4, big.NewInt(11)},
		{5, big.NewInt(34)},
		{6, big.NewInt(156)},
		{7, big.NewInt(1044)},
		{8, big.NewInt(12346)},
		{9, big.NewInt(274668)},
	}

	for _, tt := range tests {
		got, err := calculateA000088(tt.n)
		if err != nil {
			t.Errorf("calculateA000088(%d) error: %v", tt.n, err)
			continue
		}
		if got.Cmp(tt.want) != 0 {
			t.Errorf("calculateA000088(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
