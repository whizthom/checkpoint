package main
import "strings"

func RepeatAlpha(s string) string{
	var result strings.Builder

	for _, ch:= range s{
		if ch >= 'A' && ch <= 'Z'{
			count := int(ch - 'A' + 1)
			result.WriteString(strings.Repeat(string(ch), count))
		}else if ch >= 'a' && ch <= 'z'{
			count := int(ch - 'a' + 1)
			result.WriteString(strings.Repeat(string(ch), count))
		}else{
			result.WriteRune(ch)
		}
	}
	return result.String()
}
		
		