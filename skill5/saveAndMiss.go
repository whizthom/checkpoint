package main



// Write a function called SaveAndMiss() that takes a string and an int as an argument.
//  The function should move through the string in sets determined by the int, saving the first set, 
//  omitting the second, saving the third, and so on, in a 'save' and 'miss' fashion until the 
//  end of the string is reached. Return a string containing the saved characters.

// If the int is 0 or a negative number return the original string.

func SaveAndMiss1(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	runes := []rune(arg)
	length := len(runes)
	result := []rune{}
	save := true

	for i := 0; i < length; i += num {
		if save {
			end := i + num
			if end > length {
				end = length
			}
			result = append(result, runes[i:end]...)
		}
		save = !save
	}

	return string(result)
}

func SaveAndMiss(s string, num int) string {

	if num <= 0 {
		return s
	}

	runes:= []rune(s)
	length:= len(runes)
	result:=[]rune{}
	save:=true

	for i:=0; i < length; i+=3{

		if save{
			end:= i + num
			if end > length{
				end = length
			}
			result = append(result, runes[i:end]...)
		}
	}
	return string(result)
}