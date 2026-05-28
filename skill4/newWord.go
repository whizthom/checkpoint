package main 

// Instructions
// Write a function IsCapitalized() that takes a string as an argument and returns true if
// each word in the string begins with either an uppercase letter or a non-alphabetic 
// character.

// If any of the words begin with a lowercase letter return false.
// If the string is empty return false.

func IsCapitalized(s string)bool{
	if len(s) == 0{
		return false
	}

	newWord := true

	for _, ch:= range s{
		if ch == ' '{
			newWord = true
		}else{
			if newWord{
				if ch >= 'a' && ch <= 'z'{
					return false
				}
			}
			newWord=false
		}
	}
	return true
}