package main


// Instructions
// Write a function ConcatSlice() that takes two slices of integers as 
// arguments and returns the concatenation of the two slices.

func ConcatSlice(slice1, slice2 []int) []int {
    return append(slice1, slice2...)
}