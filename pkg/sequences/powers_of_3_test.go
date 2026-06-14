package sequences

import (
	"math/big"
	"testing"
)

func TestGetPowersOf3Sequence(t *testing.T) {
	tests := []struct {
		name         string
		maxNumber    *big.Int
		isPositional bool
		wantResult   *big.Int
		wantLen      int
	}{
		{
			name:         "a(0)",
			maxNumber:    big.NewInt(0),
			isPositional: true,
			wantResult:   big.NewInt(1), // 3^0
			wantLen:      1,
		},
		{
			name:         "a(1)",
			maxNumber:    big.NewInt(1),
			isPositional: true,
			wantResult:   big.NewInt(3), // 3^1
			wantLen:      1,
		},
		{
			name:         "a(5)",
			maxNumber:    big.NewInt(5),
			isPositional: true,
			wantResult:   big.NewInt(243), // 3^5
			wantLen:      1,
		},
		{
			name:         "sequence up to 81",
			maxNumber:    big.NewInt(81),
			isPositional: false,
			wantResult:   big.NewInt(81),
			wantLen:      5, // 1, 3, 9, 27, 81
		},
		{
			name:         "sequence up to 100",
			maxNumber:    big.NewInt(100),
			isPositional: false,
			wantResult:   big.NewInt(81),
			wantLen:      5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPowersOf3Sequence(tt.maxNumber, tt.isPositional)
			if err != nil {
				t.Errorf("GetPowersOf3Sequence() error = %v", err)
				return
			}
			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetPowersOf3Sequence() gotResult = %v, want %v", got.Result, tt.wantResult)
			}
			if len(got.Sequence) != tt.wantLen {
				t.Errorf("GetPowersOf3Sequence() gotLen = %v, want %v", len(got.Sequence), tt.wantLen)
			}
		})
	}
}
