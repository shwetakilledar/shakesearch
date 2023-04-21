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
	"strconv"
	"strings"
)

var ResultLength= 250
// Setting AllowMaxDiffInSpelling to 1 to find closest matching word
const AllowMaxDiffInSpelling= 1

func main() {
	searcher := Searcher{}
	err := searcher.Load("completeworks.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./client/build"))
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
	WordDict      map[string]string
}

func (s *Searcher) Load(filename string) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Load: %w", err)
	}
	s.CompleteWorks = string(dat)
	s.SuffixArray = suffixarray.New(dat)

	// Processing file to create word dictionary to handle misspelled words
	str := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s.CompleteWorks, " ")
	words := strings.Split(str, " ")
	wordDict := make(map[string]string)

	for _, word := range words {
		wordLowerCase := strings.ToLower(word)
		wordLowerCase = strings.TrimSpace(wordLowerCase)
		if len(wordLowerCase) > 1 {
			_, ok := wordDict[wordLowerCase]
			if ok {
				continue
			} else {
				wordDict[wordLowerCase] = wordLowerCase
			}
		}
	}
	s.WordDict = wordDict

	return nil
}

func handleSearch(searcher Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Search Query Setup
		query, ok := r.URL.Query()["q"]
		searchQuery := query[0]
		searchQuery = strings.TrimSpace(searchQuery)

		// Pagination Setup
		page := r.URL.Query().Get("page")
		pageNum, pageNotOk := strconv.Atoi(page)

		// Handle bad requests
		if !ok || len(searchQuery) < 1 || pageNotOk != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}

		// Check for Misspelling(s) and suggest closest matching word(s) from the CompleteWorks
		search := strings.Split(searchQuery, " ")
		misspelledWords := Misspell(search, searcher.WordDict, AllowMaxDiffInSpelling)
		if len(misspelledWords) > 0 {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(struct {
				Data string `json:"data"`
			}{
				Data: strings.Join(misspelledWords, " "),
			})
			return
		}
		// END

		// Handle Search Start
		results := searcher.Search(searchQuery)
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}

		returnResult := results
		if len(results) > PageSize+1 {
			start, end := Paginate(pageNum, len(results))
			returnResult = results[start:end]
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			Data       []string `json:"data"`
			TotalCount int      `json:"totalCount"`
		}{
			Data:       returnResult,
			TotalCount: len(results),
		})
		return
		// Handle Search END
	}
}

func (s *Searcher) Search(query string) []string {
	// Hanlde case insensitity
	// build permutations of query string to search result
	var result []string
	resultp := &result
	set := make(map[string]int)
	Permute(query, "", resultp)

	// Search
	results := []string{}
	for _, q := range *resultp {
		var re = regexp.MustCompile(q)
		res := s.SuffixArray.FindAllIndex(re, -1)

		// Loop over found results
		for _, pair := range res {
			wordStartsAt := pair[0]
			wordEndsAt := pair[1]

			// Build pair key to check if we already have appended the result into resulting array
			pairKey := strconv.Itoa(wordStartsAt) + strconv.Itoa(wordEndsAt)
			_, ok := set[pairKey]
			// Check if we have handled appending pair into resulting array
			if ok {
				break
			}

			set[pairKey] = 1

			// Index of starting position of word search
			startPos := wordStartsAt - ResultLength
			endPos := wordStartsAt + ResultLength

			// Update start position if word is the frist word in the document
			if startPos <= ResultLength {
				startPos = wordStartsAt
			}

			// Update end position if word is the last word in the document
			if endPos >= len(s.CompleteWorks) {
				endPos = wordEndsAt
			}

			resultingString := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s.CompleteWorks[startPos:endPos], " ")
			resultingString = "..." + resultingString + "..."
			results = append(results, resultingString)
		}
	}

	return results
}
