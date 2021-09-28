package combine

import (
	"encoding/json"
	"log"
	"sort"
	"testing"
)

func TestGetUniqueWords(t *testing.T) {
	words := []string{
		"apple",
		"elephant",
		"coctagon",
		"squidward",
		"walrus",
		"dog",
		"square",
		"cat",
		"dog",
		"apple",
	}
	sortedUniqueWords := GetUniqueWords(words)

	dups := make(map[string]bool)

	for _, word := range sortedUniqueWords {
		if _, exists := dups[word]; exists {
			t.Fatalf("duplicate word found: %q", word)
		}
		dups[word] = true
	}

	if !sort.SliceIsSorted(sortedUniqueWords, func(i, j int) bool {
		return sortedUniqueWords[i] < sortedUniqueWords[j]
	}) {
		t.Fatalf("words are not sorted alphabetically")
	}
}

func TestWordsFromTxtFile(t *testing.T) {
	words, err := WordsFromFile("test-words.txt", nil)
	if err != nil {
		t.Fatalf("reading words from file: %v", err)
	}

	compareWords := []string{
		"apple",
		"elephant",
		"coctagon",
		"squidward",
		"walrus",
		"dog",
		"square",
		"cat",
		"dog",
		"apple",
	}

	for i, compare := range compareWords {
		if words[i] != compare {
			t.Fatalf("read wrong words: expected %q got %q", compare, words[i])
		}
	}
}

func TestWordsFromJSONFile(t *testing.T) {
	type WordJSON struct {
		Word       string `json:"word"`
		Definition string `json:"definition"`
	}

	// Edit this function depending on JSON schema...
	var extractFunc ExtractionFunc = func(bytes []byte) ([]string, error) {
		var data []WordJSON
		if err := json.Unmarshal(bytes, &data); err != nil {
			log.Fatalf("decoding json: %v", err)
		}

		words := make([]string, len(data))
		for i, def := range data {
			words[i] = def.Word
		}

		return words, nil
	}

	words, err := WordsFromFile("test-words.json", extractFunc)
	if err != nil {
		t.Fatalf("reading words from file: %v", err)
	}

	compareWords := []string{
		"apple",
		"elephant",
		"coctagon",
		"squidward",
		"walrus",
		"dog",
		"square",
		"cat",
		"dog",
		"apple",
	}

	for i, compare := range compareWords {
		if words[i] != compare {
			t.Fatalf("read wrong words: expected %q got %q", compare, words[i])
		}
	}
}
