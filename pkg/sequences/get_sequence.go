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
	case "central_polygonal":
		sequence, err = GetCentralPolygonalNumbersSequence(maxNumber, positional)
	case "cubes":
		sequence, err = GetCubesA000578Sequence(maxNumber, positional)
	case "natural":
		sequence, err = GetNaturalSequence(maxNumber, positional)
	case "prime":
		sequence, err = GetPrimeSequence(maxNumber, positional)
	case "fibonacci_prime":
		sequence, err = GetFibonacciPrimeSequence(maxNumber, positional)
	case "cake":
		sequence, err = GetCakeSequence(maxNumber, positional)
	case "catalan":
		sequence, err = GetCatalanSequence(maxNumber, positional)
	case "totient":
		sequence, err = GetTotientSequence(maxNumber)
	case "totient_prime":
		sequence, err = GetTotientPrimeSequence(maxNumber)
	case "fibonacci":
		sequence, err = GetFibonacciSequence(maxNumber, positional)
	case "zekendorf":
		sequence, err = GetZekendorfRepresentationSequence(maxNumber, positional)
	case "lucas":
		sequence, err = GenerateLucas(maxNumber, positional)
	case "euler":
		sequence, err = GetEulerNumbersSequence(maxNumber, positional)
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
	default:
		fmt.Printf("Unknown sequence type: %s\n", sequenceType)
		err = fmt.Errorf("unknown sequence type: %s", sequenceType)
	}

	return sequence, err
}
