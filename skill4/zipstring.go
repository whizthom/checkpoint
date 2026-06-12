package main

import (
	"strconv"
)



// Instructions
// Write a function that takes a string and returns a new string that replaces every 
// character with the number of duplicates and the character itself, 
// deleting the extra duplications.

// The letters are from the latin alphabet list only. Any other character, symbols, shall not be tested.

func ZipString(s string) string {
	
	if len(s)==0{
		return ""
	}

	res := ""
	count := 1

	for i:=1; i <= len(s); i++{
		
		if i == len(s) || s[i] != s[i-1]{
			res += strconv.Itoa(count) + string(s[i-1])
			count = 1
		}else{
			count++
		}
	}

	return res
}
