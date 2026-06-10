package main

import (
	"fmt"
	"os"
)

// Write a program that takes two string and checks whether it is possible to write the first string with characters from the second string. 
// This rewrite must respect the order in which these characters appear in the second string.

// If it is possible, the program displays the string followed by a newline ('\n'), otherwise it simply displays nothing.

// If the number of arguments is different from 2, the program displays nothing.



func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := []rune(os.Args[1])
	s2 := []rune(os.Args[2])

	count := 0
	for i := 0; i < len(s2) && count < len(s1); i++ {
		if s2[i] == s1[count] {
			count++
		}
	}

	if count == len(s1) {
		fmt.Println(os.Args[1])
	}
}