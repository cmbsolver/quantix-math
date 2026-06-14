package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetGroupsOrderNSequence(t *testing.T) {
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
			name:       "First 10 terms of A000001",
			maxNumber:  big.NewInt(9),
			positional: false,
			wantName:   "Number of Groups (A000001)",
			wantSeq: []*big.Int{
				big.NewInt(0), // a(0)
				big.NewInt(1), // a(1)
				big.NewInt(1), // a(2)
				big.NewInt(1), // a(3)
				big.NewInt(2), // a(4)
				big.NewInt(1), // a(5)
				big.NewInt(2), // a(6)
				big.NewInt(1), // a(7)
				big.NewInt(5), // a(8)
				big.NewInt(2), // a(9)
			},
			wantResult: big.NewInt(2),
			wantErr:    false,
		},
		{
			name:       "Term at position 16 (a(16)=14)",
			maxNumber:  big.NewInt(16),
			positional: true,
			wantName:   "Number of Groups (A000001)",
			wantSeq:    []*big.Int{big.NewInt(14)},
			wantResult: big.NewInt(14),
			wantErr:    false,
		},
		{
			name:       "Term at position 64 (a(64)=267)",
			maxNumber:  big.NewInt(64),
			positional: true,
			wantName:   "Number of Groups (A000001)",
			wantSeq:    []*big.Int{big.NewInt(267)},
			wantResult: big.NewInt(267),
			wantErr:    false,
		},
		{
			name:       "Invalid position 101 (out of range)",
			maxNumber:  big.NewInt(101),
			positional: true,
			wantErr:    true,
		},
		{
			name:       "Invalid max number -1",
			maxNumber:  big.NewInt(-1),
			positional: false,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGroupsOrderNSequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupsOrderNSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetGroupsOrderNSequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetGroupsOrderNSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetGroupsOrderNSequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
