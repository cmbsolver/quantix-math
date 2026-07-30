package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetUnaryA000042Sequence_OEISExample(t *testing.T) {
	got, err := GetUnaryA000042Sequence(big.NewInt(5), false)
	if err != nil {
		t.Fatalf("GetUnaryA000042Sequence() error = %v", err)
	}

	want := []*big.Int{
		big.NewInt(1),
		big.NewInt(11),
		big.NewInt(111),
		big.NewInt(1111),
		big.NewInt(11111),
	}

	if !reflect.DeepEqual(got.Sequence, want) {
		t.Fatalf("GetUnaryA000042Sequence() = %v, want %v", got.Sequence, want)
	}
}
