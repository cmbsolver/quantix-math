package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetBellSequence(t *testing.T) {
	// A000110: 1, 1, 2, 5, 15, 52, 203, 877, 4140, 21147
	expectedValues := []int64{1, 1, 2, 5, 15, 52, 203, 877, 4140, 21147}
	expected := make([]*big.Int, len(expectedValues))
	for i, v := range expectedValues {
		expected[i] = big.NewInt(v)
	}

	maxNumber := big.NewInt(int64(len(expectedValues) - 1))
	seq, err := GetBellSequence(maxNumber, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(seq.Sequence, expected) {
		t.Errorf("Expected sequence %v, got %v", expected, seq.Sequence)
	}

	if seq.Result.Cmp(expected[len(expected)-1]) != 0 {
		t.Errorf("Expected result %v, got %v", expected[len(expected)-1], seq.Result)
	}
}

func TestGetBellNumberAtPosition(t *testing.T) {
	tests := []struct {
		position int64
		expected int64
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 5},
		{4, 15},
		{5, 52},
		{10, 115975},
	}

	for _, tt := range tests {
		n := big.NewInt(tt.position)
		seq, err := GetBellSequence(n, true)
		if err != nil {
			t.Errorf("Position %d: unexpected error: %v", tt.position, err)
			continue
		}
		if seq.Result.Cmp(big.NewInt(tt.expected)) != 0 {
			t.Errorf("Position %d: expected %d, got %v", tt.position, tt.expected, seq.Result)
		}
	}
}
