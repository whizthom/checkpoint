package main

// concatalternate
// Instructions
// Write a function ConcatAlternate() that receives two slices of 
// an int as arguments and returns a new slice with the result of the 
// alternated values of each slice.

// The input slices can be of different lengths.
// The new slice should start with an element of the largest slice.
// If the slices are of equal length, the new slice should return the 
// elements of the first slice first and then the elements of the second slice.
func ConcatAlternate1(slice1, slice2 []int) []int {

    result := []int{}
    
    // Determine which slice is larger (or equal)
    var larger, smaller []int

    if len(slice1) >= len(slice2) {
        larger, smaller = slice1, slice2
    } else {
        larger, smaller = slice2, slice1
    }
    
    for i := 0; i < len(larger); i++ {
        result = append(result, larger[i])
        if i < len(smaller) {
            result = append(result, smaller[i])
        }
    }
    
    return result
}

func ConcatAlternate(slice1, slice2 []int) []int {

	result := []int{}

	var larger, smaller [] int

	if len(slice1) >= len(slice2) {
		larger, smaller = slice1, slice2
	}else{
		larger, smaller = slice2, slice1
	}

	for i:=0 ; i < len(larger) ; i++{
		result = append(result, larger[i]) 
		if i < len(smaller){
			result = append(result, smaller[i])
		}
	} 
	return result 
}
