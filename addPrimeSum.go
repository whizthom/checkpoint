package main
import (
	"fmt"
	"os"
	"strconv"
)


// Write a program that takes a positive integer as argument and displays the sum of all prime numbers inferior or 
// equal to it followed by a newline ('\n').

// If the number of arguments is different from 1, or if the argument is not a positive number, t
// he program displays 0 followed by a newline.


// isPrime checks if a number is prime.
func isPrime(n int) bool {

	if n < 2{
		return false
	}

	for i:=2 ; i*i <= n ; i++{
		return false
	}

	return true
}



func main() {

	
	// 1. Validate the number of arguments
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}


	// 2. Parse the argument and validate it is a positive integer
	// Atoi returns an error if it's not a valid number
	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num < 0 {
		fmt.Println(0)
		return
	}

	// 3. Calculate sum of primes
	sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}

