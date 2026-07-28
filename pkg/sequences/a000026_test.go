package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetA000026Sequence(t *testing.T) {
	tests := []struct {
		name        string
		maxNumber   *big.Int
		isPositional bool
		want        []int64
		wantErr     bool
	}{
		{
			name:        "First few terms",
			maxNumber:   big.NewInt(10),
			isPositional: false,
			want:        []int64{1, 2, 3, 4, 5, 6, 7, 6, 6, 10},
			wantErr:     false,
		},
		{
			name:        "Positional term a(24)",
			maxNumber:   big.NewInt(24),
			isPositional: true,
			want:        []int64{18},
			wantErr:     false,
		},
		{
			name:        "Positional term a(1)",
			maxNumber:   big.NewInt(1),
			isPositional: true,
			want:        []int64{1},
			wantErr:     false,
		},
		{
			name:        "Negative input",
			maxNumber:   big.NewInt(-1),
			isPositional: false,
			want:        nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetA000026Sequence(tt.maxNumber, tt.isPositional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetA000026Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				var gotSeq []int64
				for _, val := range got.Sequence {
					gotSeq = append(gotSeq, val.Int64())
				}
				if !reflect.DeepEqual(gotSeq, tt.want) {
					t.Errorf("GetA000026Sequence() got = %v, want %v", gotSeq, tt.want)
				}
			}
		})
	}
}
