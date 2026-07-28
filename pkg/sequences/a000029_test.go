package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetA000029Sequence(t *testing.T) {
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
			name:       "n=6",
			maxNumber:  big.NewInt(6),
			positional: false,
			wantName:   "Bracelets with n beads of 2 colors (A000029)",
			wantSeq: []*big.Int{
				big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4),
				big.NewInt(6), big.NewInt(8), big.NewInt(13),
			},
			wantResult: big.NewInt(13),
			wantErr:    false,
		},
		{
			name:       "n=7 positional",
			maxNumber:  big.NewInt(7),
			positional: true,
			wantName:   "Bracelets with n beads of 2 colors (A000029)",
			wantSeq:    []*big.Int{big.NewInt(18)},
			wantResult: big.NewInt(18),
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetA000029Sequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetA000029Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetA000029Sequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetA000029Sequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetA000029Sequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
