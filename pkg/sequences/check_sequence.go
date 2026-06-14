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
		"natural", "partitions_distinct", "kolakoski", "zero", "divisor_count", "ways_to_make_change", "groups_order_n", "nn", "squares", "cubes", "prime", "emirp", "semiprime", "circular_prime", "fibonacci", "lucas", "fourth_powers",
		"central_polygonal", "cake", "catalan", "totient", "totient_prime",
		"tetrahedral",
		"fibonacci_prime", "zekendorf", "ramanujan_tau",
		"sum_odd_divisors", "alkanes", "abelian_groups_order_n",
		"threshold_functions", "fubini", "schroeder_fourth",
		"powers_of_4",
		"zero_characteristic", "collatz",
		"euler", "perfect", "modular_j", "square_pyramidal", "pentagonal",
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
	case "natural":
		return "Natural"
	case "partitions_distinct":
		return "Partitions into distinct parts (A000009)"
	case "squares":
		return "Squares (A000290)"
	case "cubes":
		return "Cubes (A000578)"
	case "prime":
		return "Prime"
	case "emirp":
		return "Emirp"
	case "semiprime":
		return "Semi-prime"
	case "circular_prime":
		return "Circular Prime"
	case "fibonacci":
		return "Fibonacci"
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
	case "modular_j":
		return "Modular function j (A000521)"
	case "central_polygonal":
		return "Central Polygonal Numbers"
	case "cake":
		return "Cake"
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
	case "groups_order_n":
		return "Number of Groups (A000001)"
	case "nn":
		return "n^n (A000312)"
	case "powers_of_4":
		return "Powers of 4 (A000302)"
	case "tetrahedral":
		return "Tetrahedral numbers (A000292)"
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
	case "zero":
		return "Zero Sequence (A000004)"
	case "zero_characteristic":
		return "Zero Characteristic (A000007)"
	case "divisor_count":
		return "Number of Divisors (A000005)"
	case "ways_to_make_change":
		return "Ways to Make Change (A000008)"
	case "collatz":
		return "Collatz"
	default:
		return st
	}
}

func checkExistence(n *big.Int, st string) (bool, string, error) {
	switch st {
	case "natural":
		if n.Sign() >= 0 {
			return true, n.String(), nil
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
	case "prime":
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
	case "zero":
		if n.Cmp(big.NewInt(0)) == 0 {
			return true, "any", nil
		}
	case "zero_characteristic":
		if n.Cmp(big.NewInt(0)) == 0 {
			return true, "any n > 0", nil
		}
		if n.Cmp(big.NewInt(1)) == 0 {
			return true, "0", nil
		}
	case "cubes", "central_polygonal", "cake", "catalan", "totient", "totient_prime",
		"fibonacci_prime", "zekendorf", "groups_order_n", "ramanujan_tau",
		"sum_odd_divisors", "alkanes", "abelian_groups_order_n",
		"threshold_functions", "fubini", "kolakoski", "ways_to_make_change", "collatz", "divisor_count", "lucas", "square_pyramidal", "pentagonal", "partitions_distinct", "nn", "schroeder_fourth", "powers_of_4", "tetrahedral":
		// Fallback: generate sequence and check (with reasonable limit)
		seq, err := GetSequence(n.String(), st, false)
		if err == nil && seq != nil {
			for i, val := range seq.Sequence {
				if val.Cmp(n) == 0 {
					return true, fmt.Sprintf("%d", i), nil
				}
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
