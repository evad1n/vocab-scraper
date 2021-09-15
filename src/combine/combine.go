package combine

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
)

type (
	ExtractionFunc func([]byte) ([]string, error)
)

func WriteWords(words []string, fileName string) error {
	// Open output file
	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("creating out file: %v", err)
	}
	defer f.Close()

	// Write words line by line
	for _, word := range words {
		f.WriteString(word + "\n")
	}

	return nil
}

func GetUniqueWords(words []string) []string {
	var uniques = make(map[string]bool)
	for _, w := range words {
		uniques[w] = true
	}

	// Sort alphabetically
	sortedWords := make([]string, len(uniques))
	i := 0
	for word := range uniques {
		sortedWords[i] = word
		i++
	}

	sort.Strings(sortedWords)

	return sortedWords
}

func WordsFromFile(fileName string, extractionFunc ExtractionFunc) ([]string, error) {
	dirname, _ := os.Getwd()
	fullPath := path.Join(dirname, fileName)

	ext := filepath.Ext(fileName)[1:]
	switch ext {
	case "txt":
		return WordsFromTxtFile(fullPath)
	case "json":
		return WordsFromJsonFile(fullPath, extractionFunc)
	default:
		return nil, fmt.Errorf("unsupported file type '%s'", fileName)
	}
}

// Reads words from a text file into a string array.
//
// Expects 1 word per line
func WordsFromTxtFile(fileName string) ([]string, error) {
	fmt.Println(os.Getwd())
	// inFile, err := os.Open(path.Join(root, fileName))
	inFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("opening file '%s': %v", fileName, err)
	}
	defer inFile.Close()

	var words []string

	fileScanner := bufio.NewScanner(inFile)
	// Read file
	i := 0
	for fileScanner.Scan() {
		i++
		word := fileScanner.Text()
		words = append(words, word)
	}
	if err := fileScanner.Err(); err != nil {
		return nil, fmt.Errorf("reading words: %v", err)
	}

	return words, nil
}

// Reads words from a json file into a string array.
//
// Expects json to be formatted according to define.Definition (default Lexicogn format)
func WordsFromJsonFile(fileName string, extractionFunc ExtractionFunc) ([]string, error) {
	fmt.Println(os.Getwd())
	// inFile, err := os.Open(path.Join(root, fileName))
	inFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("opening file '%s': %v", fileName, err)
	}
	defer inFile.Close()

	// Read file
	bytes, err := ioutil.ReadAll(inFile)
	if err != nil {
		return nil, fmt.Errorf("reading file '%s': %v", fileName, err)
	}

	words, err := extractionFunc(bytes)
	if err != nil {
		return nil, fmt.Errorf("extraction func: %v", err)
	}

	return words, nil
}
