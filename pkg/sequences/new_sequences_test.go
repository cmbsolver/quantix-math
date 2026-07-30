package sequences

import (
	"math/big"
	"testing"
)

func TestNewSequences(t *testing.T) {
	tests := []struct {
		seqType    string
		maxNum     string
		positional bool
		wantRes    string
	}{
		{"initial_digit_a000030", "123", true, "1"},
		{"initial_digit_a000030", "5", false, "4"}, // n=0,1,2,3,4 -> 0,1,2,3,4. Result is last term: 4
		{"period_12_a000034", "0", true, "1"},
		{"period_12_a000034", "1", true, "2"},
		{"parity_a000035", "10", true, "0"},
		{"parity_a000035", "11", true, "1"},
		{"unary_a000042", "3", true, "111"},
		{"mersenne_exponents_a000043", "8", true, "31"},
		{"form_x2_y2_a000050", "5", false, "9"}, // A000050[4] = 9 (n=0..4)
		{"2n_plus_1_a000051", "5", true, "33"},
		{"sylvester_a000058", "3", true, "43"},
		{"beatty_e_minus_2_a000062", "5", true, "6"},
		{"odious_a000069", "5", true, "11"},
		{"tribonacci_a000073", "5", true, "4"},
		{"tetranacci_a000078", "5", true, "2"},
		{"floor_n_1_5_a000093", "5", true, "11"},
		{"n_n_plus_3_2_a000096", "5", true, "20"},
		{"partitions_minus_1_a000065", "5", true, "6"},
		{"partitions_a000041", "5", true, "7"},
		{"fibonacci_minus_1_a000071", "4", true, "7"}, // a(4) = Fib(6)-1 = 8-1 = 7
	}

	for _, tt := range tests {
		t.Run(tt.seqType, func(t *testing.T) {
			got, err := GetSequence(tt.maxNum, tt.seqType, tt.positional)
			if err != nil {
				t.Errorf("GetSequence(%s) error = %v", tt.seqType, err)
				return
			}
			if got.Result.String() != tt.wantRes {
				t.Errorf("GetSequence(%s) result = %v, want %v", tt.seqType, got.Result, tt.wantRes)
			}
		})
	}
}

func TestSequenceDropdownOptionIncludesOEISForA000043(t *testing.T) {
	options := GetSequenceDropdownOptions()
	for _, option := range options {
		if option.Value == "mersenne_exponents_a000043" {
			if option.Label != "mersenne_exponents_a000043 (OEIS A000043)" {
				t.Fatalf("label = %q, want %q", option.Label, "mersenne_exponents_a000043 (OEIS A000043)")
			}
			return
		}
	}

	t.Fatalf("sequence option %q not found", "mersenne_exponents_a000043")
}

func TestOEISLookupsExtended(t *testing.T) {
	tests := []struct {
		id      string
		pos     int64
		wantVal string
	}{
		{"A000021", 1, "1"},
		{"A000026", 10, "10"},
		{"A000037", 1, "2"},
		{"A000050", 1, "1"},
		{"A000050", 36, "5473203125"},
		{"A000052", 66, "72"},
		{"A000053", 29, "242"},
		{"A000055", 37, "6226306037178"},
		{"A000088", 20, "24637809253125004524383007491432768"},
		{"A000099", 46, "28560"},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			got, err := GetOEISLookupSequence(tt.id, "test", big.NewInt(tt.pos), true)
			if err != nil {
				t.Errorf("GetOEISLookupSequence(%s) error = %v", tt.id, err)
				return
			}
			if got.Result.String() != tt.wantVal {
				t.Errorf("GetOEISLookupSequence(%s) result = %v, want %v", tt.id, got.Result, tt.wantVal)
			}
		})
	}
}
