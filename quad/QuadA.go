package main

import "fmt"

// QuadA draws a rectangle with corners 'o', horizontal sides '-', and vertical sides '|'
func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			// Check for corners
			if (i == 0 || i == y-1) && (j == 0 || j == x-1) {
				fmt.Print("o")
			} else if i == 0 || i == y-1 {
				// Top and bottom edges
				fmt.Print("-")
			} else if j == 0 || j == x-1 {
				// Left and right edges
				fmt.Print("|")
			} else {
				// Interior
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	QuadA(6, 5) // You can change this
}