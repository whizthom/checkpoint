package main

import (
	"fmt"
	"os"
)

// Instructions
// Write a program that takes one or more arguments and that, for each argument, 
// puts the last letter of each word in uppercase and the rest in lowercase.
//  It displays the result followed by a newline ('\n').

// If there are no argument, the program displays nothing.

func m() {
	if len(os.Args) < 2 {
		return
	}

	for _, arg := range os.Args[1:] {
		fmt.Println(reverseStrCap1(arg))
	}
}

func reverseStrCap1(s string) string {
	words := []rune(s)
	length := len(words)

	for i := 0; i < length; i++ {
		if isLetter(words[i]) {
			// check if it's the last letter of the word
			if i == length-1 || !isLetter(words[i+1]) {
				words[i] = toUpper(words[i])
			} else {
				words[i] = toLower(words[i])
			}
		}
	}
	return string(words)
}

func isLetter1(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func toUpper1(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - ('a' - 'A')
	}
	return ch
}

func toLower1(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + ('a' - 'A')
	}
	return ch
}





func main(){
	if len(os.Args) < 2{
		return
	}

	for _, args:= range os.Args[1:]{
		fmt.Println(reverseStrCap(args))
	}
}

	func reverseStrCap(s string)string{
		
		word:= []rune(s)
		length:= len(word)

		for i:=0; i < length; i++{
			if isLetter(word[i]){
				if i == length-1 || !isLetter(word[i+1]){
					word[i] = toUpper(word[i])
				}else{
					word[i] = toLower(word[i])
				}
			}
		}
		return string(word)
	}



func isLetter(ch rune)bool{

	if ch >= 'a' && ch <='z' || ch >='A' && ch <= 'Z'{
		return true
	}

	return false
}


func toUpper(ch rune)rune{
	if ch >= 'a' && ch <= 'z'{
		return ch - ('a' - 'A')
	}

	return ch
}

func toLower(ch rune)rune{
	if ch>= 'A' && ch <= 'Z'{
		return ch + ('a' - 'A')
	}
	return ch
}