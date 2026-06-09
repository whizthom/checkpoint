package main


// Can Jump
// Given an array of non-negative integers representing the number of steps you can take 
// forward from each position, 
// implement the function CanJump() which takes a slice of unsigned integers 
// []uint as input and returns a boolean value. This function should determine 
// if it's possible to reach and stay at the last index of the array starting from the first index, 
// based on the steps you need to advance. Be aware that:

// Each value represents the exact number of steps you must take forward from that position.
// The function should return true if it's possible to reach and stay at the last index without 
// stepping out of the array, and false otherwise.
// If the input has only one element, that is the last position in the array so the function
// will return true but if the array is empty it returns false.

func CanJump(arr []uint) bool {
	// Handle the edge case of an empty array
	if len(arr) == 0 {
		return false
	}
	// If the array has only one element, you are already at the end
	if len(arr) == 1 {
		return true
	}

	curr := 0
	lastIndex := len(arr) - 1

	// Move through the array following the exact steps
	for curr < lastIndex {
		steps := int(arr[curr])
		
		// If steps are 0, you're stuck and cannot move
		if steps == 0 {
			return false
		}
		
		curr += steps

		// Check if we stepped out of bounds
		if curr > lastIndex {
			return false
		}
	}

	// If we landed exactly on the last index, return true
	return curr == lastIndex
}