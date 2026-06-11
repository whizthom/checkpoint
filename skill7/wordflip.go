package main


// Instructions
// Write a function WordFlip() that takes a string as input and returns it in reverse order.

// The output should be followed by a newline \n.
// If the string is empty, return Invalid Output.
// Ignore multiple spaces between words and trim any leading or trailing spaces in the string.



func WordFlip(str string) string {
	// empty string only → Invalid Output
	if len(str) == 0 {
		return "Invalid Output"
	}

	// collect words ignoring multiple spaces
	words := []string{}
	current := ""
	for _, ch := range str {
		if ch == ' ' {
			if len(current) > 0 {
				words = append(words, current)
				current = ""
			}
		} else {
			current += string(ch)
		}
	}
	if len(current) > 0 {
		words = append(words, current)
	}

	// only spaces → return newline
	if len(words) == 0 {
		return "\n"
	}

	// build result in reverse
	result := ""
	for i := len(words) - 1; i >= 0; i-- {
		if i != len(words)-1 {
			result += " "
		}
		result += words[i]
	}

	return result + "\n"
}