package main
import "fmt"

// Instructions
// Write a function that takes two integers and returns a string showing the 
// range of numbers from the first to the second.

// The numbers must be separated by a comma and a space.
// If any of the arguments is bigger than 99 or less than 0, the function returns 
// Invalid followed by a newline \n.
// Prepend a 0 to any number that is less than 10.
// Add a new line \n at the end of the string.
// Expected function
// func FromTo(from int, to int) string {

func FromTo(from int, to int) string{

	if from > 99 || from < 0 || to > 99 || to < 0{
		return "Invalid\n"
	}

	result:=fmt.Sprintf("%02d",from)

	direction:=1
	if from > to {
		direction = -1
	}

	for i:= from + direction; i != to + direction; i+=direction{
		result+=fmt.Sprintf(" ,%02d",i)
	}
	return result + "\n"
}