package interactive

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/evad1n/vocab-scraper/combine"
)

type (
	CombineJSON struct {
		Word       string `json:"word"`
		Definition string `json:"definition"`
	}
)

const (
	defaultInFolder     = "in"
	defaultCombinedFile = "out/combined.txt"
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
	fileNames, combinedFileName, err := getCombineFileNames()
	if err != nil {
		return fmt.Errorf("getting file names: %v", err)
	}

	words := []string{}
	for _, fileName := range fileNames {
		addedWords, err := combine.WordsFromFile(fileName, extractFunc)
		if err != nil {
			return fmt.Errorf("reading words from file '%s': %v", fileName, err)
		}

		fmt.Printf("Read %d words from %s", len(addedWords), fileName)

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
func getCombineFileNames() ([]string, string, error) {
	inputScanner := bufio.NewScanner(os.Stdin)

	// Files
	fmt.Printf("\nFolder name with input files (default '%s'): ", defaultInFolder)
	inputScanner.Scan()
	folderName := inputScanner.Text()
	if folderName == "" {
		folderName = defaultInFolder
	}

	// Combined file
	fmt.Printf("Combine file (default '%s'): ", defaultCombinedFile)
	inputScanner.Scan()
	combinedFileName := inputScanner.Text()
	if combinedFileName == "" {
		combinedFileName = defaultCombinedFile
	}

	// Get all files in the specified folder
	files, err := ioutil.ReadDir(folderName)
	if err != nil {
		return nil, combinedFileName, fmt.Errorf("reading directory '%s': %v", folderName, err)
	}

	fileNames := make([]string, len(files))
	fmt.Printf("Found %d files to combine:\n", len(files))

	for i, f := range files {
		fmt.Println(i, f.Name())
		fileNames[i] = path.Join(folderName, f.Name())
	}

	fmt.Printf("\nWill combine to file: %s\n", combinedFileName)

	return fileNames, combinedFileName, nil
}
