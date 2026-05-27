package main

import (
	"fmt"
	"strings"
)

// QUESTION:
// Write a function WordAnatomy(word string) that returns:
// - the word
// - its length
// - its first letter
// - its last letter
// All in a formatted string.

func WordAnatomy(word string) string{

	word = strings.TrimSpace(word)

	if len(word) == 0{
		return "empty word"
	}

	length:=len(word)
	first:=word[0]
	last:=word[length-1]

	return fmt.Sprintf("word: %s | length: %d | first: %c | last: %c", 
	word, length, first, last)

}