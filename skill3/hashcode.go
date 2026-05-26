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

