package sequences

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"unicode"
)

type SequenceOption struct {
	Value string
	Label string
}

var sequenceDropdownOptions = []SequenceOption{
	{Value: "groups_order_n_a000001", Label: "groups_order_n_a000001"},
	{Value: "zero_a000004", Label: "zero_a000004"},
	{Value: "necklaces_color_swap_turnover_a000011", Label: "necklaces_color_swap_turnover_a000011"},
	{Value: "all_ones_a000012", Label: "all_ones_a000012"},
	{Value: "necklaces_color_swap_a000013", Label: "necklaces_color_swap_a000013"},
	{Value: "series_reduced_trees_a000014", Label: "series_reduced_trees_a000014"},
	{Value: "smallest_prime_power_a000015", Label: "smallest_prime_power_a000015"},
	{Value: "shift_register_sequences_a000016", Label: "shift_register_sequences_a000016"},
	{Value: "point_symmetric_queens_a000017", Label: "point_symmetric_queens_a000017"},
	{Value: "form_x2_16y2_a000018", Label: "form_x2_16y2_a000018"},
	{Value: "primitive_permutation_groups_a000019", Label: "primitive_permutation_groups_a000019"},
	{Value: "form_x2_12y2_a000021", Label: "form_x2_12y2_a000021"},
	{Value: "centered_hydrocarbons_a000022", Label: "centered_hydrocarbons_a000022"},
	{Value: "exp_minus_2x_a000023", Label: "exp_minus_2x_a000023"},
	{Value: "form_x2_10y2_a000024", Label: "form_x2_10y2_a000024"},
	{Value: "mock_theta_f_q_a000025", Label: "mock_theta_f_q_a000025"},
	{Value: "mosaic_numbers_a000026", Label: "mosaic_numbers_a000026"},
	{Value: "positive_integers_a000027", Label: "positive_integers_a000027"},
	{Value: "binary_weight_odd_a000028", Label: "binary_weight_odd_a000028"},
	{Value: "necklaces_turnover_a000029", Label: "necklaces_turnover_a000029"},
	{Value: "initial_digit_a000030", Label: "initial_digit_a000030"},
	{Value: "necklaces_no_turnover_a000031", Label: "necklaces_no_turnover_a000031 (OEIS A000031)"},
	{Value: "lucas_numbers_a000032", Label: "lucas_numbers_a000032 (OEIS A000032)"},
	{Value: "menage_hit_polynomials_a000033", Label: "menage_hit_polynomials_a000033 (OEIS A000033)"},
	{Value: "period_12_a000034", Label: "period_12_a000034 (OEIS A000034)"},
	{Value: "parity_a000035", Label: "parity_a000035 (OEIS A000035)"},
	{Value: "record_values_p_n_a000036", Label: "record_values_p_n_a000036 (OEIS A000036)"},
	{Value: "nonsquares_a000037", Label: "nonsquares_a000037 (OEIS A000037)"},
	{Value: "twice_characteristic_0_a000038", Label: "twice_characteristic_0_a000038 (OEIS A000038)"},
	{Value: "mock_theta_f_q_coeff_a000039", Label: "mock_theta_f_q_coeff_a000039 (OEIS A000039)"},
	{Value: "unary_a000042", Label: "unary_a000042"},
	{Value: "mersenne_exponents_a000043", Label: "mersenne_exponents_a000043"},
	{Value: "dying_rabbits_a000044", Label: "dying_rabbits_a000044"},
	{Value: "primitive_necklaces_complement_a000046", Label: "primitive_necklaces_complement_a000046"},
	{Value: "form_x2_minus_2y2_a000047", Label: "form_x2_minus_2y2_a000047"},
	{Value: "primitive_necklaces_color_swap_a000048", Label: "primitive_necklaces_color_swap_a000048"},
	{Value: "form_3x2_4y2_a000049", Label: "form_3x2_4y2_a000049"},
	{Value: "form_x2_y2_a000050", Label: "form_x2_y2_a000050"},
	{Value: "2n_plus_1_a000051", Label: "2n_plus_1_a000051"},
	{Value: "alphabetical_123_a000052", Label: "alphabetical_123_a000052"},
	{Value: "nyc_subway_1_a000053", Label: "nyc_subway_1_a000053"},
	{Value: "nyc_subway_a_a000054", Label: "nyc_subway_a_a000054"},
	{Value: "unlabeled_trees_a000055", Label: "unlabeled_trees_a000055"},
	{Value: "sl2_zn_order_a000056", Label: "sl2_zn_order_a000056"},
	{Value: "fibonacci_dividing_primes_a000057", Label: "fibonacci_dividing_primes_a000057"},
	{Value: "sylvester_a000058", Label: "sylvester_a000058"},
	{Value: "form_2k4_plus_1_primes_a000059", Label: "form_2k4_plus_1_primes_a000059"},
	{Value: "signed_trees_a000060", Label: "signed_trees_a000060"},
	{Value: "generalized_tangent_a000061", Label: "generalized_tangent_a000061"},
	{Value: "beatty_e_minus_2_a000062", Label: "beatty_e_minus_2_a000062"},
	{Value: "symmetrical_dissections_a000063", Label: "symmetrical_dissections_a000063"},
	{Value: "change_1_2_5_10_a000064", Label: "change_1_2_5_10_a000064"},
	{Value: "partitions_minus_1_a000065", Label: "partitions_minus_1_a000065"},
	{Value: "trivalent_graph_girth_a000066", Label: "trivalent_graph_girth_a000066"},
	{Value: "form_x2_2y2_a000067", Label: "form_x2_2y2_a000067"},
	{Value: "form_k4_plus_1_primes_a000068", Label: "form_k4_plus_1_primes_a000068"},
	{Value: "odious_a000069", Label: "odious_a000069"},
	{Value: "sum_partitions_a000070", Label: "sum_partitions_a000070"},
	{Value: "fibonacci_minus_1_a000071", Label: "fibonacci_minus_1_a000071"},
	{Value: "form_x2_4y2_a000072", Label: "form_x2_4y2_a000072"},
	{Value: "tribonacci_a000073", Label: "tribonacci_a000073"},
	{Value: "odd_form_x2_y2_a000074", Label: "odd_form_x2_y2_a000074"},
	{Value: "form_2x2_3y2_a000075", Label: "form_2x2_3y2_a000075"},
	{Value: "form_4x2_4xy_5y2_a000076", Label: "form_4x2_4xy_5y2_a000076"},
	{Value: "form_x2_6y2_a000077", Label: "form_x2_6y2_a000077"},
	{Value: "tetranacci_a000078", Label: "tetranacci_a000078"},
	{Value: "powers_of_2_a000079", Label: "powers_of_2_a000079"},
	{Value: "minimal_triangle_graphs_a000080", Label: "minimal_triangle_graphs_a000080"},
	{Value: "rooted_unlabeled_trees_a000081", Label: "rooted_unlabeled_trees_a000081"},
	{Value: "n2_phi_phi_a000082", Label: "n2_phi_phi_a000082"},
	{Value: "mixed_husimi_trees_a000083", Label: "mixed_husimi_trees_a000083"},
	{Value: "series_parallel_networks_a000084", Label: "series_parallel_networks_a000084"},
	{Value: "involutions_a000085", Label: "involutions_a000085"},
	{Value: "solutions_x2_x_1_mod_n_a000086", Label: "solutions_x2_x_1_mod_n_a000086"},
	{Value: "unrooted_maps_a000087", Label: "unrooted_maps_a000087"},
	{Value: "simple_unlabeled_graphs_a000088", Label: "simple_unlabeled_graphs_a000088"},
	{Value: "solutions_x2_1_mod_n_a000089", Label: "solutions_x2_1_mod_n_a000089"},
	{Value: "exp_minus_x3_3_a000090", Label: "exp_minus_x3_3_a000090"},
	{Value: "multiplicative_a000091", Label: "multiplicative_a000091"},
	{Value: "record_values_p_n_3d_a000092", Label: "record_values_p_n_3d_a000092"},
	{Value: "floor_n_1_5_a000093", Label: "floor_n_1_5_a000093"},
	{Value: "trees_diameter_4_a000094", Label: "trees_diameter_4_a000094"},
	{Value: "fixed_points_gamma0_n_a000095", Label: "fixed_points_gamma0_n_a000095"},
	{Value: "n_n_plus_3_2_a000096", Label: "n_n_plus_3_2_a000096"},
	{Value: "primitive_polynomials_a000020", Label: "primitive_polynomials_a000020"},
	{Value: "partitions_2kinds_1_2_a000097", Label: "partitions_2kinds_1_2_a000097"},
	{Value: "partitions_2kinds_1_2_3_a000098", Label: "partitions_2kinds_1_2_3_a000098"},
	{Value: "record_values_p_n_2d_a000099", Label: "record_values_p_n_2d_a000099"},
	{Value: "sqrt_prime_a000006", Label: "sqrt_prime_a000006"},
	{Value: "binary_quadratic_forms_a000003", Label: "binary_quadratic_forms_a000003"},
	{Value: "hamming_weight", Label: "hamming_weight"},
	{Value: "central_polygonal", Label: "central_polygonal"},
	{Value: "squares", Label: "squares"},
	{Value: "cubes", Label: "cubes"},
	{Value: "natural", Label: "natural"},
	{Value: "parity", Label: "parity"},
	{Value: "prime", Label: "prime"},
	{Value: "primes_a000040", Label: "primes_a000040"},
	{Value: "emirp", Label: "emirp"},
	{Value: "semiprime", Label: "semiprime"},
	{Value: "circular_prime", Label: "circular_prime"},
	{Value: "fibonacci_prime", Label: "fibonacci_prime"},
	{Value: "cake", Label: "cake"},
	{Value: "bell", Label: "bell"},
	{Value: "catalan", Label: "catalan"},
	{Value: "totient", Label: "totient"},
	{Value: "totient_prime", Label: "totient_prime"},
	{Value: "fibonacci", Label: "fibonacci"},
	{Value: "pell", Label: "pell"},
	{Value: "zekendorf", Label: "zekendorf"},
	{Value: "lucas", Label: "lucas"},
	{Value: "nn", Label: "nn"},
	{Value: "schroeder_fourth", Label: "schroeder_fourth"},
	{Value: "partitions_distinct", Label: "partitions_distinct"},
	{Value: "partitions", Label: "partitions"},
	{Value: "partitions_into_2_squares", Label: "partitions_into_2_squares"},
	{Value: "plane_partitions", Label: "plane_partitions"},
	{Value: "tangent", Label: "tangent"},
	{Value: "kendall_mann", Label: "kendall_mann"},
	{Value: "pentagonal", Label: "pentagonal"},
	{Value: "square_pyramidal", Label: "square_pyramidal"},
	{Value: "euler", Label: "euler"},
	{Value: "euler_zigzag", Label: "euler_zigzag"},
	{Value: "perfect", Label: "perfect"},
	{Value: "groups_order_n", Label: "groups_order_n"},
	{Value: "modular_j", Label: "modular_j"},
	{Value: "ramanujan_tau", Label: "ramanujan_tau"},
	{Value: "fourth_powers", Label: "fourth_powers"},
	{Value: "tetrahedral", Label: "tetrahedral"},
	{Value: "triangular", Label: "triangular"},
	{Value: "sum_divisors", Label: "sum_divisors"},
	{Value: "sum_odd_divisors", Label: "sum_odd_divisors"},
	{Value: "alkanes", Label: "alkanes"},
	{Value: "abelian_groups_order_n", Label: "abelian_groups_order_n"},
	{Value: "threshold_functions", Label: "threshold_functions"},
	{Value: "fubini", Label: "fubini"},
	{Value: "kolakoski", Label: "kolakoski"},
	{Value: "zero", Label: "zero"},
	{Value: "zero_characteristic_a000007", Label: "zero_characteristic_a000007"},
	{Value: "divisor_count_a000005", Label: "divisor_count_a000005"},
	{Value: "divisor_count", Label: "divisor_count"},
	{Value: "change_1_2_5_10_a000008", Label: "change_1_2_5_10_a000008"},
	{Value: "collatz", Label: "collatz"},
	{Value: "powers_of_2", Label: "powers_of_2"},
	{Value: "powers_of_4", Label: "powers_of_4"},
	{Value: "powers_of_3", Label: "powers_of_3"},
	{Value: "odious_numbers", Label: "odious_numbers"},
	{Value: "subfactorial", Label: "subfactorial"},
	{Value: "binary_partitions", Label: "binary_partitions"},
	{Value: "binary_rooted_trees", Label: "binary_rooted_trees"},
	{Value: "sqrt3_convergents", Label: "sqrt3_convergents"},
	{Value: "sqrt3_convergents_denominators", Label: "sqrt3_convergents_denominators"},
	{Value: "factorial", Label: "factorial"},
	{Value: "planted_3_trees", Label: "planted_3_trees"},
	{Value: "rooted_unlabeled_trees", Label: "rooted_unlabeled_trees"},
	{Value: "unlabeled_trees", Label: "unlabeled_trees"},
	{Value: "unlabeled_digraphs", Label: "unlabeled_digraphs"},
	{Value: "unlabeled_graphs", Label: "unlabeled_graphs"},
	{Value: "connected_planar_graphs", Label: "connected_planar_graphs"},
	{Value: "unlabeled_posets", Label: "unlabeled_posets"},
	{Value: "bicolorable_necklaces", Label: "bicolorable_necklaces"},
	{Value: "simplicial_polyhedra", Label: "simplicial_polyhedra"},
	{Value: "labeled_rooted_trees", Label: "labeled_rooted_trees"},
	{Value: "labeled_trees", Label: "labeled_trees"},
	{Value: "sets_of_lists", Label: "sets_of_lists"},
	{Value: "free_polyominoes", Label: "free_polyominoes"},
	{Value: "self_inverse_permutations", Label: "self_inverse_permutations"},
	{Value: "sylvester", Label: "sylvester"},
	{Value: "theta_series_square_lattice", Label: "theta_series_square_lattice"},
	{Value: "theta_series_d4_lattice", Label: "theta_series_d4_lattice"},
	{Value: "mersenne_numbers", Label: "mersenne_numbers"},
	{Value: "mersenne_prime_exponents", Label: "mersenne_prime_exponents"},
	{Value: "radon_hurwitz", Label: "radon_hurwitz"},
	{Value: "lcm_1_to_n", Label: "lcm_1_to_n"},
	{Value: "loeschian", Label: "loeschian"},
	{Value: "composites", Label: "composites"},
	{Value: "quarter_squares", Label: "quarter_squares"},
	{Value: "ways_two_squares", Label: "ways_two_squares"},
	{Value: "stern", Label: "stern"},
}

