package main

import "fmt"

// Print "YES" Only if an input is greater than 10
func printIfGreaterThan10(n int) {
    if n > 10 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

// write a function that takes a string as an argument and returns the letter G followed by a newline \n
//If the argument length is more or equal than 3, otherwise returns Invalid Input followed by a newline \n
// If its an empty string return G followed by a newline \n.


func PrintIf(str string) string {
	// If length is 3 or more, return "Invalid Input" + newline
	if len(str) >= 3 {
		return "Invalid Input\n"
	}

	// For empty strings or strings with length 1 or 2, return "G" + newline
	return "G\n"
}