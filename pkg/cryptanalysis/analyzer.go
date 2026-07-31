package cryptanalysis

import (
	"errors"
	"math"
	"sort"
	"strings"
	"unicode"
)

var englishFrequencies = map[rune]float64{
	'A': 8.167, 'B': 1.492, 'C': 2.782, 'D': 4.253, 'E': 12.702,
	'F': 2.228, 'G': 2.015, 'H': 6.094, 'I': 6.966, 'J': 0.153,
	'K': 0.772, 'L': 4.025, 'M': 2.406, 'N': 6.749, 'O': 7.507,
	'P': 1.929, 'Q': 0.095, 'R': 5.987, 'S': 6.327, 'T': 9.056,
	'U': 2.758, 'V': 0.978, 'W': 2.360, 'X': 0.150, 'Y': 1.974,
	'Z': 0.074,
}

type MatsuiSample struct {
	Plaintext  uint64 `json:"plaintext"`
	Ciphertext uint64 `json:"ciphertext"`
	KeyGuess   uint64 `json:"key_guess"`
}

type LinearApproximation struct {
	PlainMask  uint64 `json:"plain_mask"`
	CipherMask uint64 `json:"cipher_mask"`
}

type MatsuiInput struct {
	Samples        []MatsuiSample        `json:"samples"`
	Approximations []LinearApproximation `json:"approximations"`
}

type MatsuiApproximationResult struct {
	PlainMask   uint64  `json:"plain_mask"`
	CipherMask  uint64  `json:"cipher_mask"`
	Count       int     `json:"count"`
	Total       int     `json:"total"`
	Bias        float64 `json:"bias"`
	Probability float64 `json:"probability"`
	Score       float64 `json:"score"`
}

type MatsuiResult struct {
	Available      bool                        `json:"available"`
	Message        string                      `json:"message"`
	Algorithm1     []MatsuiApproximationResult `json:"algorithm_1"`
	Algorithm2Rank []MatsuiKeyRank             `json:"algorithm_2_rank"`
}

type MatsuiKeyRank struct {
	KeyGuess uint64  `json:"key_guess"`
	Score    float64 `json:"score"`
}

type RankedCipher struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

type KasiskiResult struct {
	Distances           []int          `json:"distances"`
	CandidateKeyLengths map[int]int    `json:"candidate_key_lengths"`
	Repeats             map[string]int `json:"repeats"`
}

type CharacterSetReport struct {
	Alphabet            string  `json:"alphabet"`
	AlphabetSize        int     `json:"alphabet_size"`
	ObservedUnique      int     `json:"observed_unique"`
	Coverage            float64 `json:"coverage"`
	ContainsLetters     bool    `json:"contains_letters"`
	ContainsNumbers     bool    `json:"contains_numbers"`
	ContainsPunctuation bool    `json:"contains_punctuation"`
}

type Result struct {
	NormalizedText    string             `json:"normalized_text"`
	CharacterSet      CharacterSetReport `json:"character_set"`
	Length            int                `json:"length"`
	IoC               float64            `json:"ioc"`
	ChiSquared        float64            `json:"chi_squared"`
	Entropy           float64            `json:"entropy"`
	NormalizedEntropy float64            `json:"normalized_entropy"`
	Kasiski           KasiskiResult      `json:"kasiski"`
	Ranking           []RankedCipher     `json:"ranking"`
}

func Analyze(text string, alphabet string) (Result, error) {
	prepared, usedAlphabet := normalizeText(text, alphabet)
	if len(prepared) < 2 {
		return Result{}, errors.New("input text is too short after normalization")
	}

	counts := make(map[rune]int)
	for _, r := range prepared {
		counts[r]++
	}

	length := len(prepared)
	ioc := computeIoC(counts, length)
	chiSquared := computeChiSquared(counts, length, usedAlphabet)
	entropy := computeEntropy(counts, length)
	normEntropy := 0.0
	if len([]rune(usedAlphabet)) > 1 {
		normEntropy = entropy / math.Log2(float64(len([]rune(usedAlphabet))))
	}
	kasiski := kasiskiExamination(prepared)
	ranking := rankCiphers(ioc, chiSquared, entropy, normEntropy, kasiski)

	charset := buildCharacterSetReport(text, usedAlphabet, counts)

	return Result{
		NormalizedText:    string(prepared),
		CharacterSet:      charset,
		Length:            length,
		IoC:               ioc,
		ChiSquared:        chiSquared,
		Entropy:           entropy,
		NormalizedEntropy: normEntropy,
		Kasiski:           kasiski,
		Ranking:           ranking,
	}, nil
}

