package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetCubesA000578Sequence(t *testing.T) {
	tests := []struct {
		name         string
		maxNumber    *big.Int
		isPositional bool
		wantName     string
		wantSequence []*big.Int
		wantResult   *big.Int
		wantErr      bool
	}{
		{
			name:         "first 10 terms",
			maxNumber:    big.NewInt(1000),
			isPositional: false,
			wantName:     "Cubes (A000578)",
			wantSequence: []*big.Int{
				big.NewInt(0), big.NewInt(1), big.NewInt(8), big.NewInt(27),
				big.NewInt(64), big.NewInt(125), big.NewInt(216), big.NewInt(343),
				big.NewInt(512), big.NewInt(729), big.NewInt(1000),
			},
			wantResult: big.NewInt(1000),
			wantErr:    false,
		},
		{
			name:         "positional 3rd term",
			maxNumber:    big.NewInt(3),
			isPositional: true,
			wantName:     "Cubes (A000578)",
			wantSequence: []*big.Int{big.NewInt(27)},
			wantResult:   big.NewInt(27),
			wantErr:      false,
		},
		{
			name:         "positional 10th term",
			maxNumber:    big.NewInt(10),
			isPositional: true,
			wantName:     "Cubes (A000578)",
			wantSequence: []*big.Int{big.NewInt(1000)},
			wantResult:   big.NewInt(1000),
			wantErr:      false,
		},
		{
			name:         "negative max number sequence",
			maxNumber:    big.NewInt(-1),
			isPositional: false,
			wantName:     "Cubes (A000578)",
			wantSequence: []*big.Int{},
			wantResult:   nil,
			wantErr:      false,
		},
		{
			name:         "negative position error",
			maxNumber:    big.NewInt(-1),
			isPositional: true,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCubesA000578Sequence(tt.maxNumber, tt.isPositional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCubesA000578Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetCubesA000578Sequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSequence) {
				t.Errorf("GetCubesA000578Sequence() Sequence = %v, want %v", got.Sequence, tt.wantSequence)
			}
			if (got.Result == nil && tt.wantResult != nil) || (got.Result != nil && tt.wantResult == nil) || (got.Result != nil && tt.wantResult != nil && got.Result.Cmp(tt.wantResult) != 0) {
				t.Errorf("GetCubesA000578Sequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
