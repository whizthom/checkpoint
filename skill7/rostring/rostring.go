package main

import (
	"fmt"
	"os"
)

// Instructions
// Write a program that takes a string and displays this string after rotating it one word to the left.

// Thus, the first word becomes the last, and others stay in the same order.

// A word is a sequence of alphanumerical characters.

// Words will be separated by only one space in the output.

// If the number of arguments is different from 1, the program displays a newline.



func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	s := os.Args[1]

	// split into words manually
	words := []string{}
	current := ""
	for _, ch := range s {
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

	// print words rotated one to the left
	// first word goes to the end
	for i := 1; i < len(words); i++ {
		if i > 1 {
			fmt.Print(" ")
		}
		fmt.Print(words[i])
	}
	if len(words) > 1 {
		fmt.Print(" ")
	}
	if len(words) > 0 {
		fmt.Print(words[0])
	}
	fmt.Println()
}


// $ go run . "abc   " | cat -e
// abc$
// $ go run . "Let there     be light"
// there be light Let
// $ go run . "     AkjhZ zLKIJz , 23y"
// zLKIJz , 23y AkjhZ
// $ go run . | cat -e
// $
// $