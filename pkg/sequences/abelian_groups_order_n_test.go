package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetAbelianGroupsOrderNSequence(t *testing.T) {
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
			name:       "First 10 terms of A000688",
			maxNumber:  big.NewInt(10),
			positional: false,
			wantName:   "Number of Abelian groups (A000688)",
			wantSeq: []*big.Int{
				big.NewInt(1), // a(1)
				big.NewInt(1), // a(2) = P(1)
				big.NewInt(1), // a(3) = P(1)
				big.NewInt(2), // a(4) = P(2)
				big.NewInt(1), // a(5) = P(1)
				big.NewInt(1), // a(6) = P(1)*P(1)
				big.NewInt(1), // a(7) = P(1)
				big.NewInt(3), // a(8) = P(3)
				big.NewInt(2), // a(9) = P(2)
				big.NewInt(1), // a(10) = P(1)*P(1)
			},
			wantResult: big.NewInt(1),
			wantErr:    false,
		},
		{
			name:       "Term at position 12 (a(12)=2)",
			maxNumber:  big.NewInt(12),
			positional: true,
			wantName:   "Number of Abelian groups (A000688)",
			wantSeq:    []*big.Int{big.NewInt(2)}, // a(12)=a(2^2*3)=P(2)*P(1)=2*1=2
			wantResult: big.NewInt(2),
			wantErr:    false,
		},
		{
			name:       "Term at position 72 (a(72)=6)",
			maxNumber:  big.NewInt(72),
			positional: true,
			wantName:   "Number of Abelian groups (A000688)",
			wantSeq:    []*big.Int{big.NewInt(6)}, // a(72)=a(2^3*3^2)=P(3)*P(2)=3*2=6
			wantResult: big.NewInt(6),
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
			got, err := GetAbelianGroupsOrderNSequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAbelianGroupsOrderNSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetAbelianGroupsOrderNSequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetAbelianGroupsOrderNSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetAbelianGroupsOrderNSequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
