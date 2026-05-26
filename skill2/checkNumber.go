package main

// Write a function that takes a string as an argument and returns true if the string contains any number, otherwise return false.


func CheckNumber(arg string)bool {
 if len(arg)==0{
	return false
 }

 for _, ch := range arg{
	if ch >= '0' && ch <= '9'{
		return true
	}
 }
 return false
}
