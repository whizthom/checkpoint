package main

// write a function CountAlpha() that takes a string as an argument and returns the number of alphabet characters in the string

func CountAlpha(s string)int{
	count := 0

	for _, ch := range s{
		if (ch >= 'a' && ch <= 'z') || (ch >='A' && ch <= 'Z'){
		 count ++ 
		}
	}

	return count
}
	