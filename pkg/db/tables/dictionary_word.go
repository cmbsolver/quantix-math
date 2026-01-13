package tables

import (
	"fmt"
	"math"
	"quantix-math/pkg/utility/runelib"
	"slices"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// DictionaryWord represents a structured word entry with various properties for linguistic and numerical analysis.
type DictionaryWord struct {
	DictionaryWordText          string `gorm:"column:dict_word"`
	RuneglishWordText           string `gorm:"column:dict_runeglish"`
	RuneWordText                string `gorm:"column:dict_rune"`
	RuneWordTextNoDoublet       string `gorm:"column:dict_rune_no_doublet"`
	GemSum                      int64  `gorm:"column:gem_sum"`
	GemSumPrime                 bool   `gorm:"column:gem_sum_prime"`
	GemProduct                  string `gorm:"column:gem_product"`
	GemProductPrime             bool   `gorm:"column:gem_product_prime"`
	DictionaryWordLength        int    `gorm:"column:dict_word_length"`
	RuneglishWordLength         int    `gorm:"column:dict_runeglish_length"`
	RuneWordLength              int    `gorm:"column:dict_rune_length"`
	DictRuneNoDoubletLength     int    `gorm:"column:dict_rune_no_doublet_length"`
	RunePattern                 string `gorm:"column:rune_pattern"`
	RunePatternNoDoubletPattern string `gorm:"column:rune_pattern_no_doublet"`
	RuneDistancePattern         string `gorm:"column:rune_distance_pattern"`
	Language                    string `gorm:"column:language"`
}

func (DictionaryWord) TableName() string {
	return "dictionary_words"
}

func AddDictionaryWords(db *gorm.DB, dictionaryWords []DictionaryWord) {
	db.Create(&dictionaryWords)
	return
}

func GetDictionaryWordsByFilters(db *gorm.DB, filters map[string]int) []DictionaryWord {
	var dictionaryWords []DictionaryWord

	query := db.Order("dict_word ASC")
	for field, value := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", field), value)
	}

	query.Find(&dictionaryWords)

	return sortDistinctDictionaryWords(dictionaryWords)
}

func sortDistinctDictionaryWords(words []DictionaryWord) []DictionaryWord {
	// Deduplicate by DictionaryWordText (dict_word) first.
	seen := make(map[string]DictionaryWord, len(words))
	keys := make([]string, 0, len(words))

	for _, w := range words {
		k := w.DictionaryWordText
		if k == "" {
			// Ignore empty keys to avoid odd "blank word" entries.
			continue
		}
		if _, exists := seen[k]; exists {
			continue
		}
		seen[k] = w
		keys = append(keys, k)
	}

	// Sort by DictionaryWordText ascending.
	slices.Sort(keys)

	// Rebuild result in sorted order.
	out := make([]DictionaryWord, 0, len(keys))
	for _, k := range keys {
		out = append(out, seen[k])
	}

	return out
}

// GetRunePattern gets the rune pattern for the dictionary word
func GetRunePattern(word string) string {
	patternDictionary := make(map[int]string)
	var runes []string
	counter := 1

	for _, character := range word {
		if character == '\'' {
			runes = append(runes, "'")
			continue
		}

		found := false
		for key, value := range patternDictionary {
			if value == string(character) {
				runes = append(runes, fmt.Sprintf("%d", key))
				found = true
				break
			}
		}

		if !found {
			runes = append(runes, fmt.Sprintf("%d", counter))
			patternDictionary[counter] = string(character)
			counter++
		}
	}

	return strings.Join(runes, ",")
}

// RemoveDoublets removes consecutive duplicate characters from a word
func RemoveDoublets(word []string) string {
	if len(word) == 0 {
		return ""
	}

	var result strings.Builder
	result.WriteString(word[0])

	for i := 1; i < len(word); i++ {
		if word[i] != word[i-1] {
			result.WriteString(word[i])
		}
	}

	return result.String()
}

// GetRuneDistancePattern generates a string representing the distances between adjacent runes in a slice of strings.
func GetRuneDistancePattern(word []string) string {
	charRepo := runelib.NewCharacterRepo()
	gemRunes := charRepo.GetGematriaRunes()

	if len(word) == 0 {
		return ""
	}

	var result strings.Builder
	currentValue := getRuneIndex(word[0], gemRunes)

	result.WriteString(strconv.Itoa(0))

	for i := 1; i < len(word); i++ {
		currentDistance := currentValue - getRuneIndex(word[i], gemRunes)
		distance := int(math.Abs(float64(currentDistance)))
		result.WriteString(fmt.Sprintf(", %s", strconv.Itoa(distance)))
		currentValue = getRuneIndex(word[i], gemRunes)
	}

	return result.String()
}

// getRuneIndex finds the index of a given rune in the provided alphabet slice.
// Returns the index if found, otherwise returns -1.
func getRuneIndex(rune string, alphabet []string) int {
	for i, r := range alphabet {
		if r == rune {
			return i
		}
	}

	return -1
}

func GetRecordCount(db *gorm.DB) int64 {
	var count int64
	db.Model(&DictionaryWord{}).Count(&count)
	return count
}
