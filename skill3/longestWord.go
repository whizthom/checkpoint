package main

import (
	"strings"
)

func LongestWord(s string) string {
	// 1. Split the string into words
	words := strings.Fields(s)

	// 2. Handle empty input case immediately
	if len(words) == 0 {
		return ""
	}

	longest := ""

	// 3. Iterate through the words
	for _, word := range words {
		// Strictly greater than ensures we keep the first occurrence in a tie
		if len(word) > len(longest) {
			longest = word
		}
	}

	return longest
}