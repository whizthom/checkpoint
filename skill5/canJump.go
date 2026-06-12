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

	if len(arr) == 0{
		return false
	}

	if len(arr) == 1{
		return true
	}

	count:= 0
	lastIndex:= len(arr) -1


	for count < lastIndex{
		steps:= int(arr[count])

		if steps == 0{
			return false
		}

		count += steps

		if count > lastIndex{
			return false
		}
	}

	return  count == lastIndex
}