package main
import "strings"


func firstWord (s string) string {

	s = strings.TrimSpace(s)

	if s == ""{
		return ""
	}

	parts := strings.SplitN(s, " ", 2)
	return parts[0]

}