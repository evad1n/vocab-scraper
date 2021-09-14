package combine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/evad1n/vocab-scraper/define"
)

func CombineFiles() {

}

// Read JSON or TXT
func combineFiles(file1Name string, file2Name string, outFileName string) {
	dirname, _ := os.Getwd()
	root := path.Join(dirname, "../")

	// Open file 1
	inFile, err := os.Open(path.Join(root, file1Name))
	if err != nil {
		log.Fatalf("opening comparison file 1: %v", err)
	}
	defer inFile.Close()
	fmt.Printf("Reading words from '%s'\n\n", inFileName)

	fileScanner := bufio.NewScanner(inFile)
	// Read file
	i := 0
	for fileScanner.Scan() {
		i++
		word := fileScanner.Text()
		fmt.Printf("%d: %s\n", i, word)
		definitions = append(definitions, getDef(word))
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("reading words: %v", err)
	}

	// Open output file
	outFile, err := os.Create(path.Join(root, outFileName))
	if err != nil {
		log.Fatalf("creating out file: %v", err)
	}
	defer outFile.Close()

	// JSON
	data, err := json.MarshalIndent(definitions, "", " ")
	if err != nil {
		log.Fatalf("marshalling json: %v", err)
	}

	// Write output
	outFile.Write(data)
	fmt.Printf("\nDefinitions written to '%s'\n", outFileName)
}

// Gets file names
func getFileNames() (string, string, string) {
	inputScanner := bufio.NewScanner(os.Stdin)

	// File 1
	fmt.Print("File 1 (default 'words.txt'): ")
	inputScanner.Scan()
	file1Name := inputScanner.Text()
	if file1Name == "" {
		file1Name = "in/test.txt"
	}

	// File 2
	fmt.Print("File 2 (default 'definitions.json'): ")
	inputScanner.Scan()
	file2Name := inputScanner.Text()
	if file2Name == "" {
		file2Name = "out/definitions.json"
	}

	// Combined file
	fmt.Print("File 2 (default 'definitions.json'): ")
	inputScanner.Scan()
	combinedFileName := inputScanner.Text()
	if combinedFileName == "" {
		combinedFileName = "out/definitions.json"
	}

	fmt.Printf("\nFile 1:  %s\nFile 2: %s\nCombined file: %s\n", file1Name, file2Name, combinedFileName)

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
