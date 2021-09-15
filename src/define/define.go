package define

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// These functions take a word parameter to search for and return a list of definitions found

type (
	DefineFunc func(string) ([]string, error)
)

// Dictionary.com
func DefineDictionaryCom(word string) ([]string, error) {
	url := fmt.Sprintf("https://www.dictionary.com/browse/%s", word)
	r, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET for '%s': %v", url, err)
	}
	defer r.Body.Close()

	defs := []string{}

	// Load HTML
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return nil, fmt.Errorf("goquery loading document: %v", err)
	}

	doc.Find("#top-definitions-section + section div:nth-of-type(2) div").Each(func(i int, s *goquery.Selection) {
		defs = append(defs, s.Text())
	})

	return defs, nil
}

// lexico.com
func DefineLexico(word string) ([]string, error) {
	url := fmt.Sprintf("https://www.lexico.com/en/definition/%s", word)
	r, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET for %s: %v", url, err)
	}
	defer r.Body.Close()

	defs := []string{}

	// Load HTML
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return nil, fmt.Errorf("goquery loading document: %v", err)
	}

	doc.Find(".ind.one-click-content").Each(func(i int, s *goquery.Selection) {
		defs = append(defs, s.Text())
	})

	return defs, nil
}

// dictionary.cambridge.org
func DefineCambridge(word string) ([]string, error) {
	url := fmt.Sprintf("https://dictionary.cambridge.org/us/dictionary/english/%s", word)
	r, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET for %s: %v", url, err)
	}
	defer r.Body.Close()

	defs := []string{}

	// Load HTML
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return nil, fmt.Errorf("goquery loading document: %v", err)
	}

	doc.Find(".def.ddef_d.db").Each(func(i int, s *goquery.Selection) {
		defs = append(defs, s.Text())
	})

	return defs, nil
}
