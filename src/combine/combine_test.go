package combine

import (
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
			t.Fatalf("duplicate word found: '%s'", word)
		}
		dups[word] = true
	}

	if !sort.SliceIsSorted(sortedUniqueWords, func(i, j int) bool {
		return sortedUniqueWords[i] < sortedUniqueWords[j]
	}) {
		t.Fatalf("words are not sorted alphabetically")
	}
}
