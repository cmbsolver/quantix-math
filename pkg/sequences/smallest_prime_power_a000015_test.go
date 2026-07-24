package sequences

import (
	"math/big"
	"testing"
)

func TestGetSmallestPrimePowerA000015Sequence(t *testing.T) {
	tests := []struct {
		maxNumber  string
		positional bool
		wantSeq    []int64
		wantResult int64
	}{
		{"10", false, []int64{1, 2, 3, 4, 5, 7, 7, 8, 9, 11}, 11},
		{"5", true, []int64{5}, 5},
		{"6", true, []int64{7}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.maxNumber, func(t *testing.T) {
			maxNum := new(big.Int)
			maxNum.SetString(tt.maxNumber, 10)
			got, err := GetSmallestPrimePowerA000015Sequence(maxNum, tt.positional)
			if err != nil {
				t.Errorf("GetSmallestPrimePowerA000015Sequence() error = %v", err)
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
