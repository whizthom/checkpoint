package main
import "unicode"


func CamelToSnake(str string) string{
	
	if str == ""{
		return ""
	}

	var result []rune
	runes:= []rune(str)

	for i, ch:= range runes{
		if (!unicode.IsLetter(ch)) ||
		(i > 0 && unicode.IsUpper(ch) && unicode.IsUpper(runes[i-1])) ||
		( i==len(runes)- 1 && unicode.IsUpper(ch)){
			return str
		}
	}


	return string(result)
}