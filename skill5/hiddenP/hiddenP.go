package main

import (
	"fmt"
	"os"
)

// Write a program named hiddenp that takes two strings as arguments. 
// The program should check if the first string s1 is hidden in the second s2.
// s1 is considered hidden in s2 if it is possible to find each character from s1 in s2, 
// in the same order as they appear in s1, but not necessarily consecutively.

// If s1 is hidden in s2, the program should display 1 followed by a newline.
// If s1 is not hidden in s2, the program should display 0 followed by a newline.
// If s1 is an empty string, it is considered hidden in any string.
// If the number of arguments is different from 2, the program should display nothing.


func m() {
	if len(os.Args) != 3 {
		return
	}

	s1 := []rune(os.Args[1])
	s2 := []rune(os.Args[2])

	count := 0
	for i := 0; i < len(s2) && count < len(s1); i++ {
		if s2[i] == s1[count] {
			count++
		}
	}

	if count == len(s1) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

// $ go run . "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
// 1$
// $ go run . "abc" "2altrb53c.sse" | cat -e
// 1$
// $ go run . "abc" "btarc" | cat -e
// 0$
// $ go run . "DD" "DABC" | cat -e
// 0$
// $ go run .
// $

// MY OWN TRY
func main(){
	if len(os.Args) !=3 {
		return
	}

	s1 := []rune(os.Args[1])
	s2 := []rune(os.Args[2])

	count:=0

	for i:=0 ; i < len(s2) && count < len(s1); i++{
		if s2[i] == s1[count]{
			count++
		}
	}

	if count == len(s1){
		fmt.Print(1)
	}else{
		fmt.Print(0)
	}
	fmt.Println()
}