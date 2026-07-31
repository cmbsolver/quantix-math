package sequences

import (
	"math/big"
	"testing"
)

func TestCalculateA010060(t *testing.T) {
	tests := []struct {
		n    int64
		want int64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 0},
		{4, 1},
		{5, 0},
		{6, 0},
		{7, 1},
		{8, 1},
		{9, 0},
		{10, 0},
		{11, 1},
		{12, 0},
		{13, 1},
		{14, 1},
		{15, 0},
	}

	for _, tt := range tests {
		if got := CalculateA010060(tt.n); got != tt.want {
			t.Errorf("CalculateA010060(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestGenerateA010060Sequence_OEISExample(t *testing.T) {
	seq, err := GetA010060Sequence(big.NewInt(16), false)
	if err != nil {
		t.Fatalf("GetA010060Sequence returned error: %v", err)
	}

	expected := []int64{0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0}
	if len(seq.Sequence) != len(expected) {
		t.Fatalf("expected sequence length %d, got %d", len(expected), len(seq.Sequence))
	}

	for i, want := range expected {
		if seq.Sequence[i].Int64() != want {
			t.Errorf("sequence[%d] = %d, want %d", i, seq.Sequence[i].Int64(), want)
		}
	}

	if seq.Result.Int64() != 0 {
		t.Fatalf("result = %d, want 0", seq.Result.Int64())
	}
}

func TestGetA010060AtPosition(t *testing.T) {
	seq, err := GetA010060Sequence(big.NewInt(31), true)
	if err != nil {
		t.Fatalf("GetA010060Sequence returned error: %v", err)
	}

	if seq.Result.Int64() != 1 {
		t.Fatalf("result at position 31 = %d, want 1", seq.Result.Int64())
	}
}

func TestCheckExistenceA010060(t *testing.T) {
	exists, index, err := checkExistence(big.NewInt(1), "thue_morse_a010060")
	if err != nil {
		t.Fatalf("checkExistence returned error: %v", err)
	}
	if !exists {
		t.Fatalf("checkExistence did not find 1 in A010060")
	}
	if index == "" {
		t.Fatalf("checkExistence returned empty index for A010060")
	}
}
