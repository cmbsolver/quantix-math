package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetInitialDigitA000030Sequence(t *testing.T) {
	tests := []struct {
		name       string
		maxNumber  *big.Int
		positional bool
		wantSeq    []*big.Int
		wantResult *big.Int
		wantErr    bool
	}{
		{
			name:       "First 11 terms (0 to 10)",
			maxNumber:  big.NewInt(10),
			positional: false,
			wantSeq:    []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(5), big.NewInt(6), big.NewInt(7), big.NewInt(8), big.NewInt(9), big.NewInt(1)},
			wantResult: big.NewInt(1),
			wantErr:    false,
		},
		{
			name:       "Position 0",
			maxNumber:  big.NewInt(0),
			positional: true,
			wantSeq:    []*big.Int{big.NewInt(0)},
			wantResult: big.NewInt(0),
			wantErr:    false,
		},
		{
			name:       "Position 20",
			maxNumber:  big.NewInt(20),
			positional: true,
			wantSeq:    []*big.Int{big.NewInt(2)},
			wantResult: big.NewInt(2),
			wantErr:    false,
		},
		{
			name:       "Position 12345",
			maxNumber:  big.NewInt(12345),
			positional: true,
			wantSeq:    []*big.Int{big.NewInt(1)},
			wantResult: big.NewInt(1),
			wantErr:    false,
		},
		{
			name:       "Negative input",
			maxNumber:  big.NewInt(-1),
			positional: false,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetInitialDigitA000030Sequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInitialDigitA000030Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetInitialDigitA000030Sequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetInitialDigitA000030Sequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
