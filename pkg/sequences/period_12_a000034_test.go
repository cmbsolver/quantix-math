package sequences

import (
	"math/big"
	"testing"
)

func TestGeneratePeriod12A000034Sequence_OEISExample(t *testing.T) {
	got, err := GeneratePeriod12A000034Sequence(big.NewInt(10))
	if err != nil {
		t.Fatalf("GeneratePeriod12A000034Sequence returned error: %v", err)
	}

	want := []int64{1, 2, 1, 2, 1, 2, 1, 2, 1, 2}
	if len(got.Sequence) != len(want) {
		t.Fatalf("sequence length = %d, want %d", len(got.Sequence), len(want))
	}

	for i, expected := range want {
		if got.Sequence[i].Cmp(big.NewInt(expected)) != 0 {
			t.Fatalf("term %d = %s, want %d", i, got.Sequence[i].String(), expected)
		}
	}
}
