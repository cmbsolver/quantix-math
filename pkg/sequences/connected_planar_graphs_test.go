package sequences

import (
	"math/big"
	"testing"
)

func TestGetConnectedPlanarGraphsSequence(t *testing.T) {
	tests := []struct {
		n        int64
		expected string
	}{
		{0, "1"},
		{1, "1"},
		{2, "1"},
		{3, "2"},
		{4, "6"},
		{5, "20"},
		{6, "99"},
		{7, "646"},
		{8, "5974"},
		{9, "71885"},
		{10, "1052805"},
	}

	for _, tt := range tests {
		seq, err := GetConnectedPlanarGraphsSequence(big.NewInt(tt.n), false)
		if err != nil {
			t.Errorf("GetConnectedPlanarGraphsSequence(%d) error: %v", tt.n, err)
			continue
		}
		if seq.Result.String() != tt.expected {
			t.Errorf("GetConnectedPlanarGraphsSequence(%d) = %s; want %s", tt.n, seq.Result.String(), tt.expected)
		}
	}
}

func TestGetConnectedPlanarGraphsAtPosition(t *testing.T) {
	tests := []struct {
		n        int64
		expected string
	}{
		{3, "2"},
		{5, "20"},
		{10, "1052805"},
	}

	for _, tt := range tests {
		seq, err := GetConnectedPlanarGraphsSequence(big.NewInt(tt.n), true)
		if err != nil {
			t.Errorf("GetConnectedPlanarGraphsAtPosition(%d) error: %v", tt.n, err)
			continue
		}
		if seq.Result.String() != tt.expected {
			t.Errorf("GetConnectedPlanarGraphsAtPosition(%d) = %s; want %s", tt.n, seq.Result.String(), tt.expected)
		}
	}
}
