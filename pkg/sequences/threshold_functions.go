package sequences

import (
	"fmt"
	"math/big"
)

// A000609: Number of threshold functions of n or fewer variables.
// URL: https://oeis.org/A000609
// The known terms of this sequence are limited (n=0 to 9).
// a(0)=2, a(1)=4, a(2)=14, a(3)=104, a(4)=1882, a(5)=94572, a(6)=15028134, a(7)=8378070864, a(8)=17561539552946, a(9)=144130531453121108.

var thresholdFunctionsTable = []string{
	"2",
	"4",
	"14",
	"104",
	"1882",
	"94572",
	"15028134",
	"8378070864",
	"17561539552946",
	"144130531453121108",
}

// GetThresholdFunctionsSequence returns the sequence of threshold functions of n or fewer variables (A000609).
func GetThresholdFunctionsSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetThresholdFunctionAtPosition(maxNumber)
	}
	return GenerateThresholdFunctionsSequence(maxNumber)
}

// GenerateThresholdFunctionsSequence generates the sequence up to maxNumber (a(0), ..., a(min(maxNumber, 9))).
func GenerateThresholdFunctionsSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	limit := int(maxNumber.Int64())
	if maxNumber.IsInt64() && limit >= len(thresholdFunctionsTable) {
		limit = len(thresholdFunctionsTable) - 1
	} else if !maxNumber.IsInt64() {
		limit = len(thresholdFunctionsTable) - 1
	}

	sequence := make([]*big.Int, limit+1)
	for i := 0; i <= limit; i++ {
		val := new(big.Int)
		val.SetString(thresholdFunctionsTable[i], 10)
		sequence[i] = val
	}

	return &NumericSequence{
		Name:     "Threshold functions (A000609)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[len(sequence)-1],
	}, nil
}

// GetThresholdFunctionAtPosition returns the n-th term of A000609 (0 <= n <= 9).
func GetThresholdFunctionAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	if !n.IsInt64() || n.Int64() >= int64(len(thresholdFunctionsTable)) {
		return nil, fmt.Errorf("position %s exceeds known terms of sequence A000609", n.String())
	}

	index := int(n.Int64())
	result := new(big.Int)
	result.SetString(thresholdFunctionsTable[index], 10)

	return &NumericSequence{
		Name:     "Threshold functions (A000609)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}
