package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetKolakoskiSequence(t *testing.T) {
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
			name:       "First 10 terms of A000002",
			maxNumber:  big.NewInt(10),
			positional: false,
			wantName:   "Kolakoski Sequence (A000002)",
			wantSeq: []*big.Int{
				big.NewInt(1),
				big.NewInt(2),
				big.NewInt(2),
				big.NewInt(1),
				big.NewInt(1),
				big.NewInt(2),
				big.NewInt(1),
				big.NewInt(2),
				big.NewInt(2),
				big.NewInt(1),
			},
			wantResult: big.NewInt(1),
			wantErr:    false,
		},
		{
			name:       "Term at position 7",
			maxNumber:  big.NewInt(7),
			positional: true,
			wantName:   "Kolakoski Sequence (A000002)",
			wantSeq:    []*big.Int{big.NewInt(1)},
			wantResult: big.NewInt(1),
			wantErr:    false,
		},
		{
			name:       "Term at position 1",
			maxNumber:  big.NewInt(1),
			positional: true,
			wantName:   "Kolakoski Sequence (A000002)",
			wantSeq:    []*big.Int{big.NewInt(1)},
			wantResult: big.NewInt(1),
			wantErr:    false,
		},
		{
			name:       "Invalid position 0",
			maxNumber:  big.NewInt(0),
			positional: true,
			wantErr:    true,
		},
		{
			name:       "Invalid max number 0",
			maxNumber:  big.NewInt(0),
			positional: false,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetKolakoskiSequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKolakoskiSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetKolakoskiSequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetKolakoskiSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetKolakoskiSequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
