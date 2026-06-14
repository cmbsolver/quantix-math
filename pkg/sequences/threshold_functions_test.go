package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetThresholdFunctionsSequence(t *testing.T) {
	tests := []struct {
		name       string
		maxNumber  *big.Int
		positional bool
		wantName   string
		wantSeq    []*big.Int
		wantResult *big.Int
		wantErr    bool
	}{
		{
			name:       "First 4 terms of A000609",
			maxNumber:  big.NewInt(3),
			positional: false,
			wantName:   "Threshold functions (A000609)",
			wantSeq: []*big.Int{
				big.NewInt(2),   // a(0)
				big.NewInt(4),   // a(1)
				big.NewInt(14),  // a(2)
				big.NewInt(104), // a(3)
			},
			wantResult: big.NewInt(104),
			wantErr:    false,
		},
		{
			name:       "Term at position 4 (a(4)=1882)",
			maxNumber:  big.NewInt(4),
			positional: true,
			wantName:   "Threshold functions (A000609)",
			wantSeq:    []*big.Int{big.NewInt(1882)},
			wantResult: big.NewInt(1882),
			wantErr:    false,
		},
		{
			name:       "Term at position 9 (large value)",
			maxNumber:  big.NewInt(9),
			positional: true,
			wantName:   "Threshold functions (A000609)",
			wantSeq: []*big.Int{func() *big.Int {
				res := new(big.Int)
				res.SetString("144130531453121108", 10)
				return res
			}()},
			wantResult: func() *big.Int {
				res := new(big.Int)
				res.SetString("144130531453121108", 10)
				return res
			}(),
			wantErr: false,
		},
		{
			name:       "Invalid position -1",
			maxNumber:  big.NewInt(-1),
			positional: true,
			wantErr:    true,
		},
		{
			name:       "Out of bounds position 10",
			maxNumber:  big.NewInt(10),
			positional: true,
			wantErr:    true,
		},
		{
			name:       "Max number exceeds table size (should cap)",
			maxNumber:  big.NewInt(20),
			positional: false,
			wantName:   "Threshold functions (A000609)",
			wantResult: func() *big.Int {
				res := new(big.Int)
				res.SetString("144130531453121108", 10)
				return res
			}(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetThresholdFunctionsSequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetThresholdFunctionsSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetThresholdFunctionsSequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) && !tt.positional && tt.maxNumber.Int64() < 10 {
				t.Errorf("GetThresholdFunctionsSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetThresholdFunctionsSequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
