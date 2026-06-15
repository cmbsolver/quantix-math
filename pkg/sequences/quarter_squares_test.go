package sequences

import (
	"math/big"
	"testing"
)

func TestGetQuarterSquaresSequence(t *testing.T) {
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
			name:       "First 11 terms of A002620 (0 to 10)",
			maxNumber:  big.NewInt(10),
			positional: false,
			wantName:   "Quarter-squares (A002620)",
			wantSeq: []*big.Int{
				big.NewInt(0),  // a(0) = floor(0/4) = 0
				big.NewInt(0),  // a(1) = floor(1/4) = 0
				big.NewInt(1),  // a(2) = floor(4/4) = 1
				big.NewInt(2),  // a(3) = floor(9/4) = 2
				big.NewInt(4),  // a(4) = floor(16/4) = 4
				big.NewInt(6),  // a(5) = floor(25/4) = 6
				big.NewInt(9),  // a(6) = floor(36/4) = 9
				big.NewInt(12), // a(7) = floor(49/4) = 12
				big.NewInt(16), // a(8) = floor(64/4) = 16
				big.NewInt(20), // a(9) = floor(81/4) = 20
				big.NewInt(25), // a(10) = floor(100/4) = 25
			},
			wantResult: big.NewInt(25),
			wantErr:    false,
		},
		{
			name:       "Term at position 5",
			maxNumber:  big.NewInt(5),
			positional: true,
			wantName:   "Quarter-squares (A002620)",
			wantSeq:    []*big.Int{big.NewInt(6)},
			wantResult: big.NewInt(6),
			wantErr:    false,
		},
		{
			name:       "Term at position 6",
			maxNumber:  big.NewInt(6),
			positional: true,
			wantName:   "Quarter-squares (A002620)",
			wantSeq:    []*big.Int{big.NewInt(9)},
			wantResult: big.NewInt(9),
			wantErr:    false,
		},
		{
			name:       "Term at position 0",
			maxNumber:  big.NewInt(0),
			positional: true,
			wantName:   "Quarter-squares (A002620)",
			wantSeq:    []*big.Int{big.NewInt(0)},
			wantResult: big.NewInt(0),
			wantErr:    false,
		},
		{
			name:       "Invalid negative position",
			maxNumber:  big.NewInt(-1),
			positional: true,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetQuarterSquaresSequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQuarterSquaresSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetQuarterSquaresSequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if len(got.Sequence) != len(tt.wantSeq) {
				t.Errorf("GetQuarterSquaresSequence() Sequence length = %v, want %v", len(got.Sequence), len(tt.wantSeq))
			} else {
				for i := range got.Sequence {
					if got.Sequence[i].Cmp(tt.wantSeq[i]) != 0 {
						t.Errorf("GetQuarterSquaresSequence() Sequence[%d] = %v, want %v", i, got.Sequence[i], tt.wantSeq[i])
					}
				}
			}
			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetQuarterSquaresSequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
