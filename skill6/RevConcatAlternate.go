package main


func RevConcatAlternate1(slice1, slice2 []int) []int {
	result := []int{}

	// reverse both slices
	rev1 := reverse1(slice1)
	rev2 := reverse1(slice2)

	len1 := len(rev1)
	len2 := len(rev2)

	// dump the extra elements of the larger slice first
	if len1 > len2 {
		for i := 0; i < len1-len2; i++ {
			result = append(result, rev1[i])
		}
	} else if len2 > len1 {
		for i := 0; i < len2-len1; i++ {
			result = append(result, rev2[i])
		}
	}

	// now alternate starting with slice1 for equal portion
	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	offset1 := len1 - minLen
	offset2 := len2 - minLen

	for i := 0; i < minLen; i++ {
		result = append(result, rev1[offset1+i])
		result = append(result, rev2[offset2+i])
	}

	return result
}

func reverse1(slice []int) []int {
	rev := []int{}
	for i := len(slice) - 1; i >= 0; i-- {
		rev = append(rev, slice[i])
	}
	return rev
}


// Write a function RevConcatAlternate() that receives two slices of int as arguments and returns a new slice with 
// alternated values of each slice in reverse order.

// The input slices can have different lengths.
// The new slice should start with the elements from the largest slice first and when they became equal size slices, 
// it should add an element of the first given slice.
// If the slices are of equal length, the new slice should start with an element of the first slice.


func RevConcatAlternate(slice1, slice2 []int) []int {
	result := []int{}

	// reverse both slices
	rev1 := reverse(slice1)
	rev2 := reverse(slice2)

	len1 := len(rev1)
	len2 := len(rev2)

	// dump the extra elements of the larger slice first
	if len1 > len2 {
		for i := 0; i < len1-len2; i++ {
			result = append(result, rev1[i])
		}
	} else if len2 > len1 {
		for i := 0; i < len2-len1; i++ {
			result = append(result, rev2[i])
		}
	}

	// find the shorter slice length for alternating
	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	// find where the equal portion starts in each reversed slice
	offset1 := len1 - minLen
	offset2 := len2 - minLen

	// alternate starting with slice1 for equal portion
	for i := 0; i < minLen; i++ {
		result = append(result, rev1[offset1+i])
		result = append(result, rev2[offset2+i])
	}

	return result
}

func reverse(slice []int) []int {
	rev := []int{}
	for i := len(slice) - 1; i >= 0; i-- {
		rev = append(rev, slice[i])
	}
	return rev
}
