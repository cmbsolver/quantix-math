package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetDivisorCountA000005Sequence(t *testing.T) {
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
			name:       "First 20 terms of A000005",
			maxNumber:  big.NewInt(20),
			positional: false,
			wantName:   "Number of Divisors (A000005)",
			wantSeq: []*big.Int{
				big.NewInt(1), big.NewInt(2), big.NewInt(2), big.NewInt(3), big.NewInt(2),
				big.NewInt(4), big.NewInt(2), big.NewInt(4), big.NewInt(3), big.NewInt(4),
				big.NewInt(2), big.NewInt(6), big.NewInt(2), big.NewInt(4), big.NewInt(4),
				big.NewInt(5), big.NewInt(2), big.NewInt(6), big.NewInt(2), big.NewInt(6),
			},
			wantResult: big.NewInt(6),
			wantErr:    false,
		},
		{
			name:       "12th term of A000005 (Positional)",
			maxNumber:  big.NewInt(12),
			positional: true,
			wantName:   "Number of Divisors (A000005)",
			wantSeq:    []*big.Int{big.NewInt(6)},
			wantResult: big.NewInt(6),
			wantErr:    false,
		},
		{
			name:       "1st term of A000005 (Positional)",
			maxNumber:  big.NewInt(1),
			positional: true,
			wantName:   "Number of Divisors (A000005)",
			wantSeq:    []*big.Int{big.NewInt(1)},
			wantResult: big.NewInt(1),
			wantErr:    false,
		},
		{
			name:       "Invalid max number",
			maxNumber:  big.NewInt(0),
			positional: false,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDivisorCountA000005Sequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDivisorCountA000005Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetDivisorCountA000005Sequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetDivisorCountA000005Sequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetDivisorCountA000005Sequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
