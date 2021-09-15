package combine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"

	"github.com/evad1n/vocab-scraper/define"
)

// Read JSON or TXT
func CombineFiles() {
	file1Name, file2Name, combinedFileName := getFileNames()

	words1 := wordsFromFile(file1Name)
	words2 := wordsFromFile(file2Name)

	sortedWords := getUniqueWords(words1, words2)

	writeWords(sortedWords, combinedFileName)

}

func writeWords(words []string, fileName string) {
	// Open output file
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("creating out file: %v", err)
	}
	defer f.Close()

	// Write words line by line
	for _, word := range words {
		f.WriteString(word + "\n")
	}

	fmt.Printf("\nCombines words written to '%s'\n", fileName)
}

func getUniqueWords(words1 []string, words2 []string) []string {
	var uniques = make(map[string]bool)
	for _, w := range words1 {
		uniques[w] = true
	}
	for _, w := range words2 {
		uniques[w] = true
	}

	sortedWords := make([]string, len(uniques))
	i := 0
	for word := range uniques {
		sortedWords[i] = word
		i++
	}

	sort.Strings(sortedWords)

	return sortedWords
}

func wordsFromFile(fileName string) []string {
	dirname, _ := os.Getwd()
	fullPath := path.Join(dirname, fileName)

	ext := filepath.Ext(fileName)[1:]
	switch ext {
	case "txt":
		return wordsFromTxtFile(fullPath)
	case "json":
		return wordsFromJsonFile(fullPath)
	default:
		log.Fatalf("unsupported file type '%s'\n", fileName)
	}

	return []string{}
}

// Gets file names
func getFileNames() (string, string, string) {
	inputScanner := bufio.NewScanner(os.Stdin)

	// File 1
	fmt.Print("\nFile 1 (default 'words1.txt'): ")
	inputScanner.Scan()
	file1Name := inputScanner.Text()
	if file1Name == "" {
		file1Name = "in/words.txt"
	}

	// File 2
	fmt.Print("File 2 (default 'words2.txt'): ")
	inputScanner.Scan()
	file2Name := inputScanner.Text()
	if file2Name == "" {
		file2Name = "in/words2.json"
	}

	// Combined file
	fmt.Print("File 2 (default 'combined.txt'): ")
	inputScanner.Scan()
	combinedFileName := inputScanner.Text()
	if combinedFileName == "" {
		combinedFileName = "out/combined.txt"
	}

	fmt.Printf("\nFile 1: %s\nFile 2: %s\nCombined file: %s\n", file1Name, file2Name, combinedFileName)

	return file1Name, file2Name, combinedFileName
}

// Reads words from a text file into a string array.
//
// Expects 1 word per line
func wordsFromTxtFile(fileName string) []string {
	fmt.Println(os.Getwd())
	// inFile, err := os.Open(path.Join(root, fileName))
	inFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("opening file %s: %v", fileName, err)
	}
	defer inFile.Close()

	var words []string
	fmt.Printf("Reading words from '%s'\n\n", fileName)

	fileScanner := bufio.NewScanner(inFile)
	// Read file
	i := 0
	for fileScanner.Scan() {
		i++
		word := fileScanner.Text()
		words = append(words, word)
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("reading words: %v", err)
	}

	fmt.Printf("Read %d words from %s", i, fileName)

	return words
}

// Reads words from a json file into a string array.
//
// Expects json to be formatted according to define.Definition (default Lexicogn format)
func wordsFromJsonFile(fileName string) []string {
	fmt.Println(os.Getwd())
	// inFile, err := os.Open(path.Join(root, fileName))
	inFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("opening file %s: %v", fileName, err)
	}
	defer inFile.Close()

	fmt.Printf("Reading words from '%s'\n\n", fileName)

	// Read file
	bytes, err := ioutil.ReadAll(inFile)
	if err != nil {
		log.Fatalf("reading file %s: %v", fileName, err)
	}

	var definitions []define.Definition
	if err := json.Unmarshal(bytes, definitions); err != nil {
		log.Fatalf("decoding json from file %s: %v", fileName, err)
	}

	fmt.Printf("Read %d words from %s", len(definitions), fileName)

	words := make([]string, len(definitions))
	for i, def := range definitions {
		words[i] = def.Word
	}

	return words
}