func init() {
	for i := range sequenceDropdownOptions {
		if sequenceDropdownOptions[i].Label == "" || sequenceDropdownOptions[i].Label == sequenceDropdownOptions[i].Value {
			sequenceDropdownOptions[i].Label = humanizeSequenceLabel(sequenceDropdownOptions[i].Value)
		}
	}
}

func GetSequenceDropdownOptions() []SequenceOption {
	options := make([]SequenceOption, len(sequenceDropdownOptions))
	copy(options, sequenceDropdownOptions)
	for i := range options {
		if options[i].Label == "" || options[i].Label == options[i].Value {
			options[i].Label = humanizeSequenceLabel(options[i].Value)
		}
	}
	return options
}

var oeisSuffixPattern = regexp.MustCompile(`_a(\d{6})$`)
var variablePowerPattern = regexp.MustCompile(`^([a-z])(\d+)$`)
var coefficientPowerPattern = regexp.MustCompile(`^(\d+)([a-z]+)(\d+)$`)
var dimensionPattern = regexp.MustCompile(`^(\d+)d$`)

var tokenOverrides = map[string]string{
	"nn":  "n^n",
	"nyc": "NYC",
	"sl2": "SL2",
	"zn":  "Zn",
}

func humanizeSequenceToken(token string) string {
	lowerToken := strings.ToLower(token)
	if override, ok := tokenOverrides[lowerToken]; ok {
		return override
	}

	if matches := coefficientPowerPattern.FindStringSubmatch(lowerToken); len(matches) == 4 {
		return matches[1] + matches[2] + "^" + matches[3]
	}

	if matches := variablePowerPattern.FindStringSubmatch(lowerToken); len(matches) == 3 {
		return matches[1] + "^" + matches[2]
	}

	if matches := dimensionPattern.FindStringSubmatch(lowerToken); len(matches) == 2 {
		return matches[1] + "D"
	}

	runes := []rune(token)
	if len(runes) == 0 {
		return token
	}
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func humanizeSequenceLabel(value string) string {
	base := value
	oeisLabel := ""

	if matches := oeisSuffixPattern.FindStringSubmatch(value); len(matches) == 2 {
		base = strings.TrimSuffix(value, matches[0])
		oeisLabel = " (A" + matches[1] + ")"
	}

	words := strings.Split(strings.ReplaceAll(base, "_", " "), " ")
	for i := range words {
		if words[i] == "" {
			continue
		}
		words[i] = humanizeSequenceToken(words[i])
	}

	return strings.Join(words, " ") + oeisLabel
}

func GetSequence(maxNumberString, sequenceType string, positional bool) (*NumericSequence, error) {
	maxNumber := new(big.Int)
	maxNumber, ok := maxNumber.SetString(maxNumberString, 10)
	if !ok {
		fmt.Printf("Invalid max number: %s\n", maxNumberString)
		return nil, fmt.Errorf("invalid max number: %s", maxNumberString)
	}

	var sequence *NumericSequence
	var err error

	switch sequenceType {
	case "groups_order_n_a000001":
		sequence, err = GetGroupsOrderNSequence(maxNumber, positional)
	case "zero_a000004":
		sequence, err = GetZeroSequence(maxNumber, positional)
	case "necklaces_color_swap_turnover_a000011":
		sequence, err = GetOEISLookupSequence("A000011", "Necklaces with color swap and turnover allowed", maxNumber, positional)
	case "all_ones_a000012":
		sequence, err = GetAllOnesA000012Sequence(maxNumber, positional)
	case "necklaces_color_swap_a000013":
		sequence, err = GetOEISLookupSequence("A000013", "Necklaces with color swap", maxNumber, positional)
	case "series_reduced_trees_a000014":
		sequence, err = GetOEISLookupSequence("A000014", "Series-reduced trees", maxNumber, positional)
	case "smallest_prime_power_a000015":
		sequence, err = GetSmallestPrimePowerA000015Sequence(maxNumber, positional)
	case "shift_register_sequences_a000016":
		sequence, err = GetOEISLookupSequence("A000016", "Shift register output sequences", maxNumber, positional)
	case "point_symmetric_queens_a000017":
		sequence, err = GetA000017Sequence(maxNumber, positional)
	case "form_x2_16y2_a000018":
		sequence, err = GetA000018Sequence(maxNumber, positional)
	case "primitive_permutation_groups_a000019":
		sequence, err = GetPrimitivePermutationGroupsA000019Sequence(maxNumber, positional)
	case "form_x2_12y2_a000021":
		sequence, err = GetA000021Sequence(maxNumber, positional)
	case "centered_hydrocarbons_a000022":
		sequence, err = GetA000022Sequence(maxNumber, positional)
	case "exp_minus_2x_a000023":
		sequence, err = GetA000023Sequence(maxNumber, positional)
	case "form_x2_10y2_a000024":
		sequence, err = GetA000024Sequence(maxNumber, positional)
	case "mock_theta_f_q_a000025":
		sequence, err = GetA000025Sequence(maxNumber, positional)
	case "mosaic_numbers_a000026":
		sequence, err = GetA000026Sequence(maxNumber, positional)
	case "positive_integers_a000027":
		sequence, err = GetA000027Sequence(maxNumber, positional)
	case "binary_weight_odd_a000028":
		sequence, err = GetA000028Sequence(maxNumber, positional)
	case "necklaces_turnover_a000029":
		sequence, err = GetA000029Sequence(maxNumber, positional)
	case "initial_digit_a000030":
		sequence, err = GetInitialDigitA000030Sequence(maxNumber, positional)
	case "necklaces_no_turnover_a000031":
		sequence, err = GetA000031Sequence(maxNumber, positional)
	case "lucas_numbers_a000032":
		sequence, err = GetA000032Sequence(maxNumber, positional)
	case "menage_hit_polynomials_a000033":
		sequence, err = GetA000033Sequence(maxNumber, positional)
	case "period_12_a000034":
		sequence, err = GetPeriod12A000034Sequence(maxNumber, positional)
	case "parity_a000035":
		sequence, err = GetParityA000035Sequence(maxNumber, positional)
	case "record_values_p_n_a000036":
		sequence, err = GetA000036Sequence(maxNumber, positional)
	case "nonsquares_a000037":
		sequence, err = GetA000037Sequence(maxNumber, positional)
	case "twice_characteristic_0_a000038":
		sequence, err = GetA000038Sequence(maxNumber, positional)
	case "mock_theta_f_q_coeff_a000039":
		sequence, err = GetA000039Sequence(maxNumber, positional)
	case "unary_a000042":
		sequence, err = GetUnaryA000042Sequence(maxNumber, positional)
	case "mersenne_exponents_a000043":
		sequence, err = GetOEISLookupSequence("A000043", "Mersenne exponents", maxNumber, positional)
	case "dying_rabbits_a000044":
		sequence, err = GetOEISLookupSequence("A000044", "Dying rabbits", maxNumber, positional)
	case "primitive_necklaces_complement_a000046":
		sequence, err = GetOEISLookupSequence("A000046", "Primitive necklaces (turnover, complement equivalent)", maxNumber, positional)
	case "form_x2_minus_2y2_a000047":
		sequence, err = GetOEISLookupSequence("A000047", "Numbers of form x^2 - 2y^2", maxNumber, positional)
	case "primitive_necklaces_color_swap_a000048":
		sequence, err = GetOEISLookupSequence("A000048", "Primitive necklaces with color swap", maxNumber, positional)
	case "form_3x2_4y2_a000049":
		sequence, err = GetOEISLookupSequence("A000049", "Numbers of form 3x^2 + 4y^2", maxNumber, positional)
	case "form_x2_y2_a000050":
		sequence, err = GetOEISLookupSequence("A000050", "Numbers of form x^2 + y^2", maxNumber, positional)
	case "2n_plus_1_a000051":
		sequence, err = GetA000051Sequence(maxNumber, positional)
	case "alphabetical_123_a000052":
		sequence, err = GetOEISLookupSequence("A000052", "Alphabetical 1,2,3... digit numbers", maxNumber, positional)
	case "nyc_subway_1_a000053":
		sequence, err = GetOEISLookupSequence("A000053", "NYC Subway 1 Train stops", maxNumber, positional)
	case "nyc_subway_a_a000054":
		sequence, err = GetOEISLookupSequence("A000054", "NYC Subway A Train stops", maxNumber, positional)
	case "unlabeled_trees_a000055":
		sequence, err = GetOEISLookupSequence("A000055", "Unlabeled trees", maxNumber, positional)
	case "sl2_zn_order_a000056":
		sequence, err = GetOEISLookupSequence("A000056", "Order of SL(2,Zn)", maxNumber, positional)
	case "fibonacci_dividing_primes_a000057":
		sequence, err = GetOEISLookupSequence("A000057", "Fibonacci dividing primes", maxNumber, positional)
	case "sylvester_a000058":
		sequence, err = GetA000058Sequence(maxNumber, positional)
	case "form_2k4_plus_1_primes_a000059":
		sequence, err = GetOEISLookupSequence("A000059", "k such that (2k)^4 + 1 is prime", maxNumber, positional)
	case "signed_trees_a000060":
		sequence, err = GetOEISLookupSequence("A000060", "Signed trees", maxNumber, positional)
	case "generalized_tangent_a000061":
		sequence, err = GetOEISLookupSequence("A000061", "Generalized tangent numbers", maxNumber, positional)
	case "beatty_e_minus_2_a000062":
		sequence, err = GetA000062Sequence(maxNumber, positional)
	case "symmetrical_dissections_a000063":
		sequence, err = GetOEISLookupSequence("A000063", "Symmetrical dissections of n-gon", maxNumber, positional)
	case "change_1_2_5_10_a000064":
		sequence, err = GetOEISLookupSequence("A000064", "Change for n cents (1,2,5,10)", maxNumber, positional)
	case "partitions_minus_1_a000065":
		sequence, err = GetPartitionsSequence(maxNumber, positional)
		if err == nil && sequence != nil {
			sequence.Name = "Partitions of n - 1 (A000065)"
			if positional {
				if sequence.Result != nil {
					sequence.Result = new(big.Int).Sub(sequence.Result, big.NewInt(1))
					sequence.Sequence = []*big.Int{sequence.Result}
				}
			} else {
				for i := range sequence.Sequence {
					sequence.Sequence[i] = new(big.Int).Sub(sequence.Sequence[i], big.NewInt(1))
				}
				if sequence.Result != nil {
					sequence.Result = new(big.Int).Sub(sequence.Result, big.NewInt(1))
				}
			}
		}
	case "trivalent_graph_girth_a000066":
		sequence, err = GetOEISLookupSequence("A000066", "Trivalent graph girth", maxNumber, positional)
	case "form_x2_2y2_a000067":
		sequence, err = GetOEISLookupSequence("A000067", "Numbers of form x^2 + 2y^2", maxNumber, positional)
	case "form_k4_plus_1_primes_a000068":
		sequence, err = GetOEISLookupSequence("A000068", "k such that k^4 + 1 is prime", maxNumber, positional)
	case "odious_a000069":
		sequence, err = GetA000069Sequence(maxNumber, positional)
	case "sum_partitions_a000070":
		sequence, err = GetOEISLookupSequence("A000070", "Sum of partition numbers", maxNumber, positional)
	case "fibonacci_minus_1_a000071":
		nPlus2 := new(big.Int).Add(maxNumber, big.NewInt(2))
		var seqObj *NumericSequence
		seqObj, err = GetFibonacciSequence(nPlus2, true)
		if err == nil && seqObj != nil {
			sequence = &NumericSequence{
				Name:   "Fibonacci(n+2) - 1 (A000071)",
				Number: new(big.Int).Set(maxNumber),
			}
			if positional {
				if seqObj.Result != nil {
					sequence.Result = new(big.Int).Sub(seqObj.Result, big.NewInt(1))
					sequence.Sequence = []*big.Int{new(big.Int).Set(sequence.Result)}
				}
			} else {
				// a(n) = Fib(n+2)-1 for n=0..limit-1
				limit := maxNumber.Int64()
				var seq []*big.Int
				for i := int64(0); i < limit; i++ {
					f, _ := GetFibonacciSequence(big.NewInt(i+2), true)
					val := new(big.Int).Sub(f.Result, big.NewInt(1))
					seq = append(seq, val)
				}
				sequence.Sequence = seq
				if len(seq) > 0 {
					sequence.Result = new(big.Int).Set(seq[len(seq)-1])
				}
			}
		}
	case "form_x2_4y2_a000072":
		sequence, err = GetOEISLookupSequence("A000072", "Numbers of form x^2 + 4y^2", maxNumber, positional)
	case "tribonacci_a000073":
		sequence, err = GetA000073Sequence(maxNumber, positional)
	case "odd_form_x2_y2_a000074":
		sequence, err = GetOEISLookupSequence("A000074", "Odd numbers of form x^2 + y^2", maxNumber, positional)
	case "form_2x2_3y2_a000075":
		sequence, err = GetOEISLookupSequence("A000075", "Numbers of form 2x^2 + 3y^2", maxNumber, positional)
	case "form_4x2_4xy_5y2_a000076":
		sequence, err = GetOEISLookupSequence("A000076", "Numbers of form 4x^2 + 4xy + 5y^2", maxNumber, positional)
	case "form_x2_6y2_a000077":
		sequence, err = GetOEISLookupSequence("A000077", "Numbers of form x^2 + 6y^2", maxNumber, positional)
	case "tetranacci_a000078":
		sequence, err = GetA000078Sequence(maxNumber, positional)
	case "powers_of_2_a000079":
		sequence, err = GetPowersOf2Sequence(maxNumber, positional)
	case "minimal_triangle_graphs_a000080":
		sequence, err = GetOEISLookupSequence("A000080", "Minimal triangle graphs", maxNumber, positional)
	case "rooted_unlabeled_trees_a000081":
		sequence, err = GetOEISLookupSequence("A000081", "Rooted unlabeled trees", maxNumber, positional)
	case "n2_phi_phi_a000082":
		sequence, err = GetOEISLookupSequence("A000082", "n^2 * Product(1 + 1/p)", maxNumber, positional)
	case "mixed_husimi_trees_a000083":
		sequence, err = GetOEISLookupSequence("A000083", "Mixed Husimi trees", maxNumber, positional)
	case "series_parallel_networks_a000084":
		sequence, err = GetOEISLookupSequence("A000084", "Series-parallel networks", maxNumber, positional)
	case "involutions_a000085":
		sequence, err = GetOEISLookupSequence("A000085", "Involutions (self-inverse permutations)", maxNumber, positional)
	case "solutions_x2_x_1_mod_n_a000086":
		sequence, err = GetOEISLookupSequence("A000086", "Solutions to x^2 - x + 1 == 0 (mod n)", maxNumber, positional)
	case "unrooted_maps_a000087":
		sequence, err = GetOEISLookupSequence("A000087", "Unrooted nonseparable planar maps", maxNumber, positional)
	case "simple_unlabeled_graphs_a000088":
		sequence, err = GetOEISLookupSequence("A000088", "Simple unlabeled graphs", maxNumber, positional)
	case "solutions_x2_1_mod_n_a000089":
		sequence, err = GetOEISLookupSequence("A000089", "Solutions to x^2 + 1 == 0 (mod n)", maxNumber, positional)
	case "exp_minus_x3_3_a000090":
		sequence, err = GetOEISLookupSequence("A000090", "Expansion of exp(-x^3/3)/(1-x)", maxNumber, positional)
	case "multiplicative_a000091":
		sequence, err = GetOEISLookupSequence("A000091", "A000091", maxNumber, positional)
	case "record_values_p_n_3d_a000092":
		sequence, err = GetOEISLookupSequence("A000092", "Record values of |P(n)| in 3D", maxNumber, positional)
	case "floor_n_1_5_a000093":
		sequence, err = GetA000093Sequence(maxNumber, positional)
	case "trees_diameter_4_a000094":
		sequence, err = GetOEISLookupSequence("A000094", "Trees of diameter 4", maxNumber, positional)
	case "fixed_points_gamma0_n_a000095":
		sequence, err = GetOEISLookupSequence("A000095", "Fixed points of Gamma_0(n)", maxNumber, positional)
	case "n_n_plus_3_2_a000096":
		sequence, err = GetA000096Sequence(maxNumber, positional)
	case "primitive_polynomials_a000020":
		sequence, err = GetOEISLookupSequence("A000020", "Primitive polynomials", maxNumber, positional)
	case "partitions_2kinds_1_2_a000097":
		sequence, err = GetOEISLookupSequence("A000097", "Partitions (2 kinds of 1, 2)", maxNumber, positional)
	case "partitions_2kinds_1_2_3_a000098":
		sequence, err = GetOEISLookupSequence("A000098", "Partitions (2 kinds of 1, 2, 3)", maxNumber, positional)
	case "record_values_p_n_2d_a000099":
		sequence, err = GetOEISLookupSequence("A000099", "Record values of |P(n)| in 2D", maxNumber, positional)
	case "sqrt_prime_a000006":
		sequence, err = GetSqrtPrimeA000006Sequence(maxNumber, positional)
	case "binary_quadratic_forms_a000003":
		sequence, err = GetBinaryQuadraticFormsA000003Sequence(maxNumber, positional)
	case "hamming_weight":
		sequence, err = GetHammingWeightSequence(maxNumber, positional)
	case "central_polygonal":
		sequence, err = GetCentralPolygonalNumbersSequence(maxNumber, positional)
	case "squares":
		sequence, err = GetSquaresSequence(maxNumber, positional)
	case "cubes":
		sequence, err = GetCubesA000578Sequence(maxNumber, positional)
	case "natural":
		sequence, err = GetNaturalSequence(maxNumber, positional)
	case "parity":
		sequence, err = GetParitySequence(maxNumber, positional)
	case "prime":
		sequence, err = GetPrimesA000040Sequence(maxNumber, positional)
	case "primes_a000040":
		sequence, err = GetPrimesA000040Sequence(maxNumber, positional)
	case "emirp":
		sequence, err = GetEmirpSequence(maxNumber, positional)
	case "semiprime":
		sequence, err = GetSemiPrimeSequence(maxNumber, positional)
	case "circular_prime":
		sequence, err = GetCircularPrimeSequence(maxNumber, positional)
	case "fibonacci_prime":
		sequence, err = GetFibonacciPrimeSequence(maxNumber, positional)
	case "cake":
		sequence, err = GetCakeSequence(maxNumber, positional)
	case "bell":
		sequence, err = GetBellSequence(maxNumber, positional)
	case "catalan":
		sequence, err = GetCatalanSequence(maxNumber, positional)
	case "totient":
		sequence, err = GetEulerTotientA000010Sequence(maxNumber, positional)
	case "totient_prime":
		sequence, err = GetTotientPrimeSequence(maxNumber)
	case "fibonacci":
		sequence, err = GetFibonacciSequence(maxNumber, positional)
	case "pell":
		sequence, err = GetPellSequence(maxNumber, positional)
	case "zekendorf":
		sequence, err = GetZekendorfRepresentationSequence(maxNumber, positional)
	case "lucas":
		sequence, err = GenerateLucas(maxNumber, positional)
	case "nn":
		sequence, err = GetNtoNSequence(maxNumber, positional)
	case "schroeder_fourth":
		sequence, err = GetSchroederFourthSequence(maxNumber, positional)
	case "partitions_distinct":
		sequence, err = GetPartitionsDistinctSequence(maxNumber, positional)
	case "partitions":
		sequence, err = GetPartitionsSequence(maxNumber, positional)
	case "partitions_into_2_squares":
		sequence, err = GetPartitionsInto2SquaresSequence(maxNumber, positional)
	case "plane_partitions":
		sequence, err = GetPlanePartitionsSequence(maxNumber, positional)
	case "tangent":
		sequence, err = GetTangentNumbersSequence(maxNumber, positional)
	case "kendall_mann":
		sequence, err = GetKendallMannSequence(maxNumber, positional)
	case "pentagonal":
		sequence, err = GetPentagonalSequence(maxNumber, positional)
	case "square_pyramidal":
		sequence, err = GetSquarePyramidalSequence(maxNumber, positional)
	case "euler":
		sequence, err = GetEulerNumbersSequence(maxNumber, positional)
	case "euler_zigzag":
		sequence, err = GetEulerZigzagSequence(maxNumber, positional)
	case "perfect":
		sequence, err = GetPerfectNumbersSequence(maxNumber, positional)
	case "groups_order_n":
		sequence, err = GetGroupsOrderNSequence(maxNumber, positional)
	case "modular_j":
		sequence, err = GetModularJCoefficientsSequence(maxNumber, positional)
	case "ramanujan_tau":
		sequence, err = GetRamanujanTauSequence(maxNumber, positional)
	case "fourth_powers":
		sequence, err = GetFourthPowersSequence(maxNumber, positional)
	case "tetrahedral":
		sequence, err = GetTetrahedralSequence(maxNumber, positional)
	case "triangular":
		sequence, err = GetTriangularSequence(maxNumber, positional)
	case "sum_divisors":
		sequence, err = GetSumDivisorsSequence(maxNumber, positional)
	case "sum_odd_divisors":
		sequence, err = GetSumOddDivisorsSequence(maxNumber, positional)
	case "alkanes":
		sequence, err = GetAlkanesSequence(maxNumber, positional)
	case "abelian_groups_order_n":
		sequence, err = GetAbelianGroupsOrderNSequence(maxNumber, positional)
	case "threshold_functions":
		sequence, err = GetThresholdFunctionsSequence(maxNumber, positional)
	case "fubini":
		sequence, err = GetFubiniSequence(maxNumber, positional)
	case "kolakoski":
		sequence, err = GetKolakoskiSequence(maxNumber, positional)
	case "zero":
		sequence, err = GetZeroSequence(maxNumber, positional)
	case "zero_characteristic_a000007":
		sequence, err = GetZeroCharacteristicSequence(maxNumber, positional)
	case "divisor_count_a000005":
		sequence, err = GetDivisorCountA000005Sequence(maxNumber, positional)
	case "divisor_count":
		sequence, err = GetDivisorCountA000005Sequence(maxNumber, positional)
	case "change_1_2_5_10_a000008":
		sequence, err = GetA000008Sequence(maxNumber, positional)
	case "collatz":
		sequence, err = GetCollatzSequence(maxNumber.Int64(), positional)
	case "powers_of_2":
		sequence, err = GetPowersOf2Sequence(maxNumber, positional)
	case "powers_of_4":
		sequence, err = GetPowersOf4Sequence(maxNumber, positional)
	case "powers_of_3":
		sequence, err = GetPowersOf3Sequence(maxNumber, positional)
	case "odious_numbers":
		sequence, err = GetOdiousNumbersSequence(maxNumber, positional)
	case "subfactorial":
		sequence, err = GetSubfactorialSequence(maxNumber, positional)
	case "binary_partitions":
		sequence, err = GetBinaryPartitionsSequence(maxNumber, positional)
	case "binary_rooted_trees":
		sequence, err = GetBinaryRootedTreesA002572Sequence(maxNumber, positional)
	case "sqrt3_convergents":
		sequence, err = GetSqrt3ConvergentsA002531Sequence(maxNumber, positional)
	case "sqrt3_convergents_denominators":
		sequence, err = GetSqrt3ConvergentsDenominatorsSequence(maxNumber, positional)
	case "factorial":
		sequence, err = GetFactorialSequence(maxNumber, positional)
	case "planted_3_trees":
		sequence, err = GetPlanted3TreesSequence(maxNumber, positional)
	case "rooted_unlabeled_trees":
		sequence, err = GetRootedUnlabeledTreesSequence(maxNumber, positional)
	case "unlabeled_trees":
		sequence, err = GetUnlabeledTreesSequence(maxNumber, positional)
	case "unlabeled_digraphs":
		sequence, err = GetUnlabeledDigraphsSequence(maxNumber, positional)
	case "unlabeled_graphs":
		sequence, err = GetUnlabeledGraphsSequence(maxNumber, positional)
	case "connected_planar_graphs":
		sequence, err = GetConnectedPlanarGraphsSequence(maxNumber, positional)
	case "unlabeled_posets":
		sequence, err = GetUnlabeledPosetsSequence(maxNumber, positional)
	case "bicolorable_necklaces":
		sequence, err = GetBicolorableNecklacesSequence(maxNumber, positional)
	case "simplicial_polyhedra":
		sequence, err = GetSimplicialPolyhedraSequence(maxNumber, positional)
	case "labeled_rooted_trees":
		sequence, err = GetLabeledRootedTreesSequence(maxNumber, positional)
	case "labeled_trees":
		sequence, err = GetLabeledTreesSequence(maxNumber, positional)
	case "sets_of_lists":
		sequence, err = GetSetsOfListsSequence(maxNumber, positional)
	case "free_polyominoes":
		sequence, err = GetFreePolyominoesSequence(maxNumber, positional)
	case "self_inverse_permutations":
		sequence, err = GetSelfInversePermutationsSequence(maxNumber, positional)
	case "sylvester":
		sequence, err = GetSylvesterSequence(maxNumber, positional)
	case "theta_series_square_lattice":
		sequence, err = GetThetaSeriesSquareLatticeSequence(maxNumber, positional)
	case "theta_series_d4_lattice":
		sequence, err = GetThetaSeriesD4LatticeSequence(maxNumber, positional)
	case "mersenne_numbers":
		sequence, err = GetMersenneNumbersSequence(maxNumber, positional)
	case "mersenne_prime_exponents":
		sequence, err = GetMersennePrimeExponentsSequence(maxNumber, positional)
	case "radon_hurwitz":
		sequence, err = GetRadonSequence(maxNumber, positional)
	case "lcm_1_to_n":
		sequence, err = GetLCM1ToNSequence(maxNumber, positional)
	case "loeschian":
		sequence, err = GetLoeschianSequence(maxNumber, positional)
	case "composites":
		sequence, err = GetCompositesSequence(maxNumber, positional)
	case "quarter_squares":
		sequence, err = GetQuarterSquaresSequence(maxNumber, positional)
	case "ways_two_squares":
		sequence, err = GetWaysTwoSquaresSequence(maxNumber, positional)
	case "stern":
		sequence, err = GetSternSequence(maxNumber, positional)
	default:
		fmt.Printf("Unknown sequence type: %s\n", sequenceType)
		err = fmt.Errorf("unknown sequence type: %s", sequenceType)
	}

	return sequence, err
}
