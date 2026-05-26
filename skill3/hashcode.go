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



func HashCode(dec string) string {
	size := len(dec)
	result := []byte{} // Create a slice to hold our new characters

	for _, char := range dec {
		// 1. Calculate the hash using the formula
		hash := (int(char) + size) % 127

		// 2. Check if it's unprintable (less than 33)
		if hash < 33 {
			hash += 33
		}

		// 3. Convert back to a character and add to result
		result = append(result, byte(hash))
	}

	return string(result)
}
