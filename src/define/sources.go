package define

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type (
	DefineFunc func(string) ([]string, error)
)

// Dictionary.com
func DefineDictionaryCom(word string) ([]string, error) {
	r, err := http.Get(fmt.Sprintf("https://www.dictionary.com/browse/%s", word))
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	defs := []string{}

	// Load HTML
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#top-definitions-section + section div:nth-of-type(2) div").Each(func(i int, s *goquery.Selection) {
		defs = append(defs, s.Text())
	})

	return defs, nil
}

// lexico.com
func DefineLexico(word string) ([]string, error) {
	r, err := http.Get(fmt.Sprintf("https://www.lexico.com/en/definition/%s", word))
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	defs := []string{}

	// Load HTML
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".ind.one-click-content").Each(func(i int, s *goquery.Selection) {
		defs = append(defs, s.Text())
	})

	return defs, nil
}

// dictionary.cambridge.org
func DefineCambridge(word string) ([]string, error) {
	r, err := http.Get(fmt.Sprintf("https://dictionary.cambridge.org/us/dictionary/english/%s", word))
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	defs := []string{}

	// Load HTML
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".def.ddef_d.db").Each(func(i int, s *goquery.Selection) {
		defs = append(defs, s.Text())
	})

	return defs, nil
}
