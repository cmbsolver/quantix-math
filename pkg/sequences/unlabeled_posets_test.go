package sequences

import (
	"math/big"
	"testing"
)

func TestGetUnlabeledPosetsSequence(t *testing.T) {
	tests := []struct {
		n        int64
		expected string
	}{
		{0, "1"},
		{1, "1"},
		{2, "2"},
		{3, "5"},
		{4, "16"},
		{5, "63"},
		{6, "318"},
		{7, "2045"},
		{8, "16999"},
	}

	for _, tt := range tests {
		maxNumber := big.NewInt(tt.n)
		seq, err := GetUnlabeledPosetsSequence(maxNumber, false)
		if err != nil {
			t.Errorf("GetUnlabeledPosetsSequence(%d, false) error: %v", tt.n, err)
			continue
		}
		if seq.Result.String() != tt.expected {
			t.Errorf("GetUnlabeledPosetsSequence(%d, false) = %s; want %s", tt.n, seq.Result.String(), tt.expected)
		}
	}
}

func TestGetUnlabeledPosetsAtPosition(t *testing.T) {
	tests := []struct {
		n        int64
		expected string
	}{
		{0, "1"},
		{1, "1"},
		{2, "2"},
		{3, "5"},
		{4, "16"},
		{5, "63"},
		{6, "318"},
		{7, "2045"},
		{8, "16999"},
	}

	for _, tt := range tests {
		n := big.NewInt(tt.n)
		seq, err := GetUnlabeledPosetsSequence(n, true)
		if err != nil {
			t.Errorf("GetUnlabeledPosetsSequence(%d, true) error: %v", tt.n, err)
			continue
		}
		if seq.Result.String() != tt.expected {
			t.Errorf("GetUnlabeledPosetsSequence(%d, true) = %s; want %s", tt.n, seq.Result.String(), tt.expected)
		}
	}
}
