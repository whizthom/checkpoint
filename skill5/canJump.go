package main

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