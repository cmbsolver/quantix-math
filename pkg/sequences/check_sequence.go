package sequences

import (
	"fmt"
	"math/big"
)

// CheckResult represents the result of checking if a number exists in sequences.
type CheckResult struct {
	SequenceName string `json:"sequenceName"`
	Exists       bool   `json:"exists"`
	Index        string `json:"index,omitempty"` // Position in sequence if known
}

// CheckNumberInSequences checks all available sequences for the given number.
func CheckNumberInSequences(numberStr string) ([]CheckResult, error) {
	n := new(big.Int)
	n, ok := n.SetString(numberStr, 10)
	if !ok {
		return nil, fmt.Errorf("invalid number: %s", numberStr)
	}

	sequenceTypes := []string{
		"groups_order_n_a000001",
		"zero_a000004",
		"necklaces_color_swap_turnover_a000011",
		"all_ones_a000012",
		"necklaces_color_swap_a000013",
		"series_reduced_trees_a000014",
		"smallest_prime_power_a000015",
		"shift_register_sequences_a000016",
		"point_symmetric_queens_a000017",
		"form_x2_16y2_a000018",
		"primitive_permutation_groups_a000019",
		"form_x2_12y2_a000021",
		"centered_hydrocarbons_a000022",
		"exp_minus_2x_a000023",
		"form_x2_10y2_a000024",
		"mock_theta_f_q_a000025",
		"mosaic_numbers_a000026",
		"positive_integers_a000027",
		"binary_weight_odd_a000028",
		"necklaces_turnover_a000029",
		"initial_digit_a000030",
		"necklaces_no_turnover_a000031",
		"lucas_numbers_a000032",
		"menage_hit_polynomials_a000033",
		"period_12_a000034",
		"parity_a000035",
		"record_values_p_n_a000036",
		"nonsquares_a000037",
		"twice_characteristic_0_a000038",
		"mock_theta_f_q_coeff_a000039",
		"partitions_a000041",
		"unary_a000042",
		"mersenne_exponents_a000043",
		"dying_rabbits_a000044",
		"fibonacci_numbers_a000045",
		"primitive_necklaces_complement_a000046",
		"form_x2_minus_2y2_a000047",
		"primitive_necklaces_color_swap_a000048",
		"form_3x2_4y2_a000049",
		"form_x2_y2_a000050",
		"2n_plus_1_a000051",
		"alphabetical_123_a000052",
		"nyc_subway_1_a000053",
		"nyc_subway_a_a000054",
		"unlabeled_trees_a000055",
		"sl2_zn_order_a000056",
		"fibonacci_dividing_primes_a000057",
		"sylvester_a000058",
		"form_2k4_plus_1_primes_a000059",
		"signed_trees_a000060",
		"generalized_tangent_a000061",
		"beatty_e_minus_2_a000062",
		"symmetrical_dissections_a000063",
		"change_1_2_5_10_a000064",
		"partitions_minus_1_a000065",
		"trivalent_graph_girth_a000066",
		"form_x2_2y2_a000067",
		"form_k4_plus_1_primes_a000068",
		"odious_a000069",
		"sum_partitions_a000070",
		"fibonacci_minus_1_a000071",
		"form_x2_4y2_a000072",
		"tribonacci_a000073",
		"odd_form_x2_y2_a000074",
		"form_2x2_3y2_a000075",
		"form_4x2_4xy_5y2_a000076",
		"form_x2_6y2_a000077",
		"tetranacci_a000078",
		"powers_of_2_a000079",
		"minimal_triangle_graphs_a000080",
		"rooted_unlabeled_trees_a000081",
		"n2_phi_phi_a000082",
		"mixed_husimi_trees_a000083",
		"series_parallel_networks_a000084",
		"involutions_a000085",
		"solutions_x2_x_1_mod_n_a000086",
		"unrooted_maps_a000087",
		"simple_unlabeled_graphs_a000088",
		"solutions_x2_1_mod_n_a000089",
		"exp_minus_x3_3_a000090",
		"multiplicative_a000091",
		"record_values_p_n_3d_a000092",
		"floor_n_1_5_a000093",
		"trees_diameter_4_a000094",
		"fixed_points_gamma0_n_a000095",
		"n_n_plus_3_2_a000096",
		"primitive_polynomials_a000020",
		"partitions_2kinds_1_2_a000097",
		"partitions_2kinds_1_2_3_a000098",
		"record_values_p_n_2d_a000099",
		"schroeder_second_a001003",
		"motzkin_numbers_a001006",
		"crystal_ball_squashed_d5_lattice_a010025",
		"thue_morse_a010060",
		"cotesian_numerators_a100640",
		"cotesian_denominators_a100641",
		"cotesian_numerator_c_n_1_a100643",
		"cotesian_denominator_c_n_1_a100644",
		"cotesian_numerator_c_n_2_a100645",
		"cotesian_denominator_c_n_2_a100646",
		"cotesian_numerator_c_n_3_a100647",
		"sqrt_prime_a000006",
		"binary_quadratic_forms_a000003",
		"hamming_weight",
		"central_polygonal",
		"squares",
		"cubes",
		"natural",
		"parity",
		"prime",
		"primes_a000040",
		"emirp",
		"semiprime",
		"circular_prime",
		"fibonacci_prime",
		"cake",
		"bell",
		"catalan",
		"totient",
		"totient_prime",
		"fibonacci",
		"pell",
		"zekendorf",
		"lucas",
		"nn",
		"schroeder_fourth",
		"partitions_distinct",
		"partitions",
		"partitions_a000041",
		"partitions_into_2_squares",
		"plane_partitions",
		"tangent",
		"kendall_mann",
		"pentagonal",
		"square_pyramidal",
		"euler",
		"euler_zigzag",
		"perfect",
		"groups_order_n",
		"modular_j",
		"ramanujan_tau",
		"fourth_powers",
		"tetrahedral",
		"triangular",
		"sum_divisors",
		"sum_odd_divisors",
		"alkanes",
		"abelian_groups_order_n",
		"threshold_functions",
		"fubini",
		"kolakoski",
		"zero",
		"zero_characteristic_a000007",
		"divisor_count_a000005",
		"divisor_count",
		"change_1_2_5_10_a000008",
		"collatz",
		"powers_of_2",
		"powers_of_4",
		"powers_of_3",
		"odious_numbers",
		"subfactorial",
		"binary_partitions",
		"binary_rooted_trees",
		"sqrt3_convergents",
		"sqrt3_convergents_denominators",
		"factorial",
		"planted_3_trees",
		"rooted_unlabeled_trees",
		"unlabeled_trees",
		"unlabeled_digraphs",
		"unlabeled_graphs",
		"connected_planar_graphs",
		"unlabeled_posets",
		"bicolorable_necklaces",
		"simplicial_polyhedra",
		"labeled_rooted_trees",
		"labeled_trees",
		"sets_of_lists",
		"free_polyominoes",
		"self_inverse_permutations",
		"sylvester",
		"theta_series_square_lattice",
		"theta_series_d4_lattice",
		"mersenne_numbers",
		"mersenne_prime_exponents",
		"radon_hurwitz",
		"lcm_1_to_n",
		"loeschian",
		"composites",
		"quarter_squares",
		"ways_two_squares",
		"stern",
	}

	var results []CheckResult
	for _, st := range sequenceTypes {
		exists, index, err := checkExistence(n, st)
		if err == nil && exists {
			results = append(results, CheckResult{
				SequenceName: getSequenceName(st),
				Exists:       exists,
				Index:        index,
			})
		}
	}

	return results, nil
}

