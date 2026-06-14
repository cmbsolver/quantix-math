package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetFubiniSequence(t *testing.T) {
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
			name:       "First 6 terms of A000670",
			maxNumber:  big.NewInt(5),
			positional: false,
			wantName:   "Fubini numbers (A000670)",
			wantSeq: []*big.Int{
				big.NewInt(1),   // a(0)
				big.NewInt(1),   // a(1)
				big.NewInt(3),   // a(2)
				big.NewInt(13),  // a(3)
				big.NewInt(75),  // a(4)
				big.NewInt(541), // a(5)
			},
			wantResult: big.NewInt(541),
			wantErr:    false,
		},
		{
			name:       "Term at position 6 (a(6)=4683)",
			maxNumber:  big.NewInt(6),
			positional: true,
			wantName:   "Fubini numbers (A000670)",
			wantSeq:    []*big.Int{big.NewInt(4683)},
			wantResult: big.NewInt(4683),
			wantErr:    false,
		},
		{
			name:       "Term at position 10 (a(10)=102247563)",
			maxNumber:  big.NewInt(10),
			positional: true,
			wantName:   "Fubini numbers (A000670)",
			wantSeq:    []*big.Int{big.NewInt(102247563)},
			wantResult: big.NewInt(102247563),
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
			got, err := GetFubiniSequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFubiniSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetFubiniSequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetFubiniSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetFubiniSequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
