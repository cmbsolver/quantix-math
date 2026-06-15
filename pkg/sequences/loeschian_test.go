package sequences

import (
	"math/big"
	"testing"
)

func TestGenerateLoeschianSequence(t *testing.T) {
	// Values from OEIS A003136: 0, 1, 3, 4, 7, 9, 12, 13, 16, 19, 21, 25, 27, 28, 31, 36, 37, 39, 43, 48, 49, 52, 57, 61, 63, 64, 67, 73, 75, 76, 79, 81, 84, 91, 93, 97, 100, 103, 108, 109, 111, 112, 117, 121, 124, 127, 129, 133, 139, 144, 147, 148, 151, 156, 157, 163, 169, 171, 172, 175, 181, 183, 189, 192
	expected := []int64{0, 1, 3, 4, 7, 9, 12, 13, 16, 19, 21, 25, 27, 28, 31, 36, 37, 39, 43, 48, 49, 52, 57, 61, 63, 64, 67, 73, 75, 76, 79, 81, 84, 91, 93, 97, 100, 103, 108, 109, 111, 112, 117, 121, 124, 127, 129, 133, 139, 144, 147, 148, 151, 156, 157, 163, 169, 171, 172, 175, 181, 183, 189, 192}

	maxNumber := big.NewInt(192)
	res, err := GenerateLoeschianSequence(maxNumber)
	if err != nil {
		t.Fatalf("GenerateLoeschianSequence failed: %v", err)
	}

	if len(res.Sequence) != len(expected) {
		t.Errorf("Expected %d terms, got %d", len(expected), len(res.Sequence))
	}

	for i, v := range res.Sequence {
		if v.Int64() != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], v.Int64())
		}
	}
}

func TestIsLoeschian(t *testing.T) {
	testCases := []struct {
		n        int64
		expected bool
	}{
		{0, true},
		{1, true},
		{2, false}, // 2 is 2 mod 3, exponent is 1 (odd)
		{3, true},
		{4, true}, // 2^2, 2 is 2 mod 3, exponent is 2 (even)
		{7, true},
		{9, true},
		{12, true}, // 2^2 * 3
		{13, true},
		{16, true}, // 2^4
		{19, true},
		{21, true},
		{25, true},
		{27, true},
		{28, true}, // 2^2 * 7
		{31, true},
		{42, false}, // 2 * 3 * 7 -> 2 has exponent 1
		{49, true},
		{61, true},
		{192, true},
	}

	for _, tc := range testCases {
		res := IsLoeschian(big.NewInt(tc.n))
		if res != tc.expected {
			t.Errorf("IsLoeschian(%d) = %v; expected %v", tc.n, res, tc.expected)
		}
	}
}
