package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetA001045Sequence(t *testing.T) {
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
			name:       "non-positional returns OEIS example prefix",
			maxNumber:  big.NewInt(10),
			positional: false,
			wantName:   "Jacobsthal numbers (A001045)",
			wantSeq: []*big.Int{
				big.NewInt(0), big.NewInt(1), big.NewInt(1), big.NewInt(3), big.NewInt(5), big.NewInt(11),
				big.NewInt(21), big.NewInt(43), big.NewInt(85), big.NewInt(171), big.NewInt(341),
			},
			wantResult: big.NewInt(341),
		},
		{
			name:       "positional returns single term",
			maxNumber:  big.NewInt(7),
			positional: true,
			wantName:   "Jacobsthal numbers (A001045)",
			wantSeq:    []*big.Int{big.NewInt(43)},
			wantResult: big.NewInt(43),
		},
		{
			name:       "negative input errors",
			maxNumber:  big.NewInt(-1),
			positional: false,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetA001045Sequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetA001045Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got.Name != tt.wantName {
				t.Errorf("GetA001045Sequence() Name = %v, want %v", got.Name, tt.wantName)
			}

			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetA001045Sequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}

			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetA001045Sequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
