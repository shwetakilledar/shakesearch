package main

import (
	"strings"
)

func Misspell(search []string, wordDict map[string]string, maxDistance int) []string {
	var results []string
	for _, query := range search {
		for word, value := range wordDict {
			q := strings.ToLower(query)
			q = strings.TrimSpace(q)
			_, ok := wordDict[q]

			if q == word || ok {
				break
			}

			distance := LevenshteinDistance(q, word)
			if distance > 0 && distance <= maxDistance {
				results = append(results, value)
			}
		}
	}
	return results
}
