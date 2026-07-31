package sequences

import (
	"math/big"
	"testing"
)

func TestCheckNumberInSequences(t *testing.T) {
	tests := []struct {
		number   string
		wantSeqs []string
	}{
		{
			number:   "4",
			wantSeqs: []string{"Natural", "Lucas", "Central Polygonal Numbers", "Cake"},
		},
		{
			number:   "8",
			wantSeqs: []string{"Natural", "Fibonacci"}, // 8 is not prime.
		},
		{
			number:   "13",
			wantSeqs: []string{"Natural", "Prime", "Fibonacci"},
		},
		{
			number:   "81",
			wantSeqs: []string{"Natural", "Fourth Powers (A000583)"},
		},
		{
			number:   "0",
			wantSeqs: []string{"Natural", "Fourth Powers (A000583)", "Zero Sequence (A000004)", "Zero Characteristic (A000007)"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.number, func(t *testing.T) {
			got, err := CheckNumberInSequences(tt.number)
			if err != nil {
				t.Fatalf("CheckNumberInSequences(%s) error = %v", tt.number, err)
			}

			for _, want := range tt.wantSeqs {
				found := false
				for _, res := range got {
					if res.SequenceName == want {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("CheckNumberInSequences(%s) result did not contain %s, got %v", tt.number, want, got)
				}
			}
		})
	}
}

func TestCheckNumberInNewSequences(t *testing.T) {
	tests := []struct {
		number   string
		wantName string
	}{
		{"13", "Emirp"},
		{"4", "Semi-prime"},
		{"197", "Circular Prime"},
		{"7", "Prime"},
		{"16", "Powers of 4 (A000302)"},
		{"1", "Natural"},
		{"1", "Kolakoski Sequence (A000002)"},
		{"2", "Kolakoski Sequence (A000002)"},
	}

	for _, tt := range tests {
		t.Run(tt.number, func(t *testing.T) {
			results, err := CheckNumberInSequences(tt.number)
			if err != nil {
				t.Fatalf("CheckNumberInSequences(%s) error = %v", tt.number, err)
			}
			found := false
			for _, res := range results {
				if res.SequenceName == tt.wantName {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("CheckNumberInSequences(%s) did not find sequence %s, results: %v", tt.number, tt.wantName, results)
			}
		})
	}
}

func TestCheckExistenceJacobsthalA001045(t *testing.T) {
	valueIn := big.NewInt(43)
	valueOut := big.NewInt(44)

	exists, _, err := checkExistence(valueIn, "jacobsthal_numbers_a001045")
	if err != nil {
		t.Fatalf("checkExistence(%s, jacobsthal_numbers_a001045) unexpected error: %v", valueIn.String(), err)
	}
	if !exists {
		t.Fatalf("expected %s to exist in Jacobsthal numbers", valueIn.String())
	}

	exists, _, err = checkExistence(valueOut, "jacobsthal_numbers_a001045")
	if err != nil {
		t.Fatalf("checkExistence(%s, jacobsthal_numbers_a001045) unexpected error: %v", valueOut.String(), err)
	}
	if exists {
		t.Fatalf("expected %s not to exist in Jacobsthal numbers", valueOut.String())
	}
}
