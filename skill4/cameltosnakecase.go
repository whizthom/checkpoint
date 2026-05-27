package main
import "unicode"


func CamelToSnake(str string) string{
if str == "" {
		return ""
	}

	var result []rune
	runes := []rune(str)

	// Single pass validation and transformation
	for i, ch := range runes {
		// 1. Validation: Reject invalid chars, consecutive capitals, or trailing capitals
		if !unicode.IsLetter(ch) || 
		   (i > 0 && unicode.IsUpper(ch) && unicode.IsUpper(runes[i-1])) ||
		   (i == len(runes)-1 && unicode.IsUpper(ch)) {
			return str
		}

		// 2. Transformation: Convert to snake_case
		if unicode.IsUpper(ch) {
			if i != 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(ch))
		} else {
			result = append(result, ch)
		}
	}

	return string(result)
}