package main

import "fmt"

// Write a function called Chunk that receives as parameters a slice, slice []int, 
// and a number size int. The goal of this function is to chunk a slice into many sub slices 
// where each sub slice has the length of size.

// If the size is 0 it should print a newline ('\n').

func Chunk(slice []int, size int) {
	
	if size <= 0{
		fmt.Println()
		return
	}

	if len(slice) == 0{
		fmt.Println("[]")
		return
	}

	var chunk [][] int

	for i:=0; i < len(slice); i+=size{
		chunk = append(chunk, slice[i : min(i+size, len(slice))])
	}

	fmt.Println(chunk)
}