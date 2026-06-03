package main



// Instructions
// Write a function that takes two strings's and returns the number of characters that are not included in both, 
// without repeating characters.

// If there is no unique characters return 0.
// If both strings are empty return -1.


func WeAreUnique(str1, str2 string) int {

	if len(str1) == 0 || len(str2) == 0{
		return -1
	}

	box := make(map[rune]int)

	for _, ch:= range str1{
		box[ch] |= 1
	}

	for _, ch:= range str2{
		box[ch] |= 2
	}

	count := 0
	
	for _, ch:= range box{
		if ch == 1 || ch == 2{
			count++
		}
	}

	return count 
}