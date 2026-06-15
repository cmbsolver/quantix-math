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
		"hamming_weight", "natural", "parity", "partitions_distinct", "partitions", "kolakoski", "zero", "divisor_count", "sum_divisors", "ways_to_make_change", "groups_order_n", "nn", "squares", "cubes", "prime", "primes_a000040", "emirp", "semiprime", "circular_prime", "fibonacci", "lucas", "fourth_powers",
		"triangular", "central_polygonal", "cake", "bell", "catalan", "totient", "totient_prime",
		"tetrahedral",
		"fibonacci_prime", "zekendorf", "ramanujan_tau",
		"sum_odd_divisors", "alkanes", "abelian_groups_order_n",
		"threshold_functions", "fubini", "schroeder_fourth",
		"partitions_into_2_squares",
		"powers_of_2",
		"powers_of_4",
		"subfactorial", "self_inverse_permutations",
		"binary_partitions",
		"planted_3_trees",
		"unlabeled_graphs", "connected_planar_graphs",
		"kendall_mann",
		"bicolorable_necklaces",
		"factorial",
		"labeled_rooted_trees",
		"plane_partitions",
		"tangent",
		"free_polyominoes",
		"theta_series_square_lattice",
		"theta_series_d4_lattice",
		"euler_zigzag",
		"sylvester",
		"mersenne_numbers", "mersenne_prime_exponents",
		"sets_of_lists",
		"unlabeled_trees",
		"zero_characteristic", "collatz",
		"euler", "perfect", "modular_j", "square_pyramidal", "pentagonal", "radon_hurwitz", "lcm_1_to_n", "loeschian", "composites", "stern",
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
	case "partitions_into_2_squares":
		return "Partitions into 2 squares (A000161)"
	case "divisor_count":
		return "Number of Divisors (A000005)"
	case "sum_divisors":
		return "Sum of Divisors (A000203)"
	case "ways_to_make_change":
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
		"factorial", "kendall_mann", "bicolorable_necklaces", "binary_partitions", "partitions",
		"fibonacci_prime", "zekendorf", "groups_order_n", "ramanujan_tau",
		"sum_divisors", "sum_odd_divisors", "alkanes", "abelian_groups_order_n",
		"theta_series_d4_lattice",
		"ways_two_squares",
		"threshold_functions", "fubini", "schroeder_fourth",
		"partitions_into_2_squares",
		"powers_of_2",
		"powers_of_4",
		"tangent", "self_inverse_permutations",
		"euler_zigzag",
		"sylvester",
		"mersenne_numbers", "mersenne_prime_exponents", "rooted_unlabeled_trees", "unlabeled_trees", "unlabeled_digraphs", "unlabeled_graphs", "connected_planar_graphs",
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
