package sequences

import (
	"math/big"
	"testing"
)

func TestGetEulerTotientA000010Sequence(t *testing.T) {
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
			wantName:   "Euler totient function phi(n) (A000010)",
			wantSeq:    []int64{1, 1, 2, 2, 4, 2, 6, 4, 6, 4},
			wantResult: 4,
			wantErr:    false,
		},
		{
			maxNumber:  "5",
			positional: true,
			wantName:   "Euler totient function phi(n) (A000010)",
			wantSeq:    []int64{4},
			wantResult: 4,
			wantErr:    false,
		},
		{
			maxNumber:  "0",
			positional: false,
			wantName:   "Euler totient function phi(n) (A000010)",
			wantSeq:    []int64{},
			wantResult: 0,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.maxNumber, func(t *testing.T) {
			maxNum := new(big.Int)
			maxNum.SetString(tt.maxNumber, 10)
			got, err := GetEulerTotientA000010Sequence(maxNum, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEulerTotientA000010Sequence() error = %v, wantErr %v", err, tt.wantErr)
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
				if tt.wantResult != 0 && got.Result.Int64() != tt.wantResult {
					t.Errorf("got result %v, want %v", got.Result, tt.wantResult)
				}
			}
		})
	}
}
