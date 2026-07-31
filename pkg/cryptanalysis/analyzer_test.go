package cryptanalysis

import "testing"

func TestAnalyzeEnglishText(t *testing.T) {
	res, err := Analyze("THISISASIMPLEENGLISHTEXTWITHSOMELETTERFREQUENCIES", "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if err != nil {
		t.Fatalf("Analyze failed: %v", err)
	}

	if res.IoC <= 0 {
		t.Fatalf("expected IoC > 0, got %f", res.IoC)
	}
	if res.Entropy <= 0 {
		t.Fatalf("expected entropy > 0, got %f", res.Entropy)
	}
	if len(res.Ranking) == 0 {
		t.Fatal("expected ranking results")
	}
}

func TestAnalyzeCustomAlphabetNumeric(t *testing.T) {
	res, err := Analyze("123123451212345", "12345")
	if err != nil {
		t.Fatalf("Analyze failed: %v", err)
	}

	if res.CharacterSet.AlphabetSize != 5 {
		t.Fatalf("expected alphabet size 5, got %d", res.CharacterSet.AlphabetSize)
	}
	if !res.CharacterSet.ContainsNumbers {
		t.Fatal("expected number detection")
	}
}

func TestMatsuiAlgorithms(t *testing.T) {
	input := MatsuiInput{
		Samples: []MatsuiSample{
			{Plaintext: 0x3, Ciphertext: 0x5, KeyGuess: 0x1},
			{Plaintext: 0x2, Ciphertext: 0x7, KeyGuess: 0x1},
			{Plaintext: 0x1, Ciphertext: 0x4, KeyGuess: 0x2},
		},
		Approximations: []LinearApproximation{
			{PlainMask: 0x1, CipherMask: 0x1},
			{PlainMask: 0x2, CipherMask: 0x2},
		},
	}

	res := AnalyzeMatsui(input)

	if !res.Available {
		t.Fatal("expected Matsui available")
	}
	if len(res.Algorithm1) != 2 {
		t.Fatalf("expected 2 algorithm1 results, got %d", len(res.Algorithm1))
	}
	if len(res.Algorithm2Rank) == 0 {
		t.Fatal("expected key ranking results")
	}
}
