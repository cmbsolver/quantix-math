package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetA000047Sequence(t *testing.T) {
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
			name:       "OEIS sample terms through n=10",
			maxNumber:  big.NewInt(10),
			positional: false,
			wantName:   "Number of integers <= 2^n of form x^2 - 2y^2 (A000047)",
			wantSeq: []*big.Int{
				big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(5), big.NewInt(8), big.NewInt(15),
				big.NewInt(26), big.NewInt(48), big.NewInt(87), big.NewInt(161), big.NewInt(299),
			},
			wantResult: big.NewInt(299),
			wantErr:    false,
		},
		{
			name:       "OEIS example a(3)=5",
			maxNumber:  big.NewInt(3),
			positional: true,
			wantName:   "Number of integers <= 2^n of form x^2 - 2y^2 (A000047)",
			wantSeq:    []*big.Int{big.NewInt(5)},
			wantResult: big.NewInt(5),
			wantErr:    false,
		},
		{
			name:       "negative input",
			maxNumber:  big.NewInt(-1),
			positional: false,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetA000047Sequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetA000047Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got.Name != tt.wantName {
				t.Errorf("GetA000047Sequence() Name = %v, want %v", got.Name, tt.wantName)
			}

			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetA000047Sequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}

			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetA000047Sequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}

func TestGetSequenceA000047Dispatch(t *testing.T) {
	got, err := GetSequence("3", "form_x2_minus_2y2_a000047", true)
	if err != nil {
		t.Fatalf("GetSequence() error = %v", err)
	}

	if got.Result.String() != "5" {
		t.Fatalf("GetSequence() result = %v, want %v", got.Result, "5")
	}
}
