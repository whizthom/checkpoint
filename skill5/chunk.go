package main

import "fmt"

func Chunk(slice []int, size int) {
	if size <= 0 {
		fmt.Println()
		return
	}
	if len(slice) == 0 {
		fmt.Println("[]")
		return
	}

	var chunked [][]int

	// Use min() to safely handle the remaining elements on the final slice
	for i := 0; i < len(slice); i += size {
		chunked = append(chunked, slice[i:min(i+size, len(slice))])
	}

	fmt.Println(chunked)
}