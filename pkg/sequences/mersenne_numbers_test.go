package sequences

import (
	"math/big"
	"testing"
)

func TestGetMersenneNumbersSequence(t *testing.T) {
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
			wantResult:   big.NewInt(0), // 2^0 - 1 = 0
			wantLen:      1,
		},
		{
			name:         "a(1)",
			maxNumber:    big.NewInt(1),
			isPositional: true,
			wantResult:   big.NewInt(1), // 2^1 - 1 = 1
			wantLen:      1,
		},
		{
			name:         "a(2)",
			maxNumber:    big.NewInt(2),
			isPositional: true,
			wantResult:   big.NewInt(3), // 2^2 - 1 = 3
			wantLen:      1,
		},
		{
			name:         "a(3)",
			maxNumber:    big.NewInt(3),
			isPositional: true,
			wantResult:   big.NewInt(7), // 2^3 - 1 = 7
			wantLen:      1,
		},
		{
			name:         "a(10)",
			maxNumber:    big.NewInt(10),
			isPositional: true,
			wantResult:   big.NewInt(1023), // 2^10 - 1 = 1023
			wantLen:      1,
		},
		{
			name:         "sequence up to 15",
			maxNumber:    big.NewInt(15),
			isPositional: false,
			wantResult:   big.NewInt(15),
			wantLen:      5, // 0, 1, 3, 7, 15
		},
		{
			name:         "sequence up to 30",
			maxNumber:    big.NewInt(30),
			isPositional: false,
			wantResult:   big.NewInt(15),
			wantLen:      5, // 0, 1, 3, 7, 15
		},
		{
			name:         "sequence up to 31",
			maxNumber:    big.NewInt(31),
			isPositional: false,
			wantResult:   big.NewInt(31),
			wantLen:      6, // 0, 1, 3, 7, 15, 31
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMersenneNumbersSequence(tt.maxNumber, tt.isPositional)
			if err != nil {
				t.Errorf("GetMersenneNumbersSequence() error = %v", err)
				return
			}
			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetMersenneNumbersSequence() gotResult = %v, want %v", got.Result, tt.wantResult)
			}
			if len(got.Sequence) != tt.wantLen {
				t.Errorf("GetMersenneNumbersSequence() gotLen = %v, want %v", len(got.Sequence), tt.wantLen)
			}
		})
	}
}
