package sequences

import (
	"fmt"
	"math/big"
)

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
		sequence, err = GetOEISLookupSequence("A000028", "Binary weight of exponents is odd", maxNumber, positional)
	case "necklaces_turnover_a000029":
		sequence, err = GetOEISLookupSequence("A000029", "Necklaces with turnover allowed (bracelets)", maxNumber, positional)
	case "initial_digit_a000030":
		sequence, err = GetInitialDigitA000030Sequence(maxNumber, positional)
	case "necklaces_no_turnover_a000031":
		sequence, err = GetOEISLookupSequence("A000031", "Necklaces with no turnover allowed", maxNumber, positional)
	case "menage_hit_polynomials_a000033":
		sequence, err = GetOEISLookupSequence("A000033", "Ménage hit polynomials", maxNumber, positional)
	case "period_12_a000034":
		sequence, err = GetPeriod12A000034Sequence(maxNumber, positional)
	case "parity_a000035":
		sequence, err = GetParityA000035Sequence(maxNumber, positional)
	case "record_values_p_n_a000036":
		sequence, err = GetOEISLookupSequence("A000036", "Record values of |P(n)|", maxNumber, positional)
	case "nonsquares_a000037":
		sequence, err = GetOEISLookupSequence("A000037", "Nonsquares", maxNumber, positional)
	case "twice_characteristic_0_a000038":
		sequence, err = GetOEISLookupSequence("A000038", "Twice characteristic function of {0}", maxNumber, positional)
	case "mock_theta_f_q_coeff_a000039":
		sequence, err = GetOEISLookupSequence("A000039", "Coefficients of mock theta function f(q)", maxNumber, positional)
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
