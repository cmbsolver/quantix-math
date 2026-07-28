package sequences

import (
	"fmt"
	"math/big"
)

// Erroneous version of A032522. (OEIS A000017)
// URL: https://oeis.org/A000017
// Description: Number of point symmetric solutions to non-attacking queens problem on n X n board.
// Point symmetric means the solution is invariant under a 180-degree rotation.
// The sequence defined in A000017 is considered erroneous and was replaced by A032522.
// However, per instructions, we implement the point symmetric N-queens calculation.

// GetA000017Sequence returns the A000017 sequence.
func GetA000017Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA000017AtPosition(maxNumber)
	}
	return GenerateA000017Sequence(maxNumber)
}

// GenerateA000017Sequence generates the A000017 sequence up to maxNumber terms.
func GenerateA000017Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Sign() < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	n := int(maxNumber.Int64())
	sequence := make([]*big.Int, n)

	for i := 1; i <= n; i++ {
		val := calculatePointSymmetricQueens(i)
		sequence[i-1] = big.NewInt(int64(val))
	}

	result := big.NewInt(0)
	if n > 0 {
		result = sequence[n-1]
	}

	return &NumericSequence{
		Name:     "Point symmetric queens (A000017)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   result,
	}, nil
}

// GetA000017AtPosition returns the n-th term of A000017.
func GetA000017AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Sign() <= 0 {
		return nil, fmt.Errorf("position must be positive")
	}

	val := calculatePointSymmetricQueens(int(n.Int64()))

	return &NumericSequence{
		Name:     "Point symmetric queens (A000017)",
		Number:   n,
		Sequence: []*big.Int{big.NewInt(int64(val))},
		Result:   big.NewInt(int64(val)),
	}, nil
}

// calculatePointSymmetricQueens calculates the number of point symmetric solutions to the n-queens problem.
func calculatePointSymmetricQueens(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	// For n=2, 3 no solutions exist
	if n == 2 || n == 3 {
		return 0
	}

	count := 0
	board := make([]int, n)
	for i := range board {
		board[i] = -1
	}

	// For point symmetry, a queen at (r, c) implies a queen at (n-1-r, n-1-c).
	// We only need to place queens in the first n/2 rows.
	// If n is odd, the middle queen must be at the center ((n-1)/2, (n-1)/2).

	if n%2 == 1 {
		center := (n - 1) / 2
		board[center] = center
		solvePointSymmetric(board, 0, n, &count)
	} else {
		solvePointSymmetric(board, 0, n, &count)
	}

	return count
}

func solvePointSymmetric(board []int, row int, n int, count *int) {
	half := n / 2
	if row == half {
		*count++
		return
	}

	// Skip center row if n is odd (already placed)
	if n%2 == 1 && row == (n-1)/2 {
		solvePointSymmetric(board, row+1, n, count)
		return
	}

	// Try placing a queen in row 'row' at column 'col'
	for col := 0; col < n; col++ {
		symRow := n - 1 - row
		symCol := n - 1 - col

		// Symmetric queen would be in the same row?
		// Only happens if n is odd and row is center, but we handled that.

		// Two queens at (row, col) and (symRow, symCol)
		if isSafePointSymmetric(board, row, col, n) {
			board[row] = col
			board[symRow] = symCol
			solvePointSymmetric(board, row+1, n, count)
			board[row] = -1
			board[symRow] = -1
		}
	}
}

func isSafePointSymmetric(board []int, row, col, n int) bool {
	symRow := n - 1 - row
	symCol := n - 1 - col

	// 1. Check if (row, col) and (symRow, symCol) are the same
	// In n-queens, we need one queen per row.
	// If they are in the same row, it's only possible if row == symRow,
	// which means n is odd and row is the middle row.
	// But we also need one queen per column, so col == symCol.
	if row == symRow && col != symCol {
		return false // Should not happen with our row-by-row approach
	}

	// 2. Check if the two queens attack each other
	if col == symCol {
		return false // Same column
	}
	if row-col == symRow-symCol {
		return false // Same diagonal
	}
	if row+col == symRow+symCol {
		return false // Same anti-diagonal
	}

	// 3. Check against previously placed queens
	// We only need to check against rows < row and their symmetric counterparts.
	for i := 0; i < row; i++ {
		c := board[i]
		if c == -1 {
			continue
		}
		sc := n - 1 - c
		sr := n - 1 - i

		// Current queen at (row, col)
		if c == col || sc == col ||
			i-c == row-col || sr-sc == row-col ||
			i+c == row+col || sr+sc == row+col {
			return false
		}

		// Symmetric queen at (symRow, symCol)
		// Actually, if (row, col) is safe against (i, c) and (sr, sc),
		// then by symmetry (symRow, symCol) is also safe.
	}

	// If n is odd, check against the center queen
	if n%2 == 1 {
		center := (n - 1) / 2
		if row != center {
			if col == center ||
				row-col == 0 || // center is (mid, mid), so row-col == mid-mid == 0
				row+col == n-1 { // center is (mid, mid), so row+col == mid+mid == n-1
				return false
			}
		}
	}

	return true
}

// IsPointSymmetricQueensA000017 checks if a number exists in the A000017 sequence.
func IsPointSymmetricQueensA000017(n *big.Int) (bool, string, error) {
	if n.Sign() < 0 {
		return false, "", nil
	}

	// Since this sequence grows rapidly, we can check up to a reasonable n.
	for i := 1; i <= 20; i++ {
		val := calculatePointSymmetricQueens(i)
		if big.NewInt(int64(val)).Cmp(n) == 0 {
			return true, fmt.Sprintf("%d", i), nil
		}
		if int64(val) > n.Int64() && n.IsInt64() {
			break
		}
	}

	return false, "", nil
}
