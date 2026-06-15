package sequences

import (
	"fmt"
	"math/big"
	"slices"
	"sort"
)

// Free Polyominoes (OEIS A000105)
// URL: https://oeis.org/A000105
// a(n) is the number of free polyominoes (or square animals) with n cells.
// Free polyominoes are considered identical if they are congruent by rotation or reflection.

// GetFreePolyominoesSequence returns the number of free polyominoes with n cells (OEIS A000105).
func GetFreePolyominoesSequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetFreePolyominoesAtPosition(maxNumber)
	}
	return GenerateFreePolyominoesSequence(maxNumber)
}

// GenerateFreePolyominoesSequence generates the A000105 sequence up to maxNumber.
// It returns a(0), a(1), ..., a(maxNumber).
func GenerateFreePolyominoesSequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number must be at least 0")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n+1)

	// A000105: 1, 1, 1, 2, 5, 12, 35, 108, 369, 1285, 4655, 17073
	for i := 0; i <= n; i++ {
		val := calculateA000105(i)
		sequence[i] = big.NewInt(int64(val))
	}

	return &NumericSequence{
		Name:     "Free Polyominoes (A000105)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   sequence[n],
	}, nil
}

// GetFreePolyominoesAtPosition returns the n-th term of A000105 (n >= 0).
func GetFreePolyominoesAtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("position must be at least 0")
	}

	pos := int(n.Int64())
	result := big.NewInt(int64(calculateA000105(pos)))

	return &NumericSequence{
		Name:     "Free Polyominoes (A000105)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

type point struct {
	x, y int
}

type polyomino []point

func (p polyomino) Len() int { return len(p) }
func (p polyomino) Less(i, j int) bool {
	if p[i].x != p[j].x {
		return p[i].x < p[j].x
	}
	return p[i].y < p[j].y
}
func (p polyomino) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// canonical returns the canonical representation of a polyomino under D8 symmetry.
func canonical(p polyomino) string {
	if len(p) == 0 {
		return ""
	}

	variants := make([][]point, 8)
	for i := range 8 {
		variants[i] = make([]point, len(p))
	}

	for i, pt := range p {
		x, y := pt.x, pt.y
		variants[0][i] = point{x, y}   // Identity
		variants[1][i] = point{-y, x}  // 90 deg
		variants[2][i] = point{-x, -y} // 180 deg
		variants[3][i] = point{y, -x}  // 270 deg
		variants[4][i] = point{x, -y}  // Reflection x
		variants[5][i] = point{-x, y}  // Reflection y
		variants[6][i] = point{y, x}   // Reflection y=x
		variants[7][i] = point{-y, -x} // Reflection y=-x
	}

	results := make([]string, 8)
	for i := range 8 {
		// Normalize to min x, min y = 0
		minX, minY := variants[i][0].x, variants[i][0].y
		for _, pt := range variants[i] {
			if pt.x < minX {
				minX = pt.x
			}
			if pt.y < minY {
				minY = pt.y
			}
		}
		for j := range variants[i] {
			variants[i][j].x -= minX
			variants[i][j].y -= minY
		}
		sort.Sort(polyomino(variants[i]))

		s := ""
		for _, pt := range variants[i] {
			s += fmt.Sprintf("(%d,%d)", pt.x, pt.y)
		}
		results[i] = s
	}

	return slices.Min(results)
}

var a000105Cache = map[int]int{
	0: 1,
}

func calculateA000105(n int) int {
	if n == 0 {
		return 1
	}
	if val, ok := a000105Cache[n]; ok {
		return val
	}

	// For small n, we can use recursive generation.
	// For larger n, this will be slow, but for a typical exercise it should suffice.
	// A000105: 1, 1, 1, 2, 5, 12, 35, 108, 369...

	currentLevel := map[string]polyomino{
		"(0,0)": []point{{0, 0}},
	}

	for i := 1; i < n; i++ {
		nextLevel := make(map[string]polyomino)
		for _, p := range currentLevel {
			// Find all neighbors
			neighbors := make(map[point]bool)
			for _, pt := range p {
				neighbors[point{pt.x + 1, pt.y}] = true
				neighbors[point{pt.x - 1, pt.y}] = true
				neighbors[point{pt.x, pt.y + 1}] = true
				neighbors[point{pt.x, pt.y - 1}] = true
			}
			// Remove points already in p
			for _, pt := range p {
				delete(neighbors, pt)
			}

			for neighbor := range neighbors {
				newP := make(polyomino, len(p)+1)
				copy(newP, p)
				newP[len(p)] = neighbor
				nextLevel[canonical(newP)] = newP
			}
		}
		currentLevel = nextLevel
	}

	a000105Cache[n] = len(currentLevel)
	return len(currentLevel)
}
