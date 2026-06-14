package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetZeroCharacteristicSequence(t *testing.T) {
	tests := []struct {
		maxNumber  *big.Int
		positional bool
		wantName   string
		wantSeq    []*big.Int
		wantResult *big.Int
		wantErr    bool
	}{
		{
			maxNumber:  big.NewInt(0),
			positional: false,
			wantName:   "Zero Characteristic (A000007)",
			wantSeq:    []*big.Int{big.NewInt(1)},
			wantResult: big.NewInt(1),
			wantErr:    false,
		},
		{
			maxNumber:  big.NewInt(5),
			positional: false,
			wantName:   "Zero Characteristic (A000007)",
			wantSeq:    []*big.Int{big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)},
			wantResult: big.NewInt(0),
			wantErr:    false,
		},
		{
			maxNumber:  big.NewInt(0),
			positional: true,
			wantName:   "Zero Characteristic (A000007)",
			wantSeq:    []*big.Int{big.NewInt(1)},
			wantResult: big.NewInt(1),
			wantErr:    false,
		},
		{
			maxNumber:  big.NewInt(1),
			positional: true,
			wantName:   "Zero Characteristic (A000007)",
			wantSeq:    []*big.Int{big.NewInt(0)},
			wantResult: big.NewInt(0),
			wantErr:    false,
		},
		{
			maxNumber:  big.NewInt(10),
			positional: true,
			wantName:   "Zero Characteristic (A000007)",
			wantSeq:    []*big.Int{big.NewInt(0)},
			wantResult: big.NewInt(0),
			wantErr:    false,
		},
		{
			maxNumber:  big.NewInt(-1),
			positional: false,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(reflect.TypeOf(tt.maxNumber).String(), func(t *testing.T) {
			got, err := GetZeroCharacteristicSequence(tt.maxNumber, tt.positional)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetZeroCharacteristicSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.wantName {
				t.Errorf("GetZeroCharacteristicSequence() Name = %v, want %v", got.Name, tt.wantName)
			}
			if !reflect.DeepEqual(got.Sequence, tt.wantSeq) {
				t.Errorf("GetZeroCharacteristicSequence() Sequence = %v, want %v", got.Sequence, tt.wantSeq)
			}
			if got.Result.Cmp(tt.wantResult) != 0 {
				t.Errorf("GetZeroCharacteristicSequence() Result = %v, want %v", got.Result, tt.wantResult)
			}
		})
	}
}
