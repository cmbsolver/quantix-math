package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetPowersOf2Sequence(t *testing.T) {
	tests := []struct {
		maxNumber  string
		positional bool
		want       []string
		wantRes    string
		wantErr    bool
	}{
		{
			maxNumber:  "5",
			positional: false,
			want:       []string{"1", "2", "4", "8", "16", "32"},
			wantRes:    "32",
			wantErr:    false,
		},
		{
			maxNumber:  "0",
			positional: false,
			want:       []string{"1"},
			wantRes:    "1",
			wantErr:    false,
		},
		{
			maxNumber:  "10",
			positional: true,
			want:       []string{"1024"},
			wantRes:    "1024",
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.maxNumber, func(t *testing.T) {
			n := new(big.Int)
			n.SetString(tt.maxNumber, 10)
			got, err := GetPowersOf2Sequence(n, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPowersOf2Sequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				var gotStrings []string
				for _, v := range got.Sequence {
					gotStrings = append(gotStrings, v.String())
				}
				if !reflect.DeepEqual(gotStrings, tt.want) {
					t.Errorf("GetPowersOf2Sequence() sequence = %v, want %v", gotStrings, tt.want)
				}
				if got.Result.String() != tt.wantRes {
					t.Errorf("GetPowersOf2Sequence() result = %v, want %v", got.Result.String(), tt.wantRes)
				}
			}
		})
	}
}
