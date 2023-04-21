package main

import (
	"testing"
)

func TestMin(t *testing.T) {
	got := Min(1, 2, 3)
	want := 1
	if got != want {
		t.Errorf("Min function got: %q, but wanted: %q", got, want)
	}
}

type levenshteinDistanceTest struct {
	searchQuery string
	wordInDict  string
	expected    int
}

var levenshteinDistanceTests = []levenshteinDistanceTest{
	{"", "", 0},
	{"", "hamlet", 6},
	{"Hamet", "Hamlet", 1},
	{"hamettt", "Hamlet", 3},
}

func TestLevenshteinDistance(t *testing.T) {
	for _, test := range levenshteinDistanceTests {
		if output := LevenshteinDistance(test.searchQuery, test.wordInDict); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
