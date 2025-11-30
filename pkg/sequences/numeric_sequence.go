package sequences

import (
	"math/big"
)

// NumericSequence represents a sequence of numbers.
type NumericSequence struct {
	Name     string
	Number   *big.Int
	Sequence []*big.Int
	Result   *big.Int
}

// NumericSequence64 represents a sequence of numbers.
type NumericSequence64 struct {
	Name     string
	Number   int64
	Sequence []int64
	Result   int64
}
