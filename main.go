package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type (
	Definition struct {
		Word       string `json:"word"`
		Definition string `json:"definition"`
	}
)

const (
	outFileName = "words.json"
	inFileName  = "words.txt"
)

var (
	definitions []Definition
)

func main() {
	inFile, err := os.Open(inFileName)
	if err != nil {
		log.Fatalf("opening words file: %v", err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		word := scanner.Text()
		fmt.Println(word)
		definitions = append(definitions, getDef(word))
	}

	outFile, err := os.Create(outFileName)
	if err != nil {
		log.Fatalf("creating out file: %v", err)
	}
	defer outFile.Close()

	data, err := json.MarshalIndent(definitions, "", " ")
	if err != nil {
		log.Fatalf("marshalling json: %v", err)
	}

	outFile.Write(data)
}

func getDef(word string) Definition {
	r, err := http.Get(fmt.Sprintf("https://www.dictionary.com/browse/%s", word))
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	def := Definition{
		Word: word,
	}

	// Load HTML
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	def.Definition = doc.Find("#top-definitions-section + section div:nth-of-type(2) div span").First().Clone().Children().Remove().End().Text()

	fmt.Println(def)
	return def
}

// https://www.google.com/#q=define+term

// Opts

// https://www.dictionary.com/browse/garrulous
