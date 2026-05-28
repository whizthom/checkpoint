package main

import "fmt"

// Instructions
// Write a function that takes (arr [10]byte), and displays the memory as in the example.

// After displaying the memory the function must display all the ASCII graphic characters. 
// The non printable characters must be replaced by a dot.

// The ASCII graphic characters are any characters intended to be written, printed, or otherwise displayed in a 
// form that can be read by humans, present on the ASCII encoding.


func PrintMemory(arr [10]byte) {

	for i :=0; i<10 ; i++{
		fmt.Printf("%02x", arr[i])

		if (i+1)%4 == 0 || i == 9{
			fmt.Println()
		}else{
			fmt.Print(" ")
		}
	}

	for _, b := range arr {
		if b >= 32 && b <= 126{
			fmt.Printf("%0c", b)
		}else{
			fmt.Print(".")
		}
	}
	fmt.Println()
}