func AnalyzeMatsui(input MatsuiInput) MatsuiResult {
	return runMatsui(input)
}

func normalizeText(text string, alphabet string) ([]rune, string) {
	set := map[rune]bool{}
	for _, r := range alphabet {
		if !unicode.IsSpace(r) {
			set[unicode.ToUpper(r)] = true
		}
	}

	prepared := make([]rune, 0, len(text))
	if len(set) > 0 {
		for _, r := range text {
			ru := unicode.ToUpper(r)
			if set[ru] {
				prepared = append(prepared, ru)
			}
		}
		return prepared, sortedAlphabetString(set)
	}

	for _, r := range text {
		ru := unicode.ToUpper(r)
		if unicode.IsLetter(ru) || unicode.IsNumber(ru) {
			set[ru] = true
			prepared = append(prepared, ru)
		}
	}

	return prepared, sortedAlphabetString(set)
}

func sortedAlphabetString(set map[rune]bool) string {
	keys := make([]rune, 0, len(set))
	for r := range set {
		keys = append(keys, r)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return string(keys)
}

func computeIoC(counts map[rune]int, n int) float64 {
	if n < 2 {
		return 0
	}
	numerator := 0.0
	for _, c := range counts {
		numerator += float64(c * (c - 1))
	}
	return numerator / float64(n*(n-1))
}

func computeChiSquared(counts map[rune]int, n int, alphabet string) float64 {
	if n == 0 {
		return 0
	}
	isEnglish := alphabet == "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	chi := 0.0
	for _, r := range alphabet {
		observed := float64(counts[r])
		expected := float64(n) / float64(len([]rune(alphabet)))
		if isEnglish {
			expected = (englishFrequencies[r] / 100.0) * float64(n)
		}
		if expected > 0 {
			diff := observed - expected
			chi += (diff * diff) / expected
		}
	}
	return chi
}

func computeEntropy(counts map[rune]int, n int) float64 {
	entropy := 0.0
	for _, c := range counts {
		p := float64(c) / float64(n)
		entropy -= p * math.Log2(p)
	}
	return entropy
}

func kasiskiExamination(text []rune) KasiskiResult {
	repeats := map[string][]int{}
	for size := 3; size <= 5; size++ {
		for i := 0; i+size <= len(text); i++ {
			gram := string(text[i : i+size])
			repeats[gram] = append(repeats[gram], i)
		}
	}

	distances := make([]int, 0)
	repeatCounts := make(map[string]int)
	for gram, positions := range repeats {
		if len(positions) < 2 {
			continue
		}
		repeatCounts[gram] = len(positions)
		for i := 0; i < len(positions)-1; i++ {
			d := positions[i+1] - positions[i]
			if d > 0 {
				distances = append(distances, d)
			}
		}
	}

	candidates := make(map[int]int)
	for _, d := range distances {
		for f := 2; f <= 20; f++ {
			if d%f == 0 {
				candidates[f]++
			}
		}
	}

	return KasiskiResult{
		Distances:           distances,
		CandidateKeyLengths: candidates,
		Repeats:             flattenRepeatCounts(repeatCounts),
	}
}

func flattenRepeatCounts(in map[string]int) map[string]int {
	out := make(map[string]int)
	for k, v := range in {
		out[k] = v
	}
	return out
}

func buildCharacterSetReport(original string, alphabet string, counts map[rune]int) CharacterSetReport {
	hasLetters := false
	hasNumbers := false
	hasPunct := false
	for _, r := range original {
		switch {
		case unicode.IsLetter(r):
			hasLetters = true
		case unicode.IsNumber(r):
			hasNumbers = true
		case !unicode.IsSpace(r):
			hasPunct = true
		}
	}

	coverage := 0.0
	if len([]rune(alphabet)) > 0 {
		coverage = float64(len(counts)) / float64(len([]rune(alphabet)))
	}

	return CharacterSetReport{
		Alphabet:            alphabet,
		AlphabetSize:        len([]rune(alphabet)),
		ObservedUnique:      len(counts),
		Coverage:            coverage,
		ContainsLetters:     hasLetters,
		ContainsNumbers:     hasNumbers,
		ContainsPunctuation: hasPunct,
	}
}

func rankCiphers(ioc float64, chiSquared float64, entropy float64, normEntropy float64, kasiski KasiskiResult) []RankedCipher {
	englishDistance := math.Abs(ioc-0.0667) + math.Abs(entropy-4.1)/4.0 + math.Min(chiSquared/150.0, 1.0)
	vigenereSignal := math.Abs(ioc-0.045) + math.Abs(normEntropy-0.92)
	if len(kasiski.CandidateKeyLengths) > 0 {
		vigenereSignal -= 0.2
	}
	randomSignal := math.Abs(normEntropy - 1.0)

	candidates := []RankedCipher{
		{Name: "Monoalphabetic Substitution / Caesar", Score: 1 - clamp01(englishDistance/2)},
		{Name: "Polyalphabetic (Vigenere-like)", Score: 1 - clamp01(vigenereSignal)},
		{Name: "Transposition", Score: 1 - clamp01(math.Abs(ioc-0.0667)+math.Abs(normEntropy-0.85))},
		{Name: "Random / One-time-pad-like", Score: 1 - clamp01(randomSignal)},
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Score > candidates[j].Score
	})
	return candidates
}

