package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetPowersOf4Sequence(t *testing.T) {
	tests := []struct {
		name         string
		maxNumber    *big.Int
		isPositional bool
		want         []*big.Int
		wantResult   *big.Int
		wantErr      bool
	}{
		{
			name:         "Sequence up to 100",
			maxNumber:    big.NewInt(100),
			isPositional: false,
			want: []*big.Int{
				big.NewInt(1),
				big.NewInt(4),
				big.NewInt(16),
				big.NewInt(64),
			},
			wantErr: false,
		},
		{
			name:         "Position 0",
			maxNumber:    big.NewInt(0),
			isPositional: true,
			want:         []*big.Int{big.NewInt(1)},
			wantResult:   big.NewInt(1),
			wantErr:      false,
		},
		{
			name:         "Position 3",
			maxNumber:    big.NewInt(3),
			isPositional: true,
			want:         []*big.Int{big.NewInt(64)},
			wantResult:   big.NewInt(64),
			wantErr:      false,
		},
		{
			name:         "Negative position",
			maxNumber:    big.NewInt(-1),
			isPositional: true,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPowersOf4Sequence(tt.maxNumber, tt.isPositional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPowersOf4Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(got.Sequence, tt.want) {
					t.Errorf("GetPowersOf4Sequence() Sequence = %v, want %v", got.Sequence, tt.want)
				}
				if tt.isPositional && got.Result.Cmp(tt.wantResult) != 0 {
					t.Errorf("GetPowersOf4Sequence() Result = %v, want %v", got.Result, tt.wantResult)
				}
			}
		})
	}
}
