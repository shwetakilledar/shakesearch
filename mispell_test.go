package main

import (
	"strings"
	"testing"
)

type misspellTest struct {
	search   []string
	expected []string
}

var misspellTests = []misspellTest{
	{[]string{" "}, []string{}},
	// no missplleing
	{[]string{"hamlet"}, []string{}},
	// expect result as below words have 1 levenshtein distance
	{[]string{"hamlett"}, []string{"hamlet"}},
	{[]string{"hamlot"}, []string{"hamlet"}},
	{[]string{"kamlet"}, []string{"hamlet"}},
	// helmet is not present in wordDict; so nothing to lookup to to make decision
	{[]string{"helmet"}, []string{}},
	// expect no result as below words have more than 1 levenshtein distnace
	{[]string{"helmettt"}, []string{}},
	{[]string{"helmott"}, []string{}},
}

func TestMisspell(t *testing.T) {
	maxDistance := 1
	wordDict := map[string]string{"hamlet": "hamlet"}

	for _, test := range misspellTests {
		output := Misspell(test.search, wordDict, maxDistance)
		result := strings.Join(output, " ")
		expected := strings.Join(test.expected, " ")
		if result != expected {
			t.Errorf("Output %q not equal to expected %q", output, expected)
		}
	}
}