func getSequenceName(st string) string {
	switch st {
	case "hamming_weight":
		return "Hamming weight (A000120)"
	case "powers_of_3":
		return "Powers of 3 (A000244)"
	case "powers_of_2":
		return "Powers of 2 (A000079)"
	case "subfactorial":
		return "Subfactorial (A000166)"
	case "binary_partitions":
		return "Binary partitions (A000123)"
	case "binary_rooted_trees":
		return "Binary rooted trees (A002572)"
	case "sqrt3_convergents":
		return "Numerators of convergents to sqrt(3) (A002531)"
	case "kendall_mann":
		return "Kendall-Mann numbers (A000140)"
	case "bicolorable_necklaces":
		return "Bicolorable Primitive Necklaces (A000048)"
	case "factorial":
		return "Factorial numbers (A000142)"
	case "unlabeled_trees":
		return "Unlabeled Trees (A000055)"
	case "unlabeled_digraphs":
		return "Unlabeled Directed Graphs (A000273)"
	case "unlabeled_graphs":
		return "Unlabeled Graphs (A000088)"
	case "connected_planar_graphs":
		return "Connected Planar Graphs (A003094)"
	case "unlabeled_posets":
		return "Unlabeled Posets (A000112)"
	case "simplicial_polyhedra":
		return "Simplicial polyhedra (A000109)"
	case "labeled_rooted_trees":
		return "Labeled Rooted Trees (A000169)"
	case "plane_partitions":
		return "Plane partitions (A000219)"
	case "tangent":
		return "Tangent Numbers (A000182)"
	case "euler_zigzag":
		return "Euler zigzag numbers (A000111)"
	case "mersenne_numbers":
		return "Mersenne numbers (A000225)"
	case "mersenne_prime_exponents":
		return "Mersenne prime exponents (A000043)"
	case "mersenne_exponents_a000043":
		return "Mersenne prime exponents (A000043)"
	case "sylvester":
		return "Sylvester's sequence (A000058)"
	case "labeled_trees":
		return "Labeled Trees (A000272)"
	case "sets_of_lists":
		return "Sets of Lists (A000262)"
	case "free_polyominoes":
		return "Free Polyominoes (A000105)"
	case "self_inverse_permutations":
		return "Self-inverse permutations (A000085)"
	case "natural":
		return "Natural"
	case "parity":
		return "n mod 2; parity of n (A000035)"
	case "partitions_distinct":
		return "Partitions into distinct parts (A000009)"
	case "partitions_a000041":
		return "Partitions of n (A000041)"
	case "partitions":
		return "Partitions of n (A000041)"
	case "squares":
		return "Squares (A000290)"
	case "cubes":
		return "Cubes (A000578)"
	case "prime":
		return "The prime numbers (A000040)"
	case "primes_a000040":
		return "The prime numbers (A000040)"
	case "emirp":
		return "Emirp"
	case "semiprime":
		return "Semi-prime"
	case "circular_prime":
		return "Circular Prime"
	case "fibonacci":
		return "Fibonacci"
	case "pell":
		return "Pell numbers (A000129)"
	case "lucas":
		return "Lucas"
	case "pentagonal":
		return "Pentagonal numbers (A000326)"
	case "square_pyramidal":
		return "Square pyramidal numbers (A000330)"
	case "euler":
		return "Euler numbers (A000364)"
	case "perfect":
		return "Perfect numbers (A000396)"
	case "fourth_powers":
		return "Fourth Powers (A000583)"
	case "triangular":
		return "Triangular numbers (A000217)"
	case "modular_j":
		return "Modular function j (A000521)"
	case "central_polygonal":
		return "Central polygonal numbers (the Lazy Caterer's sequence) (A000124)"
	case "cake":
		return "Cake"
	case "bell":
		return "Bell numbers (A000110)"
	case "catalan":
		return "Catalan"
	case "totient":
		return "Totient"
	case "totient_prime":
		return "Totient Prime"
	case "fibonacci_prime":
		return "Fibonacci Prime"
	case "zekendorf":
		return "Zekendorf Representation"
	case "groups_order_n_a000001":
		return "Number of groups of order n (A000001)"
	case "nn":
		return "n^n (A000312)"
	case "powers_of_4":
		return "Powers of 4 (A000302)"
	case "tetrahedral":
		return "Tetrahedral numbers (A000292)"
	case "schroeder_second_a001003":
		return "Schroeder's second problem (A001003)"
	case "schroeder_fourth":
		return "Schroeder's fourth problem (A000311)"
	case "ramanujan_tau":
		return "Ramanujan's tau function (A000594)"
	case "sum_odd_divisors":
		return "Sum of Odd Divisors (A000593)"
	case "alkanes":
		return "Number of Alkanes (A000602)"
	case "abelian_groups_order_n":
		return "Number of Abelian groups (A000688)"
	case "threshold_functions":
		return "Threshold functions (A000609)"
	case "fubini":
		return "Fubini numbers (A000670)"
	case "kolakoski":
		return "Kolakoski Sequence (A000002)"
	case "zero_a000004":
		return "Zero Sequence (A000004)"
	case "zero_characteristic_a000007":
		return "Zero Characteristic (A000007)"
	case "partitions_into_2_squares":
		return "Partitions into 2 squares (A000161)"
	case "divisor_count_a000005":
		return "Number of Divisors (A000005)"
	case "divisor_count":
		return "Number of Divisors (A000005)"
	case "sum_divisors":
		return "Sum of Divisors (A000203)"
	case "change_1_2_5_10_a000008":
		return "Ways to Make Change (A000008)"
	case "rooted_unlabeled_trees":
		return "Rooted Unlabeled Trees (A000081)"
	case "collatz":
		return "Collatz"
	case "theta_series_square_lattice":
		return "Theta series of square lattice (A004018)"
	case "theta_series_d4_lattice":
		return "Theta series of D_4 lattice (A004011)"
	case "radon_hurwitz":
		return "Radon function (A003484)"
	case "lcm_1_to_n":
		return "Least common multiple of {1, 2, ..., n} (A003418)"
	case "loeschian":
		return "Loeschian numbers (A003136)"
	case "planted_3_trees":
		return "Planted 3-trees of height n (A002658)"
	case "quarter_squares":
		return "Quarter-squares (A002620)"
	case "composites":
		return "Composite numbers (A002808)"
	case "ways_two_squares":
		return "Ways as sum of at most two nonzero squares (A002654)"
	case "sqrt3_convergents_denominators":
		return "Sqrt(3) Convergents Denominators (A002530)"
	case "odious_numbers":
		return "Odious numbers (A000069)"
	case "stern":
		return "Stern's diatomic series (A002487)"
	case "sqrt_prime_a000006":
		return "Integer part of square root of n-th prime (A000006)"
	case "binary_quadratic_forms_a000003":
		return "Binary quadratic forms (A000003)"
	case "point_symmetric_queens_a000017":
		return "Point symmetric queens (A000017)"
	case "primitive_permutation_groups_a000019":
		return "Primitive permutation groups (A000019)"
	case "primitive_polynomials_a000020":
		return "Primitive polynomials of degree n over GF(2) (A000020)"
	case "form_x2_16y2_a000018":
		return "Numbers of form x^2 + 16y^2 (A000018)"
	case "form_x2_12y2_a000021":
		return "Numbers of form x^2 + 12y^2 (A000021)"
	case "centered_hydrocarbons_a000022":
		return "Centered hydrocarbons (A000022)"
	case "form_x2_10y2_a000024":
		return "Numbers of form x^2 + 10y^2 (A000024)"
	case "mock_theta_f_q_a000025":
		return "Mock theta function f(q) (A000025)"
	case "mosaic_numbers_a000026":
		return "Mosaic numbers (A000026)"
	case "positive_integers_a000027":
		return "Positive integers (A000027)"
	case "initial_digit_a000030":
		return "Initial digit (A000030)"
	case "necklaces_turnover_a000029":
		return "Bracelets with n beads of 2 colors (A000029)"
	case "binary_weight_odd_a000028":
		return "Binary weight of exponents is odd (A000028)"
	case "exp_minus_2x_a000023":
		return "Expansion of e.g.f. exp(-2x)/(1-x) (A000023)"
	case "crystal_ball_squashed_d5_lattice_a010025":
		return "Crystal ball sequence for squashed {D_5}^* lattice (A010025)"
	case "thue_morse_a010060":
		return "Thue-Morse sequence (A010060)"
	case "cotesian_numerators_a100640":
		return "Numerators of Cotesian numbers C(n,k) (A100640)"
	case "cotesian_denominators_a100641":
		return "Denominators of Cotesian numbers C(n,k) (A100641)"
	case "cotesian_numerator_c_n_1_a100643":
		return "Numerator of Cotesian number C(n,1) (A100643)"
	case "cotesian_denominator_c_n_1_a100644":
		return "Denominator of Cotesian number C(n,1) (A100644)"
	case "cotesian_numerator_c_n_2_a100645":
		return "Numerator of Cotesian number C(n,2) (A100645)"
	case "cotesian_denominator_c_n_2_a100646":
		return "Denominator of Cotesian number C(n,2) (A100646)"
	case "cotesian_numerator_c_n_3_a100647":
		return "Numerator of Cotesian number C(n,3) (A100647)"
	default:
		return st
	}
}

