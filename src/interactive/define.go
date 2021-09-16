package interactive

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/evad1n/vocab-scraper/define"
)

type (
	Definition struct {
		Word       string `json:"word"`
		Definition string `json:"definition"`
	}
)

const (
	defaultInFile  = "in/test.txt"
	defaultOutFile = "out/definitions.json"
)

var (
	definitions        []Definition
	automaticSelection = false // Automatically select the first definition
)

func FindDefinitions() error {
	dirname, _ := os.Getwd()
	root := path.Join(dirname, "../")

	inFileName, outFileName := getFileNames()

	// Automatic/Manual selection
	automaticSelection = getSelectionType()

	// Open input file
	inFile, err := os.Open(path.Join(root, inFileName))
	if err != nil {
		log.Fatalf("opening words file: %v", err)
	}
	defer inFile.Close()
	fmt.Printf("Reading words from %q\n\n", inFileName)

	fileScanner := bufio.NewScanner(inFile)
	// Read file
	i := 0
	for fileScanner.Scan() {
		i++
		word := fileScanner.Text()
		fmt.Println(strings.Repeat("=", 20))
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
	data, err := json.MarshalIndent(definitions, "", strings.Repeat(" ", 4))
	if err != nil {
		log.Fatalf("marshalling json: %v", err)
	}

	// Write output
	outFile.Write(data)
	fmt.Printf("\nDefinitions written to %q\n", outFileName)

	return nil
}

// Gets input and output filenames and returns them
func getFileNames() (string, string) {
	inputScanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("\nInput file (default %q): ", defaultInFile)
	inputScanner.Scan()
	inFileName := inputScanner.Text()
	if inFileName == "" {
		inFileName = defaultInFile
	}

	fmt.Printf("Output file (default %q): ", defaultOutFile)
	inputScanner.Scan()
	outFileName := inputScanner.Text()
	if outFileName == "" {
		outFileName = defaultOutFile
	}

	fmt.Printf("\nIn file:  %s\nOut file: %s\n\n", inFileName, outFileName)

	return inFileName, outFileName
}

// Asks user for manual/automatic definition selection
//
// Returns if selection is automatic
func getSelectionType() bool {
	fmt.Println("Selection type:")
	fmt.Println("1: Automatic")
	fmt.Println("2: Manual")
	var choice int
	fmt.Print("\nYour choice: ")
	fmt.Scanln(&choice)

	return (choice == 1)
}

func getDef(word string) Definition {
	def := Definition{
		Word: word,
	}

	i := 1

	dictDefs, err := define.DictionaryCom.Define(word)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\ndictionary.com:")
	for _, def := range dictDefs {
		fmt.Printf("%d:  %s\n", i, def)
		i++
	}

	lexicoDefs, err := define.Lexico.Define(word)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nlexico:")
	for _, def := range lexicoDefs {
		fmt.Printf("%d:  %s\n", i, def)
		i++
	}

	cambridgeDefs, err := define.Cambridge.Define(word)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\ncambridge:")
	for _, def := range cambridgeDefs {
		fmt.Printf("%d:  %s\n", i, def)
		i++
	}

	defs := append(append(dictDefs, lexicoDefs...), cambridgeDefs...)

	if automaticSelection {
		// Auto select first option
		// def.Definition = defs[0]
	} else {
		// Choose from options
		var choice int
		fmt.Println("\nChoose a definition:")
		fmt.Scanln(&choice)

		def.Definition = defs[choice-1]
	}

	fmt.Printf("\n%s\n\n", def)

	return def
}

func (d Definition) String() string {
	return fmt.Sprintf("%s : %s", d.Word, d.Definition)
}

// Endpoints

// Dictionary.com: https://www.dictionary.com/browse/%s

// https://www.google.com/#q=define+term ??
