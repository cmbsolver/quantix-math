package sequences

import (
	"math/big"
	"testing"
)

func TestGetSimplicialPolyhedraSequence(t *testing.T) {
	tests := []struct {
		maxNumber   int64
		positional  bool
		expected    string
		expectError bool
	}{
		{3, false, "1", false},
		{4, false, "1", false},
		{5, false, "1", false},
		{6, false, "2", false},
		{7, false, "5", false},
		{8, false, "14", false},
		{9, false, "50", false},
		{10, false, "233", false},
		{11, false, "1249", false},
		{12, false, "7595", false},
		{3, true, "1", false},
		{4, true, "1", false},
		{12, true, "7595", false},
		{23, true, "28615703421545", false},
		{2, false, "", true},
		{24, false, "", true},
	}

	for _, tt := range tests {
		maxBig := big.NewInt(tt.maxNumber)
		seq, err := GetSimplicialPolyhedraSequence(maxBig, tt.positional)

		if tt.expectError {
			if err == nil {
				t.Errorf("Expected error for n=%d, positional=%v, but got none", tt.maxNumber, tt.positional)
			}
			continue
		}

		if err != nil {
			t.Errorf("Unexpected error for n=%d, positional=%v: %v", tt.maxNumber, tt.positional, err)
			continue
		}

		if seq.Result.String() != tt.expected {
			t.Errorf("For n=%d, positional=%v, expected %s, but got %s", tt.maxNumber, tt.positional, tt.expected, seq.Result.String())
		}
	}
}
