package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetModularJCoefficientsSequence(t *testing.T) {
	tests := []struct {
		name         string
		n            *big.Int
		isPositional bool
		wantName     string
		wantSequence []*big.Int
		wantErr      bool
	}{
		{
			name:         "first 5 terms (n=3)",
			n:            big.NewInt(3),
			isPositional: false,
			wantName:     "Modular function j (A000521)",
			wantSequence: []*big.Int{
				big.NewInt(1),         // a(-1)
				big.NewInt(744),       // a(0)
				big.NewInt(196884),    // a(1)
				big.NewInt(21493760),  // a(2)
				big.NewInt(864299970), // a(3)
			},
			wantErr: false,
		},
		{
			name:         "positional n=-1",
			n:            big.NewInt(-1),
			isPositional: true,
			wantName:     "Modular function j (A000521)",
			wantSequence: []*big.Int{big.NewInt(1)},
			wantErr:      false,
		},
		{
			name:         "positional n=0",
			n:            big.NewInt(0),
			isPositional: true,
			wantName:     "Modular function j (A000521)",
			wantSequence: []*big.Int{big.NewInt(744)},
			wantErr:      false,
		},
		{
			name:         "positional n=1",
			n:            big.NewInt(1),
			isPositional: true,
			wantName:     "Modular function j (A000521)",
			wantSequence: []*big.Int{big.NewInt(196884)},
			wantErr:      false,
		},
		{
			name:         "invalid position n=-2",
			n:            big.NewInt(-2),
			isPositional: true,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetModularJCoefficientsSequence(tt.n, tt.isPositional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetModularJCoefficientsSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetModularJCoefficientsSequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSequence) {
				for i, v := range got.Sequence {
					if i < len(tt.wantSequence) {
						if v.Cmp(tt.wantSequence[i]) != 0 {
							t.Errorf("At index %d: got %v, want %v", i, v, tt.wantSequence[i])
						}
					} else {
						t.Errorf("Extra element at index %d: %v", i, v)
					}
				}
				if len(got.Sequence) < len(tt.wantSequence) {
					t.Errorf("Missing elements starting from index %d", len(got.Sequence))
				}
			}
		})
	}
}
