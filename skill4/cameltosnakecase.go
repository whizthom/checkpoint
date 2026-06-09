package main

import "unicode"

// Instructions
// Write a function that converts a string from camelCase to snake_case.

// If the string is empty, return an empty string.
// If the string is not camelCase, return the string unchanged.
// If the string is camelCase, return the snake_case version of the string.
// For this exercise you need to know that camelCase has two different writing
// alternatives that will be accepted:

// lowerCamelCase
// UpperCamelCase
// Rules for writing in camelCase:

// The word does not end on a capitalized letter (CamelCasE).
// No two capitalized letters shall follow directly each other (CamelCAse).
// Numbers or punctuation are not allowed in the word anywhere (camelCase1).

func CamelToSnake1(s string) string {
	if len(s) == 0{
		return ""
	}

	word:= []rune(s)
	length := len(word)

	for i, ch:= range word{

		if !unicode.IsLetter(ch){
			return s
		}

		if i == length-1 && unicode.IsUpper(ch){
			return s
		}

		if i > 0 && unicode.IsUpper(ch) && unicode.IsUpper(word[i-1]){
			return s
		}
	}

	var result [] rune

		for i, ch:= range word {
			if unicode.IsUpper(ch){
				if i != 0 {
					result = append (result, '_')
				}
				result = append(result, ch)
			}else{
				result = append(result, ch)
			}
		}
	return string(result)
}

func isLetter(ch rune) bool {
    return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isUpper(ch rune) bool {
    return ch >= 'A' && ch <= 'Z'
}

func toLower(ch rune) rune {
    return ch + ('a' - 'A')
}

func CamelToSnake(s string) string {
    if len(s) == 0 {
        return ""
    }

    word := []rune(s)
    length := len(word)

    for i, ch := range word {
        // No letters, numbers or punctuation allowed
        if !isLetter(ch) {
            return s
        }
        // Word cannot end on a capital letter
        if i == length-1 && isUpper(ch) {
            return s
        }
        // No two consecutive capital letters
        if i > 0 && isUpper(ch) && isUpper(word[i-1]) {
            return s
        }
    }

    var result []rune
    for i, ch := range word {
        if isUpper(ch) {
            if i != 0 {
                result = append(result, '_')
            }
            result = append(result, toLower(ch))
        } else {
            result = append(result, ch)
        }
    }
    return string(result)
}

