package runer

import (
	"quantix-math/pkg/utility/runelib"
	"strings"
)

// TransposeLatinToRune transposes a Latin text to a runic text.
func TransposeLatinToRune(text string, encodeBackwards bool) string {
	var sb strings.Builder
	repo := runelib.NewCharacterRepo()
	text = strings.ToUpper(text)

	if encodeBackwards {
		text = reverseWords(text, repo)
	}

	for i := 0; i < len(text); i++ {
		xchar := string(text[i])
		if !repo.IsRune(xchar, true) {
			switch xchar {
			case "A":
				if i+1 < len(text) && text[i+1] == 'E' {
					sb.WriteString(repo.GetRuneFromChar("AE"))
					i++
				} else {
					sb.WriteString(repo.GetRuneFromChar("A"))
				}
			case "E":
				if i+1 < len(text) && text[i+1] == 'A' {
					sb.WriteString(repo.GetRuneFromChar("EA"))
					i++
				} else if i+1 < len(text) && text[i+1] == 'O' {
					sb.WriteString(repo.GetRuneFromChar("EO"))
					i++
				} else {
					sb.WriteString(repo.GetRuneFromChar("E"))
				}
			case "O":
				if i+1 < len(text) && text[i+1] == 'E' {
					sb.WriteString(repo.GetRuneFromChar("OE"))
					i++
				} else {
					sb.WriteString(repo.GetRuneFromChar("O"))
				}
			case "T":
				if i+1 < len(text) && text[i+1] == 'H' {
					sb.WriteString(repo.GetRuneFromChar("TH"))
					i++
				} else {
					sb.WriteString(repo.GetRuneFromChar("T"))
				}
			case "I":
				if i+1 < len(text) && text[i+1] == 'O' {
					sb.WriteString(repo.GetRuneFromChar("IO"))
					i++
				} else if i+2 < len(text) && text[i+1] == 'N' && text[i+2] == 'G' {
					sb.WriteString(repo.GetRuneFromChar("ING"))
					i += 2
				} else if i+1 < len(text) && text[i+1] == 'A' {
					sb.WriteString(repo.GetRuneFromChar("IA"))
					i++
				} else {
					sb.WriteString(repo.GetRuneFromChar("I"))
				}
			case "N":
				if i+1 < len(text) && text[i+1] == 'G' {
					sb.WriteString(repo.GetRuneFromChar("NG"))
					i++
				} else {
					sb.WriteString(repo.GetRuneFromChar("N"))
				}
			default:
				sb.WriteString(repo.GetRuneFromChar(xchar))
			}
		} else {
			sb.WriteString(repo.GetRuneFromChar(xchar))
		}
	}

	return sb.String()
}

// reverseString reverses the given string and returns the reversed result.
func reverseString(text string) string {
	var sb strings.Builder
	for i := len(text) - 1; i >= 0; i-- {
		sb.WriteString(string(text[i]))
	}
	return sb.String()
}

// reverseWords reverses all words in the given text while preserving the original order of delimiters.
// A word is defined as a sequence of characters not classified as separators by the provided CharacterRepo.
func reverseWords(text string, repo *runelib.CharacterRepo) string {
	var retval []string
	charArray := strings.Split(text, "")
	var sb strings.Builder
	for i := 0; i < len(charArray); i++ {
		if repo.IsSeperator(charArray[i]) {
			retval = append(retval, reverseString(sb.String()))
			retval = append(retval, charArray[i])
			sb.Reset()
		} else {
			sb.WriteString(charArray[i])
		}
	}

	if sb.Len() > 0 {
		retval = append(retval, reverseString(sb.String()))
	}

	return strings.Join(retval, "")
}
