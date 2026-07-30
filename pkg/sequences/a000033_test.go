package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetA000033Sequence(t *testing.T) {
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
			wantName:   "Ménage hit polynomials (A000033)",
			wantSeq: []*big.Int{
				big.NewInt(0), big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(40), big.NewInt(210),
				big.NewInt(1477), big.NewInt(11672), big.NewInt(104256), big.NewInt(1036050), big.NewInt(11338855),
			},
			wantResult: big.NewInt(11338855),
			wantErr:    false,
		},
		{
			name:       "positional n=10",
			maxNumber:  big.NewInt(10),
			positional: true,
			wantName:   "Ménage hit polynomials (A000033)",
			wantSeq:    []*big.Int{big.NewInt(11338855)},
			wantResult: big.NewInt(11338855),
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
			got, err := GetA000033Sequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetA000033Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got.Name != tt.wantName {
				t.Errorf("GetA000033Sequence() Name = %v, want %v", got.Name, tt.wantName)
			}

			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetA000033Sequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}

			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetA000033Sequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
