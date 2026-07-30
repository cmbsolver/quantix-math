package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetA000049Sequence(t *testing.T) {
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
			wantName:   "Number of positive integers <= 2^n of the form 3*x^2 + 4*y^2 (A000049)",
			wantSeq: []*big.Int{
				big.NewInt(0), big.NewInt(0), big.NewInt(2), big.NewInt(3), big.NewInt(5), big.NewInt(9),
				big.NewInt(16), big.NewInt(29), big.NewInt(53), big.NewInt(98), big.NewInt(181),
			},
			wantResult: big.NewInt(181),
			wantErr:    false,
		},
		{
			name:       "OEIS example a(4)=5",
			maxNumber:  big.NewInt(4),
			positional: true,
			wantName:   "Number of positive integers <= 2^n of the form 3*x^2 + 4*y^2 (A000049)",
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
			got, err := GetA000049Sequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetA000049Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got.Name != tt.wantName {
				t.Errorf("GetA000049Sequence() Name = %v, want %v", got.Name, tt.wantName)
			}

			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetA000049Sequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}

			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetA000049Sequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}

func TestGetSequenceA000049Dispatch(t *testing.T) {
	got, err := GetSequence("4", "form_3x2_4y2_a000049", true)
	if err != nil {
		t.Fatalf("GetSequence() error = %v", err)
	}

	if got.Result.String() != "5" {
		t.Fatalf("GetSequence() result = %v, want %v", got.Result, "5")
	}
}
