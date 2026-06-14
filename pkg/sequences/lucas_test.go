package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGenerateLucas(t *testing.T) {
	tests := []struct {
		name         string
		maxNumber    *big.Int
		isPositional bool
		wantSeq      []*big.Int
		wantResult   *big.Int
		wantErr      bool
	}{
		{
			name:         "Lucas sequence up to 10",
			maxNumber:    big.NewInt(10),
			isPositional: false,
			wantSeq: []*big.Int{
				big.NewInt(2),
				big.NewInt(1),
				big.NewInt(3),
				big.NewInt(4),
				big.NewInt(7),
			},
		},
		{
			name:         "Lucas number at position 0",
			maxNumber:    big.NewInt(0),
			isPositional: true,
			wantSeq:      []*big.Int{big.NewInt(2)},
			wantResult:   big.NewInt(2),
		},
		{
			name:         "Lucas number at position 1",
			maxNumber:    big.NewInt(1),
			isPositional: true,
			wantSeq:      []*big.Int{big.NewInt(1)},
			wantResult:   big.NewInt(1),
		},
		{
			name:         "Lucas number at position 4",
			maxNumber:    big.NewInt(4),
			isPositional: true,
			wantSeq:      []*big.Int{big.NewInt(7)},
			wantResult:   big.NewInt(7),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateLucas(tt.maxNumber, tt.isPositional)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GenerateLucas() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GenerateLucas() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
			if tt.isPositional && got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GenerateLucas() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
