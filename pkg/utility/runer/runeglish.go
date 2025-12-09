package runer

import (
	"strings"
)

// PrepLatinToRune converts Latin text to a format suitable for rune conversion.
func PrepLatinToRune(text string) string {
	text = strings.ToUpper(text)

	text = strings.ReplaceAll(text, "QU", "CW")
	text = strings.ReplaceAll(text, "Z", "S")
	text = strings.ReplaceAll(text, "K", "C")
	text = strings.ReplaceAll(text, "Q", "C")
	text = strings.ReplaceAll(text, "V", "U")

	var sb strings.Builder

	for i := 0; i < len(text); i++ {
		xchar := text[i]

		switch xchar {
		case 'I':
			if i+1 < len(text) && text[i+1] == 'O' {
				sb.WriteString("IO")
				i++
			} else if i+1 < len(text) && text[i+1] == 'A' {
				sb.WriteString("IO")
				i++
			} else {
				sb.WriteByte('I')
			}
		default:
			sb.WriteByte(xchar)
		}
	}

	return sb.String()
}
