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
		sequence, err = GetTotientSequence(maxNumber)
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
	case "zero_characteristic":
		sequence, err = GetZeroCharacteristicSequence(maxNumber, positional)
	case "divisor_count":
		sequence, err = GetDivisorCountSequence(maxNumber, positional)
	case "ways_to_make_change":
		sequence, err = GetWaysToMakeChangeSequence(maxNumber, positional)
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
	case "factorial":
		sequence, err = GetFactorialSequence(maxNumber, positional)
	case "rooted_unlabeled_trees":
		sequence, err = GetRootedUnlabeledTreesSequence(maxNumber, positional)
	case "unlabeled_trees":
		sequence, err = GetUnlabeledTreesSequence(maxNumber, positional)
	case "unlabeled_digraphs":
		sequence, err = GetUnlabeledDigraphsSequence(maxNumber, positional)
	case "unlabeled_graphs":
		sequence, err = GetUnlabeledGraphsSequence(maxNumber, positional)
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
	case "mersenne_numbers":
		sequence, err = GetMersenneNumbersSequence(maxNumber, positional)
	case "mersenne_prime_exponents":
		sequence, err = GetMersennePrimeExponentsSequence(maxNumber, positional)
	default:
		fmt.Printf("Unknown sequence type: %s\n", sequenceType)
		err = fmt.Errorf("unknown sequence type: %s", sequenceType)
	}

	return sequence, err
}
