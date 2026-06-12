package main


// Instructions
// Write a function ConcatSlice() that takes two slices of integers as 
// arguments and returns the concatenation of the two slices.

func ConcatSlice1(slice1, slice2 []int) []int {
    return append(slice1, slice2...)
}


func ConcatSlice(slice1, slice2 []int) []int {
result := make([]int, 0, len(slice1)+len(slice2))

for _, v := range slice1 {
result = append(result, v)
}

for _, v := range slice2 {
result = append(result, v)
}

return result
}