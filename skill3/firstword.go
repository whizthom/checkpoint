package main
import "strings"


// Write a function that takes a string and return a string containing its first word,
//  followed by a newline ('\n').

// A word is a sequence of characters delimited by spaces or by the start/end of the argument.

func firstWord1 (s string) string {

	s = strings.TrimSpace(s)

	if s == ""{
		return ""
	}

	parts := strings.SplitN(s, " ",2)

	return parts[0]
}

func firstWord (s string) string {

	result := ""
	inWord := false

	for i:=0; i < len(s); i++{

		if s[i] != ' '{
			result = result + string(s[i])
			inWord = true
		}else if inWord{
			break
		}
	}
	return result + "\n"
}

