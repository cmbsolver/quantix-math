package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetFourthPowersSequence(t *testing.T) {
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
			name:       "Terms up to 100",
			maxNumber:  big.NewInt(100),
			positional: false,
			wantName:   "Fourth Powers (A000583)",
			wantSeq:    []int64{0, 1, 16, 81},
			wantResult: 81,
			wantErr:    false,
		},
		{
			name:       "Term at position 0",
			maxNumber:  big.NewInt(0),
			positional: true,
			wantName:   "Fourth Powers (A000583)",
			wantSeq:    []int64{0},
			wantResult: 0,
			wantErr:    false,
		},
		{
			name:       "Term at position 1",
			maxNumber:  big.NewInt(1),
			positional: true,
			wantName:   "Fourth Powers (A000583)",
			wantSeq:    []int64{1},
			wantResult: 1,
			wantErr:    false,
		},
		{
			name:       "Term at position 2",
			maxNumber:  big.NewInt(2),
			positional: true,
			wantName:   "Fourth Powers (A000583)",
			wantSeq:    []int64{16},
			wantResult: 16,
			wantErr:    false,
		},
		{
			name:       "Term at position 3",
			maxNumber:  big.NewInt(3),
			positional: true,
			wantName:   "Fourth Powers (A000583)",
			wantSeq:    []int64{81},
			wantResult: 81,
			wantErr:    false,
		},
		{
			name:       "Term at position 10",
			maxNumber:  big.NewInt(10),
			positional: true,
			wantName:   "Fourth Powers (A000583)",
			wantSeq:    []int64{10000},
			wantResult: 10000,
			wantErr:    false,
		},
		{
			name:       "Negative position",
			maxNumber:  big.NewInt(-1),
			positional: true,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFourthPowersSequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFourthPowersSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetFourthPowersSequence() Name = %v, want %v", got.Name, tt.wantName)
			}

			gotSeq := make([]int64, len(got.Sequence))
			for i, v := range got.Sequence {
				gotSeq[i] = v.Int64()
			}

			if !reflect.DeepEqual(gotSeq, tt.wantSeq) {
				t.Errorf("GetFourthPowersSequence() Sequence = %v, want %v", gotSeq, tt.wantSeq)
			}
			if got.Result.Int64() != tt.wantResult {
				t.Errorf("GetFourthPowersSequence() Result = %v, want %v", got.Result.Int64(), tt.wantResult)
			}
		})
	}
}
