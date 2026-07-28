package sequences

import (
	"math/big"
	"testing"
)

func TestGetPrimitivePermutationGroupsA000019Sequence(t *testing.T) {
	tests := []struct {
		n            int64
		isPositional bool
		want         []int64
	}{
		{1, true, []int64{1}},
		{2, true, []int64{1}},
		{3, true, []int64{2}},
		{4, true, []int64{2}},
		{5, true, []int64{5}},
		{16, true, []int64{22}},
		{25, true, []int64{28}},
		{5, false, []int64{1, 1, 2, 2, 5}},
	}

	for _, tt := range tests {
		n := big.NewInt(tt.n)
		res, err := GetPrimitivePermutationGroupsA000019Sequence(n, tt.isPositional)
		if err != nil {
			t.Errorf("GetPrimitivePermutationGroupsA000019Sequence(%d, %v) error: %v", tt.n, tt.isPositional, err)
			continue
		}

		if len(res.Sequence) != len(tt.want) {
			t.Errorf("GetPrimitivePermutationGroupsA000019Sequence(%d, %v) got length %d, want %d", tt.n, tt.isPositional, len(res.Sequence), len(tt.want))
			continue
		}

		for i, v := range tt.want {
			if res.Sequence[i].Int64() != v {
				t.Errorf("GetPrimitivePermutationGroupsA000019Sequence(%d, %v) index %d: got %d, want %d", tt.n, tt.isPositional, i, res.Sequence[i].Int64(), v)
			}
		}
	}
}