func clamp01(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}

func runMatsui(input MatsuiInput) MatsuiResult {
	if len(input.Samples) == 0 || len(input.Approximations) == 0 {
		return MatsuiResult{
			Available: false,
			Message:   "Provide known plaintext/ciphertext pairs and linear approximations to run Matsui algorithms.",
		}
	}

	alg1 := make([]MatsuiApproximationResult, 0, len(input.Approximations))
	total := len(input.Samples)
	for _, approx := range input.Approximations {
		count := 0
		for _, s := range input.Samples {
			lhs := parity(s.Plaintext & approx.PlainMask)
			rhs := parity((s.Ciphertext ^ s.KeyGuess) & approx.CipherMask)
			if lhs == rhs {
				count++
			}
		}
		prob := float64(count) / float64(total)
		bias := prob - 0.5
		score := math.Abs(bias) * math.Sqrt(float64(total))
		alg1 = append(alg1, MatsuiApproximationResult{
			PlainMask:   approx.PlainMask,
			CipherMask:  approx.CipherMask,
			Count:       count,
			Total:       total,
			Bias:        bias,
			Probability: prob,
			Score:       score,
		})
	}

	keyScores := map[uint64]float64{}
	for _, s := range input.Samples {
		for _, approx := range input.Approximations {
			lhs := parity(s.Plaintext & approx.PlainMask)
			rhs := parity((s.Ciphertext ^ s.KeyGuess) & approx.CipherMask)
			if lhs == rhs {
				keyScores[s.KeyGuess] += 1
			} else {
				keyScores[s.KeyGuess] -= 1
			}
		}
	}

	ranks := make([]MatsuiKeyRank, 0, len(keyScores))
	for k, score := range keyScores {
		ranks = append(ranks, MatsuiKeyRank{KeyGuess: k, Score: score})
	}
	sort.Slice(ranks, func(i, j int) bool { return ranks[i].Score > ranks[j].Score })

	return MatsuiResult{
		Available:      true,
		Message:        "Matsui algorithm 1 statistics and algorithm 2 key ranking computed.",
		Algorithm1:     alg1,
		Algorithm2Rank: ranks,
	}
}

func parity(v uint64) uint8 {
	v ^= v >> 32
	v ^= v >> 16
	v ^= v >> 8
	v ^= v >> 4
	v ^= v >> 2
	v ^= v >> 1
	return uint8(v & 1)
}

func IsEnglishAlphabet(alphabet string) bool {
	return strings.EqualFold(alphabet, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
