package main

// Write a function that takes a string and a character as arguments and returns the number of times the character appears in the string.
// If the character is not in the string return 0
// If the string is empty return 0


func countChar(str string, c rune)int{

	if len(str)==0{
		return 0
	}

	count := 0

	for _, ch := range str{
		
		if ch == c{
			count++
		}
	}
return count
}
