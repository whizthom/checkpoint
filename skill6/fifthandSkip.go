package main


// Write a function FifthAndSkip() that takes a string and returns another string. 
// The function separates every five characters of the string with a space and removes the sixth one.

// If there are spaces in the middle of a word, ignore them and get the first character after the spaces until you 
// reach a length of 5.
// If the string is less than 5 characters return Invalid Input followed by a newline \n.
// If the string is empty return a newline \n.

func FifthAndSkip1(str string) string {
	// "If the string is empty return a newline \n"
	if len(str) == 0 {
		return "\n"
	}

	// "If there are spaces in the middle of a word, ignore them"
	// strip all spaces into cleaned slice
	cleaned := []rune{}
	for _, ch := range str{
		if ch != ' ' {
			cleaned = append(cleaned, ch)
		}
	}

	// "If the string is less than 5 characters return Invalid Input followed by a newline"
	if len(cleaned) < 5 {
		return "Invalid Input\n"
	}

	result := []rune{}
	count := 0
	for count < len(cleaned) {
		// "separates every five characters of the string with a space"
		// take up to 5 characters
		for i := 0; i < 5 && count < len(cleaned); i++ {
			result = append(result, cleaned[count])
			count++
		}
		// "removes the sixth one"
		// skip the 6th character
		if count < len(cleaned) {
			count++
		}
		// add space separator only if more characters remain
		// prevents a trailing space at the end
		if count < len(cleaned) {
			result = append(result, ' ')
		}
	}

	return string(result) + "\n"
}

func FifthAndSkip(str string) string {

	if len(str) == 0 {
		return "\n"
	}

	newWord:= []rune{}

	for _, ch:= range str{

		if ch != ' '{
			newWord = append(newWord, ch)
		}
	}

	count:=0
	result:= []rune{}

	for count < len(newWord){
		
		for i:=0 ; i < 5 && count < len(newWord); i++{
			result = append(result, newWord[count])
			count++
		}

		if count < len(newWord){
			count++
		}

		if count < len(newWord){
			result = append(result, ' ')
		}

	}
	return string(result) + "\n"
}