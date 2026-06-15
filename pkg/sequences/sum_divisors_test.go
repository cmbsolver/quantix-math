package sequences

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetSumDivisorsSequence(t *testing.T) {
	tests := []struct {
		maxNumber    int64
		isPositional bool
		expected     []int64
		expectErr    bool
	}{
		{
			maxNumber:    10,
			isPositional: false,
			expected:     []int64{1, 3, 4, 7, 6, 12, 8, 15, 13, 18},
			expectErr:    false,
		},
		{
			maxNumber:    12,
			isPositional: true,
			expected:     []int64{28},
			expectErr:    false,
		},
		{
			maxNumber:    1,
			isPositional: false,
			expected:     []int64{1},
			expectErr:    false,
		},
		{
			maxNumber:    0,
			isPositional: false,
			expected:     nil,
			expectErr:    true,
		},
	}

	for _, tt := range tests {
		n := big.NewInt(tt.maxNumber)
		result, err := GetSumDivisorsSequence(n, tt.isPositional)

		if tt.expectErr {
			if err == nil {
				t.Errorf("GetSumDivisorsSequence(%d, %v) expected error, got nil", tt.maxNumber, tt.isPositional)
			}
			continue
		}

		if err != nil {
			t.Errorf("GetSumDivisorsSequence(%d, %v) unexpected error: %v", tt.maxNumber, tt.isPositional, err)
			continue
		}

		actual := make([]int64, len(result.Sequence))
		for i, v := range result.Sequence {
			actual[i] = v.Int64()
		}

		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("GetSumDivisorsSequence(%d, %v) = %v, want %v", tt.maxNumber, tt.isPositional, actual, tt.expected)
		}
	}
}

func TestCalculateSumDivisors(t *testing.T) {
	tests := []struct {
		n        int64
		expected int64
	}{
		{1, 1},
		{2, 3},
		{3, 4},
		{4, 7},
		{5, 6},
		{6, 12},
		{12, 28},
		{24, 60},
		{25, 31},
		{28, 56},
		{30, 72},
		{100, 217},
	}

	for _, tt := range tests {
		actual := calculateSumDivisors(tt.n)
		if actual != tt.expected {
			t.Errorf("calculateSumDivisors(%d) = %d, want %d", tt.n, actual, tt.expected)
		}
	}
}
