package main 

import (
    "fmt"                       
    "os"                        
)


// Write a program that takes two string and displays, without doubles, the characters that appear in
//  both string, in the order they appear in the first one.

// The display will be followed by a newline ('\n').

// If the number of arguments is different from 2, the program displays nothing.


func m() {
    if len(os.Args) != 3 {      // os.Args[0]=program, os.Args[1]=s1, os.Args[2]=s2, so we need exactly 3
        return                  // wrong number of args, print nothing and exit
    }

    s1 := []rune(os.Args[1])   // convert first string to rune slice (to handle characters safely)
    s2 := []rune(os.Args[2])   // convert second string to rune slice

    seen := [256]bool{}   // array of 256 booleans all set to false, tracks already printed characters

    for _, ch := range s1 {  // loop through every character in s1, _ discards the index we dont need it
        if !seen[ch] && contain(s2, ch) {  // if ch hasnt been printed yet AND ch exists in s2
            seen[ch] = true                 // mark ch as printed to prevent duplicates
            fmt.Printf("%c", ch)            // print the character
        }
    }
    fmt.Println()               // print newline at the end
}

func contain(s []rune, ch rune) bool {  // helper function, checks if ch exists anywhere in s
    for _, c := range s {               // loop through every character in s
        if c == ch {                    // if we find a match
            return true                 // ch is in s
        }
    }
    return false                        // ch was never found in s
}


// $ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
// padinto
// $ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
// df6ewg4
// $

func main(){

	if len(os.Args) != 3{
		return
	}

	s1 := []rune(os.Args[1])
	s2 := []rune(os.Args[2])

	seen := [256]bool{}

	for _, ch:= range s1{

		if !seen[ch] && contains(s2, ch){
			seen[ch] = true
			fmt.Printf("%c", ch)
		}
	}
	fmt.Println()
}

func contains(s []rune , c rune)bool{
	for _, ch:= range s{
		if ch == c{
			return true
		}
	}
	return false
}