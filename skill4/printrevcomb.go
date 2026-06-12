package main

import (
	"fmt"
)


// Instructions
// Write a program that prints in descending order on a single line all unique combinations of three different 
// digits so that the first digit is greater than the second and the second is greater than the third.

// These combinations are separated by a comma and a space.

// Usage
// Here is an incomplete output :

// $ go run . | cat -e
// 987, 986, 985, 984, 983, 982, 981, 980, 976, ..., 310, 210$
// $

func printRevComb() {

	for first:= '9'; first >= '2'; first--{
		for second:= first - 1; second >= '1'; second--{
			for third:= second - 1; third >= '0' ; third--{
				fmt.Printf("%c%c%c", first,second,third)
				if !(first=='2' && second=='1' && third == '0'){
					fmt.Printf(", ")
				}
			}
		}
	}
	fmt.Println()
}