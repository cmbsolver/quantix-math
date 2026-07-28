package sequences

import (
	"math/big"
	"testing"
)

func TestGetA000017Sequence(t *testing.T) {
	tests := []struct {
		maxNumber  *big.Int
		positional bool
		wantName   string
		wantResult *big.Int
		wantSeq    []int64
	}{
		{
			maxNumber:  big.NewInt(10),
			positional: false,
			wantName:   "Point symmetric queens (A000017)",
			wantResult: big.NewInt(12),
			wantSeq:    []int64{1, 0, 0, 2, 2, 4, 8, 4, 16, 12},
		},
		{
			maxNumber:  big.NewInt(4),
			positional: true,
			wantName:   "Point symmetric queens (A000017)",
			wantResult: big.NewInt(2),
			wantSeq:    []int64{2},
		},
		{
			maxNumber:  big.NewInt(16),
			positional: true,
			wantName:   "Point symmetric queens (A000017)",
			wantResult: big.NewInt(3000),
			wantSeq:    []int64{3000},
		},
	}

	for _, tt := range tests {
		got, err := GetA000017Sequence(tt.maxNumber, tt.positional)
		if err != nil {
			t.Errorf("GetA000017Sequence(%v, %v) error = %v", tt.maxNumber, tt.positional, err)
			continue
		}
		if got.Name != tt.wantName {
			t.Errorf("got.Name = %v, want %v", got.Name, tt.wantName)
		}
		if got.Result.Cmp(tt.wantResult) != 0 {
			t.Errorf("got.Result = %v, want %v", got.Result, tt.wantResult)
		}
		if len(got.Sequence) != len(tt.wantSeq) {
			t.Errorf("len(got.Sequence) = %v, want %v", len(got.Sequence), len(tt.wantSeq))
		} else {
			for i := range got.Sequence {
				if got.Sequence[i].Int64() != tt.wantSeq[i] {
					t.Errorf("got.Sequence[%d] = %v, want %v", i, got.Sequence[i], tt.wantSeq[i])
				}
			}
		}
	}
}
