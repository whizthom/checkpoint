package main
import "strings"


// QUESTION:
// Given a sentence, return the longest word.
// If tie, return the first longest.

func LongestWord(s string) string{

	s = strings.TrimSpace(s)

	words := strings.Fields(s)

	longest := words[0]

	if len(words) == 0{
			return ""
	}
	 
	for _, ch := range words{

		if len(ch) > len(longest){
			longest = ch
		}
	}
	return longest
}
