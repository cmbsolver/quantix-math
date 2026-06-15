package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetSubfactorialSequence(t *testing.T) {
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
			name:       "First 10 terms of A000166",
			maxNumber:  big.NewInt(9),
			positional: false,
			wantName:   "Subfactorial (A000166)",
			wantSeq: []*big.Int{
				big.NewInt(1),      // a(0)
				big.NewInt(0),      // a(1)
				big.NewInt(1),      // a(2)
				big.NewInt(2),      // a(3)
				big.NewInt(9),      // a(4)
				big.NewInt(44),     // a(5)
				big.NewInt(265),    // a(6)
				big.NewInt(1854),   // a(7)
				big.NewInt(14833),  // a(8)
				big.NewInt(133496), // a(9)
			},
			wantResult: big.NewInt(133496),
			wantErr:    false,
		},
		{
			name:       "Term at position 4",
			maxNumber:  big.NewInt(4),
			positional: true,
			wantName:   "Subfactorial (A000166)",
			wantSeq:    []*big.Int{big.NewInt(9)},
			wantResult: big.NewInt(9),
			wantErr:    false,
		},
		{
			name:       "Term at position 0",
			maxNumber:  big.NewInt(0),
			positional: true,
			wantName:   "Subfactorial (A000166)",
			wantSeq:    []*big.Int{big.NewInt(1)},
			wantResult: big.NewInt(1),
			wantErr:    false,
		},
		{
			name:       "Term at position 1",
			maxNumber:  big.NewInt(1),
			positional: true,
			wantName:   "Subfactorial (A000166)",
			wantSeq:    []*big.Int{big.NewInt(0)},
			wantResult: big.NewInt(0),
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
			got, err := GetSubfactorialSequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubfactorialSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetSubfactorialSequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetSubfactorialSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetSubfactorialSequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
