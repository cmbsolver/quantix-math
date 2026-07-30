package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetA000031Sequence(t *testing.T) {
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
			wantName:   "Necklaces with no turnover allowed (A000031)",
			wantSeq: []*big.Int{
				big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(6), big.NewInt(8),
				big.NewInt(14), big.NewInt(20), big.NewInt(36), big.NewInt(60), big.NewInt(108),
			},
			wantResult: big.NewInt(108),
			wantErr:    false,
		},
		{
			name:       "positional n=7",
			maxNumber:  big.NewInt(7),
			positional: true,
			wantName:   "Necklaces with no turnover allowed (A000031)",
			wantSeq:    []*big.Int{big.NewInt(20)},
			wantResult: big.NewInt(20),
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
			got, err := GetA000031Sequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetA000031Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got.Name != tt.wantName {
				t.Errorf("GetA000031Sequence() Name = %v, want %v", got.Name, tt.wantName)
			}

			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetA000031Sequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}

			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetA000031Sequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
