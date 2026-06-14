package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetSquaresSequence(t *testing.T) {
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
			wantName:   "Squares (A000290)",
			wantSeq:    []int64{0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100},
			wantResult: 100,
			wantErr:    false,
		},
		{
			name:       "Terms up to 50",
			maxNumber:  big.NewInt(50),
			positional: false,
			wantName:   "Squares (A000290)",
			wantSeq:    []int64{0, 1, 4, 9, 16, 25, 36, 49},
			wantResult: 49,
			wantErr:    false,
		},
		{
			name:       "Term at position 0",
			maxNumber:  big.NewInt(0),
			positional: true,
			wantName:   "Squares (A000290)",
			wantSeq:    []int64{0},
			wantResult: 0,
			wantErr:    false,
		},
		{
			name:       "Term at position 5",
			maxNumber:  big.NewInt(5),
			positional: true,
			wantName:   "Squares (A000290)",
			wantSeq:    []int64{25},
			wantResult: 25,
			wantErr:    false,
		},
		{
			name:       "Term at position 10",
			maxNumber:  big.NewInt(10),
			positional: true,
			wantName:   "Squares (A000290)",
			wantSeq:    []int64{100},
			wantResult: 100,
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
			got, err := GetSquaresSequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSquaresSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetSquaresSequence() Name = %v, want %v", got.Name, tt.wantName)
			}

			gotSeq := make([]int64, len(got.Sequence))
			for i, v := range got.Sequence {
				gotSeq[i] = v.Int64()
			}

			if !reflect.DeepEqual(gotSeq, tt.wantSeq) {
				t.Errorf("GetSquaresSequence() Sequence = %v, want %v", gotSeq, tt.wantSeq)
			}
			if got.Result.Int64() != tt.wantResult {
				t.Errorf("GetSquaresSequence() Result = %v, want %v", got.Result.Int64(), tt.wantResult)
			}
		})
	}
}
