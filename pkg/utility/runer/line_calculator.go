package runer

import (
	"math/big"
	"quantix-math/pkg/utility/runelib"
	"strings"
)

type LineCalculator struct {
	charRepo *runelib.CharacterRepo
}

func NewLineCalculator(charRepo *runelib.CharacterRepo) *LineCalculator {
	return &LineCalculator{
		charRepo: charRepo,
	}
}

type LineCalculations struct {
	Lines []LineCalculation `json:"lines"`
}

type LineCalculation struct {
	LineNumber int    `json:"line_number"`
	LineTotal  string `json:"line_total"`
}

func (lc *LineCalculator) CalculateLines(input string) LineCalculations {
	lineCalculations := LineCalculations{
		Lines: make([]LineCalculation, 0),
	}

	inputArray := strings.Split(input, "")
	line := make([]string, 0)

	// Are the lines in the input.
	lines := make([]string, 0)

	// This is to build out our lines.
	for _, char := range inputArray {
		if lc.charRepo.IsRune(char, true) {
			line = append(line, char)
		}

		if lc.charRepo.IsLineSeperator(char) {
			lines = append(lines, strings.Join(line, ""))
			line = make([]string, 0)
		}
	}

	lines = append(lines, strings.Join(line, ""))
	line = make([]string, 0)

	// This is to calculate the value of the line.
	for lineCounter, tempLine := range lines {
		wordInts := make([]int64, 0)
		wordInt := int64(0)
		lineTotal := big.NewInt(1)
		lineArray := strings.Split(tempLine, "")

		for _, char := range lineArray {
			if lc.charRepo.IsRuneSeparator(char) {
				if wordInt > 0 {
					wordInts = append(wordInts, wordInt)
					wordInt = 0
				}
			} else {
				wordInt = wordInt + lc.charRepo.GetValueFromRune64(char)
			}
		}

		if wordInt > 0 {
			wordInts = append(wordInts, wordInt)
		}

		for _, tempWordInt := range wordInts {
			lineTotal.Mul(lineTotal, big.NewInt(tempWordInt))
		}

		lineCalculation := LineCalculation{
			LineNumber: lineCounter + 1,
			LineTotal:  lineTotal.String(),
		}

		if lineTotal.Cmp(big.NewInt(2)) > 0 {
			lineCalculations.Lines = append(lineCalculations.Lines, lineCalculation)
		}

		lineTotal = big.NewInt(1)
	}

	return lineCalculations
}
