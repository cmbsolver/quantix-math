package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetDivisorCountSequence(t *testing.T) {
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
			name:       "First 10 terms of A000005",
			maxNumber:  big.NewInt(10),
			positional: false,
			wantName:   "Number of Divisors (A000005)",
			wantSeq: []*big.Int{
				big.NewInt(1), // d(1) = 1
				big.NewInt(2), // d(2) = 1, 2
				big.NewInt(2), // d(3) = 1, 3
				big.NewInt(3), // d(4) = 1, 2, 4
				big.NewInt(2), // d(5) = 1, 5
				big.NewInt(4), // d(6) = 1, 2, 3, 6
				big.NewInt(2), // d(7) = 1, 7
				big.NewInt(4), // d(8) = 1, 2, 4, 8
				big.NewInt(3), // d(9) = 1, 3, 9
				big.NewInt(4), // d(10) = 1, 2, 5, 10
			},
			wantResult: big.NewInt(4),
			wantErr:    false,
		},
		{
			name:       "Term at position 12",
			maxNumber:  big.NewInt(12),
			positional: true,
			wantName:   "Number of Divisors (A000005)",
			wantSeq:    []*big.Int{big.NewInt(6)}, // d(12) = 1, 2, 3, 4, 6, 12
			wantResult: big.NewInt(6),
			wantErr:    false,
		},
		{
			name:       "Term at position 1",
			maxNumber:  big.NewInt(1),
			positional: true,
			wantName:   "Number of Divisors (A000005)",
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
			got, err := GetDivisorCountSequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDivisorCountSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetDivisorCountSequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetDivisorCountSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetDivisorCountSequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
