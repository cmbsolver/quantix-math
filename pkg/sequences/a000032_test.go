package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetA000032Sequence(t *testing.T) {
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
			name:       "OEIS sample terms through n=10",
			maxNumber:  big.NewInt(10),
			positional: false,
			wantName:   "Lucas numbers (A000032)",
			wantSeq: []*big.Int{
				big.NewInt(2), big.NewInt(1), big.NewInt(3), big.NewInt(4), big.NewInt(7), big.NewInt(11),
				big.NewInt(18), big.NewInt(29), big.NewInt(47), big.NewInt(76), big.NewInt(123),
			},
			wantResult: big.NewInt(123),
			wantErr:    false,
		},
		{
			name:       "positional n=8",
			maxNumber:  big.NewInt(8),
			positional: true,
			wantName:   "Lucas numbers (A000032)",
			wantSeq:    []*big.Int{big.NewInt(47)},
			wantResult: big.NewInt(47),
			wantErr:    false,
		},
		{
			name:       "negative input",
			maxNumber:  big.NewInt(-1),
			positional: false,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetA000032Sequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetA000032Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got.Name != tt.wantName {
				t.Errorf("GetA000032Sequence() Name = %v, want %v", got.Name, tt.wantName)
			}

			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetA000032Sequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}

			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetA000032Sequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
