package main

// Instructions
// Write a function that simulates the behavior of the Itoa function in Go.
//  Itoa transforms a number represented as an int in a number represented as a string.

// For this exercise the handling of the signs + or - does have to be taken into
//  account.

// Expected function
// func Itoa(n int) string {

// }

func Itoa(n int) string {
	if n == 0{
		return "0"
	}

	isNegative := false

	if n < 0 {
		isNegative = true
		n = -n
	}

	result := ""

	for n > 0 {
		digit:=n%10

		result = string(rune('0' + digit)) + result

		n/=10
	}

	if isNegative{
		result = "-" + result
	}

	return result
}