package sequences

import (
	"math/big"
	"testing"
)

func TestGetTriangularSequence(t *testing.T) {
	tests := []struct {
		maxNumber *big.Int
		expected  []*big.Int
	}{
		{
			maxNumber: big.NewInt(0),
			expected:  []*big.Int{big.NewInt(0)},
		},
		{
			maxNumber: big.NewInt(10),
			expected:  []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(3), big.NewInt(6), big.NewInt(10)},
		},
		{
			maxNumber: big.NewInt(55),
			expected:  []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(3), big.NewInt(6), big.NewInt(10), big.NewInt(15), big.NewInt(21), big.NewInt(28), big.NewInt(36), big.NewInt(45), big.NewInt(55)},
		},
	}

	for _, tt := range tests {
		seq, err := GetTriangularSequence(tt.maxNumber, false)
		if err != nil {
			t.Errorf("GetTriangularSequence(%v) returned error: %v", tt.maxNumber, err)
			continue
		}

		if len(seq.Sequence) != len(tt.expected) {
			t.Errorf("GetTriangularSequence(%v) length = %d, expected %d", tt.maxNumber, len(seq.Sequence), len(tt.expected))
			continue
		}

		for i, v := range seq.Sequence {
			if v.Cmp(tt.expected[i]) != 0 {
				t.Errorf("GetTriangularSequence(%v) at index %d = %v, expected %v", tt.maxNumber, i, v, tt.expected[i])
			}
		}
	}
}

func TestGetTriangularAtPosition(t *testing.T) {
	tests := []struct {
		n        *big.Int
		expected *big.Int
	}{
		{n: big.NewInt(0), expected: big.NewInt(0)},
		{n: big.NewInt(1), expected: big.NewInt(1)},
		{n: big.NewInt(2), expected: big.NewInt(3)},
		{n: big.NewInt(3), expected: big.NewInt(6)},
		{n: big.NewInt(10), expected: big.NewInt(55)},
	}

	for _, tt := range tests {
		seq, err := GetTriangularSequence(tt.n, true)
		if err != nil {
			t.Errorf("GetTriangularSequence(%v, true) returned error: %v", tt.n, err)
			continue
		}

		if seq.Result.Cmp(tt.expected) != 0 {
			t.Errorf("GetTriangularSequence(%v, true) = %v, expected %v", tt.n, seq.Result, tt.expected)
		}
	}
}
