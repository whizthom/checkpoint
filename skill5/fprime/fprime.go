package main

import (
	"fmt"
	"os"
	"strconv"
)


// Write a program that takes a positive int and displays its prime factors, 
// followed by a newline ('\n').

// Factors must be displayed in ascending order and separated by *.

// If the number of arguments is different from 1, if the argument is invalid, or
//  if the integer does not have a prime factor, the program displays nothing.

func m() {
	if len(os.Args) != 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 1 {
		return
	}

	factors := []int{}
	for d := 2; d*d <= n; d++ {
		for n%d == 0 {
			factors = append(factors, d)
			n /= d
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}

	for i, f := range factors {
		if i > 0 {
			fmt.Print("*")
		}
		fmt.Print(f)
	}
	fmt.Println()
}

func main(){
	if len(os.Args) !=2{
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 2{
		return
	}

	factors := []int{}

	for i:=2; i*i <= n; i++{

		for n%i==0{
			factors = append(factors, i)
			n/=i
		}
	}
	if n > 1{
		factors = append(factors, n)
	}

	for i, f:= range factors{
		if i > 0{
			fmt.Print("*")
		}
		fmt.Print(f)
	}
	fmt.Println()
}


// $ go run . 225225
// 3*3*5*5*7*11*13
// $ go run . 8333325
// 3*3*5*5*7*11*13*37
// $ go run . 9539
// 9539
// $ go run . 804577
// 804577
// $ go run . 42
// 2*3*7
// $ go run . a
// $ go run . 0
// $ go run . 1
// $