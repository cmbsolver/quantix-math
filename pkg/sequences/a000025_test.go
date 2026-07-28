package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetA000025Sequence(t *testing.T) {
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
			want:        []int64{1, 1, -2, 3, -3, 3, -5, 7, -6, 6, -10},
			wantErr:     false,
		},
		{
			name:        "Positional term a(6)",
			maxNumber:   big.NewInt(6),
			isPositional: true,
			want:        []int64{-5},
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
			got, err := GetA000025Sequence(tt.maxNumber, tt.isPositional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetA000025Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				var gotSeq []int64
				for _, val := range got.Sequence {
					gotSeq = append(gotSeq, val.Int64())
				}
				if !reflect.DeepEqual(gotSeq, tt.want) {
					t.Errorf("GetA000025Sequence() got = %v, want %v", gotSeq, tt.want)
				}
			}
		})
	}
}
