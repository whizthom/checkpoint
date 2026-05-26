package main
import "fmt"

// Print "NO" if the input integar is NOT divisible by 5.

func printIfNotDivisibleBy5(n int){
	if n%5 != 0{
		fmt.Println("NO")
	}else{
		fmt.Println("YES")
	}
}


// Write a function that takes a string and returns:
//  1. "G\n" if the string's length is less than 3 (including empty string)
//  2. "Invalid Input\n" otherwise.

func PrintIfNot(str string) string{

	if len(str) < 3 {
		return "G\n"
	}

	return "Invalid Input\n"
}