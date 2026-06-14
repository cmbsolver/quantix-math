package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetSumOddDivisorsSequence(t *testing.T) {
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
			name:       "First 10 terms of A000593",
			maxNumber:  big.NewInt(10),
			positional: false,
			wantName:   "Sum of Odd Divisors (A000593)",
			wantSeq:    []int64{1, 1, 4, 1, 6, 4, 8, 1, 13, 6},
			wantResult: 6,
			wantErr:    false,
		},
		{
			name:       "Term at position 1",
			maxNumber:  big.NewInt(1),
			positional: true,
			wantName:   "Sum of Odd Divisors (A000593)",
			wantSeq:    []int64{1},
			wantResult: 1,
			wantErr:    false,
		},
		{
			name:       "Term at position 3",
			maxNumber:  big.NewInt(3),
			positional: true,
			wantName:   "Sum of Odd Divisors (A000593)",
			wantSeq:    []int64{4},
			wantResult: 4,
			wantErr:    false,
		},
		{
			name:       "Term at position 9",
			maxNumber:  big.NewInt(9),
			positional: true,
			wantName:   "Sum of Odd Divisors (A000593)",
			wantSeq:    []int64{13},
			wantResult: 13,
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
			got, err := GetSumOddDivisorsSequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSumOddDivisorsSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetSumOddDivisorsSequence() Name = %v, want %v", got.Name, tt.wantName)
			}

			gotSeq := make([]int64, len(got.Sequence))
			for i, v := range got.Sequence {
				gotSeq[i] = v.Int64()
			}

			if !reflect.DeepEqual(gotSeq, tt.wantSeq) {
				t.Errorf("GetSumOddDivisorsSequence() Sequence = %v, want %v", gotSeq, tt.wantSeq)
			}
			if got.Result.Int64() != tt.wantResult {
				t.Errorf("GetSumOddDivisorsSequence() Result = %v, want %v", got.Result.Int64(), tt.wantResult)
			}
		})
	}
}
