package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetSylvesterSequence(t *testing.T) {
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
			name:       "First 5 terms of A000058 (max 2000)",
			maxNumber:  big.NewInt(2000),
			positional: false,
			wantName:   "Sylvester's sequence (A000058)",
			wantSeq: []*big.Int{
				big.NewInt(2),
				big.NewInt(3),
				big.NewInt(7),
				big.NewInt(43),
				big.NewInt(1807),
			},
			wantResult: big.NewInt(1807),
			wantErr:    false,
		},
		{
			name:       "Term at position 0",
			maxNumber:  big.NewInt(0),
			positional: true,
			wantName:   "Sylvester's sequence (A000058)",
			wantSeq:    []*big.Int{big.NewInt(2)},
			wantResult: big.NewInt(2),
			wantErr:    false,
		},
		{
			name:       "Term at position 4",
			maxNumber:  big.NewInt(4),
			positional: true,
			wantName:   "Sylvester's sequence (A000058)",
			wantSeq:    []*big.Int{big.NewInt(1807)},
			wantResult: big.NewInt(1807),
			wantErr:    false,
		},
		{
			name:       "Invalid position -1",
			maxNumber:  big.NewInt(-1),
			positional: true,
			wantErr:    true,
		},
		{
			name:       "Invalid max number 1",
			maxNumber:  big.NewInt(1),
			positional: false,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSylvesterSequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSylvesterSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetSylvesterSequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetSylvesterSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetSylvesterSequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
