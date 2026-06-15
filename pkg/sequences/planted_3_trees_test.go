package sequences

import (
	"math/big"
	"testing"
)

func TestGeneratePlanted3TreesSequence(t *testing.T) {
	tests := []struct {
		maxNumber int64
		expected  []string
	}{
		{
			maxNumber: 5,
			expected:  []string{"1", "1", "2", "7", "56", "2212"},
		},
		{
			maxNumber: 6,
			expected:  []string{"1", "1", "2", "7", "56", "2212", "2595782"},
		},
	}

	for _, tt := range tests {
		seq, err := GeneratePlanted3TreesSequence(big.NewInt(tt.maxNumber))
		if err != nil {
			t.Errorf("GeneratePlanted3TreesSequence(%d) error: %v", tt.maxNumber, err)
			continue
		}

		if len(seq.Sequence) != int(tt.maxNumber)+1 {
			t.Errorf("GeneratePlanted3TreesSequence(%d) length: expected %d, got %d", tt.maxNumber, tt.maxNumber+1, len(seq.Sequence))
		}

		for i, val := range seq.Sequence {
			if val.String() != tt.expected[i] {
				t.Errorf("GeneratePlanted3TreesSequence(%d) at index %d: expected %s, got %s", tt.maxNumber, i, tt.expected[i], val.String())
			}
		}
	}
}

func TestIsPlanted3Tree(t *testing.T) {
	tests := []struct {
		number   string
		expected bool
	}{
		{"1", true},
		{"2", true},
		{"7", true},
		{"56", true},
		{"2212", true},
		{"2595782", true},
		{"3", false},
		{"8", false},
		{"100", false},
	}

	for _, tt := range tests {
		n := new(big.Int)
		n.SetString(tt.number, 10)
		exists, _ := IsPlanted3Tree(n)
		if exists != tt.expected {
			t.Errorf("IsPlanted3Tree(%s): expected %v, got %v", tt.number, tt.expected, exists)
		}
	}
}
