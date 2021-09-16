package define

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// These functions take a word parameter to search for and return a list of definitions found

type (
	Endpoint struct {
		QueryURL func(string) string // Returns the url to search for a word
		GetDefs  ScrapeDefsFunc      // Returns the list of definitions
	}

	// ScrapeDefsFunc will take a goquery document and scrape the definitions off of it
	ScrapeDefsFunc func(*goquery.Document) []string
)

var (
	DictionaryCom = Endpoint{
		QueryURL: urlDictionaryCom,
		GetDefs:  scrapeDictionaryCom,
	}

	Lexico = Endpoint{
		QueryURL: urlLexico,
		GetDefs:  scrapeLexico,
	}

	Cambridge = Endpoint{
		QueryURL: urlCambridge,
		GetDefs:  scrapeCambridge,
	}
)

func (e Endpoint) Define(word string) ([]string, error) {
	url := e.QueryURL(word)
	r, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET for %q: %v", url, err)
	}
	defer r.Body.Close()

	// Load HTML
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return nil, fmt.Errorf("goquery loading document: %v", err)
	}

	return e.GetDefs(doc), nil
}

// Endpoint specifics

// Dictionary.com
func urlDictionaryCom(word string) string {
	return fmt.Sprintf("https://www.dictionary.com/browse/%s", word)
}

func scrapeDictionaryCom(doc *goquery.Document) []string {
	defs := []string{}

	doc.Find("#top-definitions-section + section > div:nth-of-type(2) > div > span").Each(func(i int, s *goquery.Selection) {
		html, _ := s.Html()

		// Cut inner html after first text
		// Find first tag after text
		cutIdx := strings.Index(html, "<")
		var txt string
		if cutIdx != -1 {
			txt = html[:cutIdx-2]
		} else {
			txt = html
		}
		defs = append(defs, txt)
	})

	return defs
}

// lexico.com
func urlLexico(word string) string {
	return fmt.Sprintf("https://www.lexico.com/en/definition/%s", word)
}

func scrapeLexico(doc *goquery.Document) []string {
	defs := []string{}

	doc.Find(".ind.one-click-content").Each(func(i int, s *goquery.Selection) {
		defs = append(defs, s.Text())
	})

	return defs
}

// dictionary.cambridge.org
func urlCambridge(word string) string {
	return fmt.Sprintf("https://dictionary.cambridge.org/us/dictionary/english/%s", word)
}

func scrapeCambridge(doc *goquery.Document) []string {
	defs := []string{}

	doc.Find(".def.ddef_d.db").Each(func(i int, s *goquery.Selection) {
		txt := s.Text()
		defs = append(defs, txt[:len(txt)-2])
	})

	return defs
}
