package loader

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"quantix-math/pkg/db"
	"quantix-math/pkg/db/tables"
	"quantix-math/pkg/sequences"
	"quantix-math/pkg/utility/runer"
	"regexp"
	"strings"
	"unicode/utf8"
)

// LoadWords loads all words from a directory into the database.
func LoadWords() {
	dbConn, _ := db.InitConnection()
	if tables.GetRecordCount(dbConn) > 0 {
		return
	}

	// Define flags for the directory path and output SQL file
	dir := "./assets/files/words"

	wordList := make(map[string]bool)
	dictList := make([]tables.DictionaryWord, 0, 16384)

	// Scan the directory for text files
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".txt") {
			processFile(path, wordList)
			fmt.Printf("Loaded %d words from %s\n", len(wordList), path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Failed to read directory: %v\n", err)
		return
	}

	// Read every word from the word list
	for word := range wordList {
		word = strings.ToUpper(word)
		word = filterNumbersOut(word)

		if len(word) <= 0 {
			continue
		}

		runeglish := runer.PrepLatinToRune(word)
		runeText := runer.TransposeLatinToRune(runeglish, false)
		runeTextNoDoublet := tables.RemoveDoublets(strings.Split(runeText, ""))
		gemSum := runer.CalculateGemSum(runeText, runer.Runes, false)
		gemProd := runer.CalculateGemProduct(runeText, runer.Runes, false)

		dictWord := tables.DictionaryWord{
			DictionaryWordText:          word,
			RuneglishWordText:           runeglish,
			RuneWordText:                runeText,
			RuneWordTextNoDoublet:       runeTextNoDoublet,
			GemSum:                      gemSum,
			GemSumPrime:                 sequences.IsPrime(big.NewInt(gemSum)),
			GemProduct:                  gemProd.String(),
			GemProductPrime:             sequences.IsPrime(&gemProd),
			DictionaryWordLength:        len(word),
			RuneglishWordLength:         len(runeglish),
			DictRuneNoDoubletLength:     utf8.RuneCountInString(runeTextNoDoublet),
			RuneWordLength:              utf8.RuneCountInString(runeText),
			RunePattern:                 tables.GetRunePattern(word),
			RunePatternNoDoubletPattern: tables.GetRunePattern(runeTextNoDoublet),
			RuneDistancePattern:         tables.GetRuneDistancePattern(strings.Split(runeText, "")),
			Language:                    "English",
		}

		dictList = append(dictList, dictWord)
		delete(wordList, word) // Remove the word from wordList

		if len(dictList) >= 500 {
			tables.AddDictionaryWords(dbConn, dictList)
			dictList = dictList[:0]
		}
	}

	if len(dictList) > 0 {
		tables.AddDictionaryWords(dbConn, dictList)
	}

	_ = db.CloseConnection(dbConn)
}

// processFile reads a file and adds all words to the word list
func processFile(filePath string, wordList map[string]bool) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer func(file *os.File) {
		closeError := file.Close()
		if closeError != nil {
			fmt.Printf("Failed to close file: %v\n", closeError)
		}
	}(file)

	reader := bufio.NewReader(file)
	re := regexp.MustCompile(`[^\w]+`)
	for {
		line, readError := reader.ReadString('\n')
		if readError != nil {
			if readError.Error() != "EOF" {
				fmt.Printf("Error reading file: %v\n", err)
			}
			break
		}
		words := re.Split(line, -1)
		for _, word := range words {
			if word != "" {
				wordList[word] = true
			}
		}
	}
}

// filterNumbersOut removes all numeric characters from the input string, retaining only alphabetic characters and valid symbols.
func filterNumbersOut(text string) string {
	wordArray := strings.Split(text, "")
	var newWordArray []string
	for _, character := range wordArray {
		if strings.ContainsAny(character, "ABCDEFGHIJKLMOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'-") {
			newWordArray = append(newWordArray, character)
		}
	}
	return strings.Join(newWordArray, "")
}
