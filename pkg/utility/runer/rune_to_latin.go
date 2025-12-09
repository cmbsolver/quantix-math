package runer

import (
	"quantix-math/pkg/utility/runelib"
	"strings"
)

// TransposeRuneToLatin transposes a string of runes to Latin characters
func TransposeRuneToLatin(text string) string {
	var sb strings.Builder
	repo := runelib.NewCharacterRepo()

	for _, runeCharacter := range text {
		character := repo.GetCharFromRune(string(runeCharacter))
		if character != "" {
			sb.WriteString(character)
		} else {
			sb.WriteRune(runeCharacter)
		}
	}

	return sb.String()
}
