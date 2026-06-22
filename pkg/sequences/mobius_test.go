package sequences

import (
	"testing"
)

func TestMobiusCalculator(t *testing.T) {
	mc := NewMobiusCalculator(100)

	tests := []struct {
		n    int
		want int8
	}{
		{1, 1},
		{2, -1},
		{3, -1},
		{4, 0},
		{5, -1},
		{6, 1},
		{7, -1},
		{8, 0},
		{9, 0},
		{10, 1},
		{30, -1}, // 2*3*5 -> (-1)^3 = -1
	}

	for _, tt := range tests {
		got := mc.GetMu(tt.n)
		if got != tt.want {
			t.Errorf("GetMu(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

func TestDivisorSummation(t *testing.T) {
	mc := NewMobiusCalculator(100)

	tests := []int{1, 2, 6, 12, 30, 60}
	for _, n := range tests {
		got := mc.DivisorSummation(n)
		want := 0
		if n == 1 {
			want = 1
		}
		if got != want {
			t.Errorf("DivisorSummation(%d) = %d; want %d", n, got, want)
		}
	}
}
