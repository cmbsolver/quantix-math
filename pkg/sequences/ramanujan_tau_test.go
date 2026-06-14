package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetRamanujanTauSequence(t *testing.T) {
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
			name:       "First 10 terms of A000594",
			maxNumber:  big.NewInt(10),
			positional: false,
			wantName:   "Ramanujan's tau function (A000594)",
			wantSeq:    []int64{1, -24, 252, -1472, 4830, -6048, -16744, 84480, -113643, -115920},
			wantResult: -115920,
			wantErr:    false,
		},
		{
			name:       "Term at position 1",
			maxNumber:  big.NewInt(1),
			positional: true,
			wantName:   "Ramanujan's tau function (A000594)",
			wantSeq:    []int64{1},
			wantResult: 1,
			wantErr:    false,
		},
		{
			name:       "Term at position 2",
			maxNumber:  big.NewInt(2),
			positional: true,
			wantName:   "Ramanujan's tau function (A000594)",
			wantSeq:    []int64{-24},
			wantResult: -24,
			wantErr:    false,
		},
		{
			name:       "Term at position 5",
			maxNumber:  big.NewInt(5),
			positional: true,
			wantName:   "Ramanujan's tau function (A000594)",
			wantSeq:    []int64{4830},
			wantResult: 4830,
			wantErr:    false,
		},
		{
			name:       "Invalid position 0",
			maxNumber:  big.NewInt(0),
			positional: true,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRamanujanTauSequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRamanujanTauSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetRamanujanTauSequence() Name = %v, want %v", got.Name, tt.wantName)
			}

			gotSeq := make([]int64, len(got.Sequence))
			for i, v := range got.Sequence {
				gotSeq[i] = v.Int64()
			}

			if !reflect.DeepEqual(gotSeq, tt.wantSeq) {
				t.Errorf("GetRamanujanTauSequence() Sequence = %v, want %v", gotSeq, tt.wantSeq)
			}
			if got.Result.Int64() != tt.wantResult {
				t.Errorf("GetRamanujanTauSequence() Result = %v, want %v", got.Result.Int64(), tt.wantResult)
			}
		})
	}
}
