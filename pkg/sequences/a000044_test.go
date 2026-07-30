package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetA000044Sequence(t *testing.T) {
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
			name:       "OEIS sample terms through n=15",
			maxNumber:  big.NewInt(15),
			positional: false,
			wantName:   "Dying rabbits (A000044)",
			wantSeq: []*big.Int{
				big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(5),
				big.NewInt(8), big.NewInt(13), big.NewInt(21), big.NewInt(34), big.NewInt(55), big.NewInt(89),
				big.NewInt(144), big.NewInt(232), big.NewInt(375), big.NewInt(606),
			},
			wantResult: big.NewInt(606),
			wantErr:    false,
		},
		{
			name:       "positional n=14",
			maxNumber:  big.NewInt(14),
			positional: true,
			wantName:   "Dying rabbits (A000044)",
			wantSeq:    []*big.Int{big.NewInt(375)},
			wantResult: big.NewInt(375),
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
			got, err := GetA000044Sequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetA000044Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got.Name != tt.wantName {
				t.Errorf("GetA000044Sequence() Name = %v, want %v", got.Name, tt.wantName)
			}

			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetA000044Sequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}

			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetA000044Sequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
