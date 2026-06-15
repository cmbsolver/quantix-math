package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateLCM1ToNSequence(t *testing.T) {
	tests := []struct {
		maxNumber int64
		want      []int64
	}{
		{0, []int64{1}},
		{1, []int64{1, 1}},
		{2, []int64{1, 1, 2}},
		{3, []int64{1, 1, 2, 6}},
		{4, []int64{1, 1, 2, 6, 12}},
		{5, []int64{1, 1, 2, 6, 12, 60}},
		{10, []int64{1, 1, 2, 6, 12, 60, 60, 420, 840, 2520, 2520}},
	}

	for _, tt := range tests {
		maxNum := big.NewInt(tt.maxNumber)
		seq, err := GenerateLCM1ToNSequence(maxNum)
		if err != nil {
			t.Errorf("GenerateLCM1ToNSequence(%d) error = %v", tt.maxNumber, err)
			continue
		}

		if len(seq.Sequence) != len(tt.want) {
			t.Errorf("GenerateLCM1ToNSequence(%d) got sequence length %d, want %d", tt.maxNumber, len(seq.Sequence), len(tt.want))
			continue
		}

		for i, v := range tt.want {
			if seq.Sequence[i].Int64() != v {
				t.Errorf("GenerateLCM1ToNSequence(%d) sequence[%d] = %v, want %d", tt.maxNumber, i, seq.Sequence[i], v)
			}
		}
	}
}

func TestGetLCM1ToNAtPosition(t *testing.T) {
	tests := []struct {
		n    int64
		want int64
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 12},
		{5, 60},
		{10, 2520},
	}

	for _, tt := range tests {
		n := big.NewInt(tt.n)
		seq, err := GetLCM1ToNAtPosition(n)
		if err != nil {
			t.Errorf("GetLCM1ToNAtPosition(%d) error = %v", tt.n, err)
			continue
		}

		if seq.Result.Int64() != tt.want {
			t.Errorf("GetLCM1ToNAtPosition(%d) = %v, want %d", tt.n, seq.Result, tt.want)
		}
	}
}

func TestIsLCM1ToN(t *testing.T) {
	tests := []struct {
		n        int64
		isLCM    bool
		expected string
	}{
		{1, true, "0 or 1"},
		{2, true, "2"},
		{6, true, "3"},
		{12, true, "4"},
		{60, true, "5-6"},
		{420, true, "7"},
		{840, true, "8"},
		{2520, true, "9-10"},
		{3, false, ""},
		{4, false, ""},
		{10, false, ""},
	}

	for _, tt := range tests {
		n := big.NewInt(tt.n)
		isLCM, pos := IsLCM1ToN(n)
		if isLCM != tt.isLCM {
			t.Errorf("IsLCM1ToN(%d) isLCM = %v, want %v", tt.n, isLCM, tt.isLCM)
		}
		if pos != tt.expected {
			t.Errorf("IsLCM1ToN(%d) pos = %s, want %s", tt.n, pos, tt.expected)
		}
	}
}
