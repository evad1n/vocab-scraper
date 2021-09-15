package interactive

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/evad1n/vocab-scraper/combine"
)

type (
	CombineJSON struct {
		Word       string `json:"word"`
		Definition string `json:"definition"`
	}
)

// Edit this function depending on JSON schema...
var extractFunc combine.ExtractionFunc = func(bytes []byte) ([]string, error) {
	var data []CombineJSON
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatalf("decoding json: %v", err)
	}

	words := make([]string, len(data))
	for i, def := range data {
		words[i] = def.Word
	}

	return words, nil
}

// Read JSON or TXT
func CombineFiles() error {
	fileNames, combinedFileName := getCombineFileNames()

	words := []string{}
	for _, fileName := range fileNames {
		addedWords, err := combine.WordsFromFile(fileName, extractFunc)
		if err != nil {
			return fmt.Errorf("reading words from file '%s': %v", fileName, err)
		}

		fmt.Printf("Read %d words from %s", len(definitions), fileName)

		words = append(words, addedWords...)
	}

	sortedWords := combine.GetUniqueWords(words)

	if err := combine.WriteWords(sortedWords, combinedFileName); err != nil {
		return fmt.Errorf("writing words to file: %v", err)
	}

	fmt.Printf("\nCombines words written to '%s'\n", combinedFileName)

	return nil
}

// Gets file names
//
// Returns a list of input files, and the name of the output file
func getCombineFileNames() ([]string, string) {
	inputScanner := bufio.NewScanner(os.Stdin)

	// Files
	fmt.Print("\nFolder name with input files (default 'in'): ")
	inputScanner.Scan()
	folderName := inputScanner.Text()
	if folderName == "" {
		folderName = "in/words.txt"
	}
	fmt.Println(fileNames)

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

	return nil, combinedFileName
}
