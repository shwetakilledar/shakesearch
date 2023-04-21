package main

import "strings"

var PageSize = 20

func Paginate(pageNum int, sliceLength int) (int, int) {
	start := pageNum * PageSize

	if start > sliceLength {
		start = sliceLength
	}

	end := start + PageSize
	if end > sliceLength {
		end = sliceLength
	}

	return start, end
}

func Permute(word string, perOfWord string, resultp *[]string) {
	if len(word) == 0 {
		*resultp = append(*resultp, perOfWord)
		return
	}

	ch := string(word[0])
	ch2 := string(word[0])
	ch = strings.ToLower(ch)
	ch2 = strings.ToUpper(ch2)
	word = word[1:]

	Permute(word, perOfWord+ch, resultp)
	Permute(word, perOfWord+ch2, resultp)
}

func Min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	} else {
		return c
	}
}

func LevenshteinDistance(search, wordInDict string) int {
	search = strings.ToLower(search)
	wordInDict = strings.ToLower(wordInDict)
	m := len(search)
	n := len(wordInDict)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			if search[i-1] == wordInDict[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(dp[i-1][j-1]+1, dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}

	return dp[m][n]
}
