package sequences

import (
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
