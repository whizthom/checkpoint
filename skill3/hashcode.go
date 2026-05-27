package main

// QUESTION:
// Write a function HashCode(s string) int that sums each character’s ASCII value.
// Example: 'A' = 65, 'B' = 66 → "AB" = 131.

func HashCode(s string)int{

	count:= 0

	for _, ch:= range s{
		count+=int(ch)
	}

	return count
}

// QUESTION
// Write a function called HashCode() that takes a string as an argument and returns a new hashed string.

// The hash equation is computed as follows:
// (ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.

// If the resulting character is unprintable add 33 to it.



func HashCode1(dec string) string {
	
	size := len(dec)
	result := []byte{}

	for _, ch := range dec{
		hash := (int(ch) + size) % 127

		if hash < 33{
			hash += 33
		}
		result = append(result, byte(hash))
	}
	return string(result)
}
