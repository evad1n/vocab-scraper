package define

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

type (
	Definition struct {
		Word       string `json:"word"`
		Definition string `json:"definition"`
	}

	Endpoint struct {
		URL      string
		Selector string
	}
)

var (
	definitions []Definition
	auto        = false
)

func FindDefinitions() {
	dirname, _ := os.Getwd()
	root := path.Join(dirname, "../")

	inFileName, outFileName := getFileNames()

	// Automatic/Manual selection
	auto = getSelectionType()

	// Open input file
	inFile, err := os.Open(path.Join(root, inFileName))
	if err != nil {
		log.Fatalf("opening words file: %v", err)
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

// Gets input and output filenames and returns them
func getFileNames() (string, string) {
	inputScanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\nInput file (default 'words.txt'): ")
	inputScanner.Scan()
	inFileName := inputScanner.Text()
	if inFileName == "" {
		inFileName = "in/test.txt"
	}

	fmt.Print("Output file (default 'definitions.json'): ")
	inputScanner.Scan()
	outFileName := inputScanner.Text()
	if outFileName == "" {
		outFileName = "out/definitions.json"
	}

	fmt.Printf("\nIn file:  %s\nOut file: %s\n\n", inFileName, outFileName)

	return inFileName, outFileName
}

// Returns if selection type is automatic
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

	defs, err := DefineDictionaryCom(word)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dictionary.com: %v\n\n", defs)

	more, err := DefineLexico(word)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("lexico: %v\n\n", more)
	defs = append(defs, more...)

	more, err = DefineCambridge(word)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("cambridge: %v\n\n", more)
	defs = append(defs, more...)

	if auto {
		// Auto select first option
		def.Definition = defs[0]
	} else {
		// Choose from options
		fmt.Println("Choose a definition:")
		for i, d := range defs {
			fmt.Printf("%d:  %s\n", i+1, d)
		}
		var choice int
		fmt.Print("\nYour choice: ")
		fmt.Scanln(&choice)

		def.Definition = defs[choice-1]
	}

	fmt.Printf("\n%s\n", def)

	return def
}

func (d Definition) String() string {
	return fmt.Sprintf("%s : %s", d.Word, d.Definition)
}

// Endpoints

// Dictionary.com: https://www.dictionary.com/browse/%s

// https://www.google.com/#q=define+term ??
