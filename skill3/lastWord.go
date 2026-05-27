package main
import "strings"

func LastWord(s string) string{
	s = strings.TrimSpace(s)
	if s == ""{
		return ""
	}
	parts := strings.Split(s, " ")

	return parts[len(parts) -1] 
}

// QUESTION 
// Write a function LastWord that takes a sentense and returns its last word followed by a \n.
// A word is a section of string delimited by spaces or by the start/end of the string.

func LastWord1 (s string) string{
	s = strings.TrimSpace(s)
	words := strings.Fields(s)

	if len(s) == 0{
		return "\n"
	}

	return words[len(words)-1] + "\n"
}