package sequences

import (
	"math/big"
	"testing"
)

func TestOEISLookups(t *testing.T) {
	tests := []struct {
		id         string
		maxNumber  string
		positional bool
		wantFirst  string
		wantLast   string
		wantLen    int
	}{
		{"A000011", "5", false, "1", "4", 5},
		{"A000013", "10", false, "1", "30", 10},
		{"A000014", "1", true, "0", "0", 1},
		{"A000014", "2", true, "1", "1", 1},
		{"A000020", "5", false, "2", "6", 5},
	}

	for _, tt := range tests {
		t.Run(tt.id+"_"+tt.maxNumber, func(t *testing.T) {
			maxNum := new(big.Int)
			maxNum.SetString(tt.maxNumber, 10)
			got, err := GetOEISLookupSequence(tt.id, "Test", maxNum, tt.positional)
			if err != nil {
				t.Errorf("GetOEISLookupSequence() error = %v", err)
				return
			}
			if len(got.Sequence) != tt.wantLen {
				t.Errorf("got length %v, want %v", len(got.Sequence), tt.wantLen)
			}
			if got.Sequence[0].String() != tt.wantFirst {
				t.Errorf("got first %v, want %v", got.Sequence[0], tt.wantFirst)
			}
			if got.Sequence[len(got.Sequence)-1].String() != tt.wantLast {
				t.Errorf("got last %v, want %v", got.Sequence[len(got.Sequence)-1], tt.wantLast)
			}
		})
	}
}
