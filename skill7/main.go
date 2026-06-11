package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n  WORD FLIP  ________________________________________\n")


	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}

