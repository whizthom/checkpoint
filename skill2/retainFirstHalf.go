package main

// Given a slice, return only the first half of it 

func retainFirstHalf[T any](s[]T)[]T{
	half := len(s)/2
	return s [:half]
}

// write a function called RetainFirstHalf() that takes a string as an argument and returns the first half of this string.
// *if the length of the strinf is odd, round it down.
// *if the string is empty, return an empty string.
// * if the string length is equal to one, return the string.


func RetainFirstHalf(str string) string {


	length := len(str)

	// if the string is empty, return an empty string
	if length == 0{
		return ""
	}

	if len(str)==1{
		return str
	}

	// if the length of the string is odd, round it down
	
	half := length/2
	 
   return str[:half]
}