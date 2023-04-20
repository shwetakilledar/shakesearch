package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"index/suffixarray"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/client9/misspell"
)

const ResultLength = 250

func main() {
	searcher := Searcher{}
	err := searcher.Load("completeworks.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/search", handleSearch(searcher))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("Listening on port %s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Searcher struct {
	CompleteWorks string
	SuffixArray   *suffixarray.Index
}

func handleSearch(searcher Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}

		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)

		checker := misspell.New()
		q := query[0]
		corrected, diff := checker.Replace(q)


		if len(diff) >0 {
			q = corrected
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(struct {
        Data string `json:"data"`
      }{
        Data: q,
      })
			return
		}

		results := searcher.Search(q)
		err := enc.Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			Data []string `json:"data"`
		}{
			Data: results,
		})
		return
	}
}

func (s *Searcher) Load(filename string) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Load: %w", err)
	}
	s.CompleteWorks = string(dat)
	s.SuffixArray = suffixarray.New(dat)
	return nil
}

func (s *Searcher) Search(query string) []string {
	var re = regexp.MustCompile(query)
	res := s.SuffixArray.FindAllIndex(re, -1)
	results := []string{}

	for _, pair := range res {
		wordStartsAt := pair[0]
		wordEndsAt := pair[1]

		// Index of starting position of word search
		startPos := wordStartsAt - ResultLength
		endPos := wordStartsAt + ResultLength

		// Update start position if word is the frist word in the document
		if startPos <= ResultLength {
			startPos= wordStartsAt
		}

		// Update end position if word is the last word in the document
		if endPos >= len(s.CompleteWorks) {
			endPos= wordEndsAt
		}

		results = append(results, s.CompleteWorks[startPos:endPos])
	}

	return results
}
