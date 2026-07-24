package sequences

import (
	"math/big"
	"testing"
)

func TestGetSqrtPrimeA000006Sequence(t *testing.T) {
	tests := []struct {
		maxNumber  string
		positional bool
		wantName   string
		wantSeq    []int64
		wantResult int64
		wantErr    bool
	}{
		{
			maxNumber:  "10",
			positional: false,
			wantName:   "Integer part of square root of n-th prime (A000006)",
			wantSeq:    []int64{1, 1, 2, 2, 3, 3, 4, 4, 4, 5},
			wantResult: 5,
			wantErr:    false,
		},
		{
			maxNumber:  "5",
			positional: true,
			wantName:   "Integer part of square root of n-th prime (A000006)",
			wantSeq:    []int64{3},
			wantResult: 3,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.maxNumber, func(t *testing.T) {
			maxNum := new(big.Int)
			maxNum.SetString(tt.maxNumber, 10)
			got, err := GetSqrtPrimeA000006Sequence(maxNum, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSqrtPrimeA000006Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.Name != tt.wantName {
					t.Errorf("got name %v, want %v", got.Name, tt.wantName)
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
			}
		})
	}
}
