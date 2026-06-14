package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetAlkanesSequence(t *testing.T) {
	tests := []struct {
		name       string
		maxNumber  *big.Int
		positional bool
		wantName   string
		wantSeq    []int64
		wantResult int64
		wantErr    bool
	}{
		{
			name:       "First 10 terms of A000602",
			maxNumber:  big.NewInt(9),
			positional: false,
			wantName:   "Number of Alkanes (A000602)",
			wantSeq:    []int64{1, 1, 1, 1, 2, 3, 5, 9, 18, 35},
			wantResult: 35,
			wantErr:    false,
		},
		{
			name:       "Term at position 10 (75)",
			maxNumber:  big.NewInt(10),
			positional: true,
			wantName:   "Number of Alkanes (A000602)",
			wantSeq:    []int64{75},
			wantResult: 75,
			wantErr:    false,
		},
		{
			name:       "Term at position 12 (355)",
			maxNumber:  big.NewInt(12),
			positional: true,
			wantName:   "Number of Alkanes (A000602)",
			wantSeq:    []int64{355},
			wantResult: 355,
			wantErr:    false,
		},
		{
			name:       "Term at position 20 (366319)",
			maxNumber:  big.NewInt(20),
			positional: true,
			wantName:   "Number of Alkanes (A000602)",
			wantSeq:    []int64{366319},
			wantResult: 366319,
			wantErr:    false,
		},
		{
			name:       "Invalid position -1",
			maxNumber:  big.NewInt(-1),
			positional: true,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAlkanesSequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAlkanesSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetAlkanesSequence() Name = %v, want %v", got.Name, tt.wantName)
			}

			gotSeq := make([]int64, len(got.Sequence))
			for i, v := range got.Sequence {
				gotSeq[i] = v.Int64()
			}

			if !reflect.DeepEqual(gotSeq, tt.wantSeq) {
				t.Errorf("GetAlkanesSequence() Sequence = %v, want %v", gotSeq, tt.wantSeq)
			}
			if got.Result.Int64() != tt.wantResult {
				t.Errorf("GetAlkanesSequence() Result = %v, want %v", got.Result.Int64(), tt.wantResult)
			}
		})
	}
}
