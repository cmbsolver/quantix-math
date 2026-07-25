package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetA000008Sequence(t *testing.T) {
	tests := []struct {
		maxNumber  *big.Int
		positional bool
		wantName   string
		wantSeq    []*big.Int
		wantResult *big.Int
		wantErr    bool
	}{
		{
			maxNumber:  big.NewInt(0),
			positional: false,
			wantName:   "Ways to Make Change (A000008)",
			wantSeq:    []*big.Int{big.NewInt(1)},
			wantResult: big.NewInt(1),
			wantErr:    false,
		},
		{
			maxNumber:  big.NewInt(5),
			positional: false,
			wantName:   "Ways to Make Change (A000008)",
			wantSeq:    []*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(2), big.NewInt(2), big.NewInt(3), big.NewInt(4)},
			wantResult: big.NewInt(4),
			wantErr:    false,
		},
		{
			maxNumber:  big.NewInt(10),
			positional: true,
			wantName:   "Ways to Make Change (A000008)",
			wantSeq:    []*big.Int{big.NewInt(11)},
			wantResult: big.NewInt(11),
			wantErr:    false,
		},
		{
			maxNumber:  big.NewInt(20),
			positional: true,
			wantName:   "Ways to Make Change (A000008)",
			wantSeq:    []*big.Int{big.NewInt(40)},
			wantResult: big.NewInt(40),
			wantErr:    false,
		},
		{
			maxNumber:  big.NewInt(60),
			positional: true,
			wantName:   "Ways to Make Change (A000008)",
			wantSeq:    []*big.Int{big.NewInt(546)},
			wantResult: big.NewInt(546),
			wantErr:    false,
		},
		{
			maxNumber:  big.NewInt(-1),
			positional: false,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.maxNumber.String(), func(t *testing.T) {
			got, err := GetA000008Sequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetA000008Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetA000008Sequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetA000008Sequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetA000008Sequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