func checkExistence(n *big.Int, st string) (bool, string, error) {
	switch st {
	case "triangular":
		// n is a triangular number iff 8n + 1 is a perfect square
		if n.Sign() >= 0 {
			v := new(big.Int).Mul(big.NewInt(8), n)
			v.Add(v, big.NewInt(1))
			if isPerfectSquare(v) {
				// (sqrt(8n+1)-1)/2 = position
				root := new(big.Int).Sqrt(v)
				pos := new(big.Int).Sub(root, big.NewInt(1))
				pos.Div(pos, big.NewInt(2))
				return true, pos.String(), nil
			}
		}
	case "natural":
		if n.Sign() >= 0 {
			return true, n.String(), nil
		}
	case "parity":
		if n.Cmp(big.NewInt(0)) == 0 {
			return true, "even indices", nil
		}
		if n.Cmp(big.NewInt(1)) == 0 {
			return true, "odd indices", nil
		}
	case "squares":
		if n.Sign() >= 0 {
			root := new(big.Int).Sqrt(n)
			if new(big.Int).Mul(root, root).Cmp(n) == 0 {
				return true, root.String(), nil
			}
		}
	case "fourth_powers":
		if n.Sign() >= 0 {
			root := new(big.Int).Sqrt(n)
			if new(big.Int).Mul(root, root).Cmp(n) == 0 {
				root2 := new(big.Int).Sqrt(root)
				if new(big.Int).Mul(root2, root2).Cmp(root) == 0 {
					return true, root2.String(), nil
				}
			}
		}
	case "sqrt3_convergents_denominators":
		// For sequences without an easy closed form for inverse, we can generate up to n or a reasonable limit
		// A002530 grows exponentially, so it's fast to generate up to n.
		if n.Sign() < 0 {
			return false, "", nil
		}
		for i := int64(0); ; i++ {
			term := calculateSqrt3ConvergentsDenominator(i)
			if term.Cmp(n) == 0 {
				return true, fmt.Sprintf("%d", i), nil
			}
			if term.Cmp(n) > 0 {
				break
			}
		}
	case "odious_numbers":
		if IsOdious(n) {
			return true, "", nil
		}
	case "prime", "primes_a000040":
		if IsPrime(n) {
			return true, "", nil
		}
	case "emirp":
		if IsEmirp(n) {
			return true, "", nil
		}
	case "semiprime":
		if IsSemiPrime(n) {
			return true, "", nil
		}
	case "circular_prime":
		if IsCircularPrime(n) {
			return true, "", nil
		}
	case "fibonacci":
		// A number is Fibonacci if and only if one or both of (5*n^2 + 4) or (5*n^2 – 4) is a perfect square
		if n.Sign() < 0 {
			return false, "", nil
		}
		n2 := new(big.Int).Mul(n, n)
		fiveN2 := new(big.Int).Mul(big.NewInt(5), n2)

		v1 := new(big.Int).Add(fiveN2, big.NewInt(4))
		v2 := new(big.Int).Sub(fiveN2, big.NewInt(4))

		if isPerfectSquare(v1) || isPerfectSquare(v2) {
			return true, "", nil
		}
	case "pell":
		// A number is a Pell number if and only if 8*n^2 + 1 or 8*n^2 - 1 is a perfect square
		if n.Sign() < 0 {
			return false, "", nil
		}
		n2 := new(big.Int).Mul(n, n)
		eightN2 := new(big.Int).Mul(big.NewInt(8), n2)
		v1 := new(big.Int).Add(eightN2, big.NewInt(1))
		v2 := new(big.Int).Sub(eightN2, big.NewInt(1))
		if isPerfectSquare(v1) || isPerfectSquare(v2) {
			return true, "", nil
		}
	case "zero_a000004":
		if n.Cmp(big.NewInt(0)) == 0 {
			return true, "any", nil
		}
	case "zero_characteristic_a000007":
		if n.Cmp(big.NewInt(0)) == 0 {
			return true, "any n > 0", nil
		}
		if n.Cmp(big.NewInt(1)) == 0 {
			return true, "0", nil
		}
	case "theta_series_square_lattice":
		if n.Sign() >= 0 {
			val := n.Int64()
			if countWaysToSumOf2Squares(val) > 0 {
				return true, "", nil
			}
		}
	case "loeschian":
		if IsLoeschian(n) {
			return true, "", nil
		}
	case "planted_3_trees":
		if exists, pos := IsPlanted3Tree(n); exists {
			return true, pos, nil
		}
	case "quarter_squares":
		if n.Sign() >= 0 {
			// n is a quarter-square if it's of the form floor(k^2/4).
			// k = 2m   => floor(4m^2/4) = m^2
			// k = 2m+1 => floor((4m^2+4m+1)/4) = m^2 + m = m(m+1)
			// So n is a quarter-square if n = m^2 or n = m(m+1) for some m.
			// m^2 <= n => m = floor(sqrt(n))
			m := new(big.Int).Sqrt(n)
			m2 := new(big.Int).Mul(m, m)
			if m2.Cmp(n) == 0 {
				// n = m^2, so k = 2m
				k := new(big.Int).Mul(m, big.NewInt(2))
				return true, k.String(), nil
			}
			mPlus1 := new(big.Int).Add(m, big.NewInt(1))
			mTimesMPlus1 := new(big.Int).Mul(m, mPlus1)
			if mTimesMPlus1.Cmp(n) == 0 {
				// n = m(m+1), so k = 2m+1
				k := new(big.Int).Mul(m, big.NewInt(2))
				k.Add(k, big.NewInt(1))
				return true, k.String(), nil
			}
		}
	case "composites":
		if n.Cmp(big.NewInt(4)) >= 0 && !IsPrime(n) {
			return true, "", nil
		}
	case "stern":
		exists, pos := IsSternNumber(n)
		if exists {
			return true, pos, nil
		}
	case "kolakoski":
		exists, pos := IsKolakoskiNumber(n)
		if exists {
			return true, pos, nil
		}
	case "binary_quadratic_forms_a000003":
		// Fallback: generate sequence and check
		seq, err := GetBinaryQuadraticFormsA000003Sequence(n, false)
		if err == nil && seq != nil {
			for i, val := range seq.Sequence {
				if val.Cmp(n) == 0 {
					return true, fmt.Sprintf("%d", i+1), nil
				}
			}
		}
	case "radon_hurwitz":
		if n.Sign() > 0 {
			return true, "", nil
		}
	case "lcm_1_to_n":
		exists, pos := IsLCM1ToN(n)
		if exists {
			return true, pos, nil
		}
	case "cubes", "central_polygonal", "cake", "bell", "catalan", "totient", "totient_prime",
		"hamming_weight",
		"binary_rooted_trees",
		"sqrt3_convergents",
		"sqrt_prime_a000006",
		"factorial", "kendall_mann", "bicolorable_necklaces", "binary_partitions", "partitions",
		"fibonacci_prime", "zekendorf", "groups_order_n_a000001", "ramanujan_tau",
		"primitive_polynomials_a000020",
		"centered_hydrocarbons_a000022",
		"sum_divisors", "divisor_count", "divisor_count_a000005", "sum_odd_divisors", "alkanes", "abelian_groups_order_n",
		"theta_series_d4_lattice",
		"ways_two_squares",
		"threshold_functions", "fubini", "schroeder_fourth", "schroeder_second_a001003",
		"partitions_into_2_squares",
		"powers_of_2",
		"powers_of_4",
		"tangent", "self_inverse_permutations",
		"euler_zigzag",
		"sylvester",
		"mersenne_numbers", "mersenne_prime_exponents", "rooted_unlabeled_trees", "unlabeled_trees", "unlabeled_digraphs", "unlabeled_graphs", "connected_planar_graphs",
		"change_1_2_5_10_a000008",
		"unlabeled_posets", "simplicial_polyhedra", "labeled_trees", "sets_of_lists", "free_polyominoes":
		// Fallback: generate sequence and check (with reasonable limit)
		seq, err := GetSequence(n.String(), st, false)
		if err == nil && seq != nil {
			for i, val := range seq.Sequence {
				if val.Cmp(n) == 0 {
					return true, fmt.Sprintf("%d", i), nil
				}
			}
		}
	case "point_symmetric_queens_a000017":
		return IsPointSymmetricQueensA000017(n)
	case "primitive_permutation_groups_a000019":
		return IsPrimitivePermutationGroupA000019(n)
	case "necklaces_turnover_a000029":
		// A000029 sequence can grow fast, but let's check first few terms or use a reasonable limit
		seq, err := GetA000029Sequence(big.NewInt(100), false)
		if err == nil {
			for i, term := range seq.Sequence {
				if term.Cmp(n) == 0 {
					return true, fmt.Sprintf("%d", i), nil
				}
				if term.Cmp(n) > 0 {
					break
				}
			}
		}
	case "exp_minus_2x_a000023":
		return IsA000023(n)
	case "mock_theta_f_q_a000025":
		// Fallback: generate sequence and check
		seq, err := GetA000025Sequence(n, false)
		if err == nil && seq != nil {
			for i, val := range seq.Sequence {
				if val.Cmp(n) == 0 {
					return true, fmt.Sprintf("%d", i), nil
				}
			}
		}
	case "mosaic_numbers_a000026":
		// Fallback: generate sequence and check
		seq, err := GetA000026Sequence(n, false)
		if err == nil && seq != nil {
			for i, val := range seq.Sequence {
				if val.Cmp(n) == 0 {
					return true, fmt.Sprintf("%d", i+1), nil
				}
			}
		}
	case "positive_integers_a000027":
		if n.Sign() > 0 {
			return true, n.String(), nil
		}
	case "initial_digit_a000030":
		if n.Sign() >= 0 {
			// Initial digit of n is n itself for 0 <= n <= 9.
			// For n > 9, a(n) is the first digit.
			// Any single digit number k can be the initial digit of many numbers.
			// The question is "does n exist in A000030?".
			// The terms of A000030 are always 0-9.
			if n.Cmp(big.NewInt(9)) <= 0 {
				return true, "", nil
			}
		}
	case "binary_weight_odd_a000028":
		if isA000028(n.Int64()) {
			return true, "", nil
		}
	case "form_x2_10y2_a000024":
		if n.Sign() <= 0 {
			return false, "", nil
		}
		// m = x^2 + 10y^2
		for y := int64(0); ; y++ {
			y2_10 := 10 * y * y
			if big.NewInt(y2_10).Cmp(n) > 0 {
				break
			}
			rem := new(big.Int).Sub(n, big.NewInt(y2_10))
			if isPerfectSquare(rem) {
				return true, "", nil
			}
		}
	case "form_x2_16y2_a000018":
		if n.Sign() <= 0 {
			return false, "", nil
		}
		// m = x^2 + 16y^2
		for y := int64(0); ; y++ {
			y2_16 := 16 * y * y
			if big.NewInt(y2_16).Cmp(n) > 0 {
				break
			}
			rem := new(big.Int).Sub(n, big.NewInt(y2_16))
			if isPerfectSquare(rem) {
				return true, "", nil
			}
		}
	case "form_x2_12y2_a000021":
		if n.Sign() <= 0 {
			return false, "", nil
		}
		// m = x^2 + 12y^2
		for y := int64(0); ; y++ {
			y2_12 := 12 * y * y
			if big.NewInt(y2_12).Cmp(n) > 0 {
				break
			}
			rem := new(big.Int).Sub(n, big.NewInt(y2_12))
			if isPerfectSquare(rem) {
				return true, "", nil
			}
		}
	case "crystal_ball_squashed_d5_lattice_a010025":
		if n.Sign() <= 0 {
			return false, "", nil
		}
		limit := int64(1)
		for {
			term := CalculateA010025(int(limit))
			termBig := new(big.Int).SetUint64(term)
			cmp := termBig.Cmp(n)
			if cmp == 0 {
				return true, fmt.Sprintf("%d", limit), nil
			}
			if cmp > 0 {
				break
			}
			limit *= 2
		}
		low, high := int64(0), limit
		for low <= high {
			mid := (low + high) / 2
			term := CalculateA010025(int(mid))
			termBig := new(big.Int).SetUint64(term)
			cmp := termBig.Cmp(n)
			if cmp == 0 {
				return true, fmt.Sprintf("%d", mid), nil
			}
			if cmp < 0 {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	case "thue_morse_a010060":
		if n.Cmp(big.NewInt(0)) == 0 {
			return true, "even binary-weight indices", nil
		}
		if n.Cmp(big.NewInt(1)) == 0 {
			return true, "odd binary-weight indices", nil
		}
	case "cotesian_numerators_a100640":
		seq, err := GenerateCotesianNumeratorsA100640Sequence(big.NewInt(1000))
		if err != nil {
			return false, "", err
		}
		for i, v := range seq.Sequence {
			if v.Cmp(n) == 0 {
				return true, fmt.Sprintf("%d", i), nil
			}
		}
	case "cotesian_denominators_a100641":
		seq, err := GenerateCotesianDenominatorsA100641Sequence(big.NewInt(1000))
		if err != nil {
			return false, "", err
		}
		for i, v := range seq.Sequence {
			if v.Cmp(n) == 0 {
				return true, fmt.Sprintf("%d", i), nil
			}
		}
	case "cotesian_numerator_c_n_1_a100643":
		seq, err := GenerateCotesianNumeratorCN1A100643Sequence(big.NewInt(1000))
		if err != nil {
			return false, "", err
		}
		for i, v := range seq.Sequence {
			if v.Cmp(n) == 0 {
				return true, fmt.Sprintf("%d", i), nil
			}
		}
	case "cotesian_denominator_c_n_1_a100644":
		seq, err := GenerateCotesianDenominatorCN1A100644Sequence(big.NewInt(1000))
		if err != nil {
			return false, "", err
		}
		for i, v := range seq.Sequence {
			if v.Cmp(n) == 0 {
				return true, fmt.Sprintf("%d", i), nil
			}
		}
	case "cotesian_numerator_c_n_2_a100645":
		seq, err := GenerateCotesianNumeratorCN2A100645Sequence(big.NewInt(1000))
		if err != nil {
			return false, "", err
		}
		for i, v := range seq.Sequence {
			if v.Cmp(n) == 0 {
				return true, fmt.Sprintf("%d", i), nil
			}
		}
	case "cotesian_denominator_c_n_2_a100646":
		seq, err := GenerateCotesianDenominatorCN2A100646Sequence(big.NewInt(1000))
		if err != nil {
			return false, "", err
		}
		for i, v := range seq.Sequence {
			if v.Cmp(n) == 0 {
				return true, fmt.Sprintf("%d", i), nil
			}
		}
	case "cotesian_numerator_c_n_3_a100647":
		seq, err := GenerateCotesianNumeratorCN3A100647Sequence(big.NewInt(1000))
		if err != nil {
			return false, "", err
		}
		for i, v := range seq.Sequence {
			if v.Cmp(n) == 0 {
				return true, fmt.Sprintf("%d", i), nil
			}
		}
	}
	return false, "", nil
}

func isPerfectSquare(n *big.Int) bool {
	if n.Sign() < 0 {
		return false
	}
	root := new(big.Int).Sqrt(n)
	return new(big.Int).Mul(root, root).Cmp(n) == 0
}
