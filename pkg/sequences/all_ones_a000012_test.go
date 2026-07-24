package sequences

import (
	"math/big"
	"testing"
)

func TestGetAllOnesA000012Sequence(t *testing.T) {
	tests := []struct {
		maxNumber  string
		positional bool
		wantSeq    []int64
		wantResult int64
	}{
		{"5", false, []int64{1, 1, 1, 1, 1}, 1},
		{"10", true, []int64{1}, 1},
		{"0", false, []int64{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.maxNumber, func(t *testing.T) {
			maxNum := new(big.Int)
			maxNum.SetString(tt.maxNumber, 10)
			got, err := GetAllOnesA000012Sequence(maxNum, tt.positional)
			if err != nil {
				t.Errorf("GetAllOnesA000012Sequence() error = %v", err)
				return
			}
			if len(got.Sequence) != len(tt.wantSeq) {
				t.Errorf("got sequence length %v, want %v", len(got.Sequence), len(tt.wantSeq))
			}
			for i, v := range tt.wantSeq {
				if got.Sequence[i].Int64() != v {
					t.Errorf("got sequence[%d] = %v, want %v", i, got.Sequence[i], v)
				}
			}
			if got.Result.Int64() != tt.wantResult {
				t.Errorf("got result %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
