package main

// QUESTION:
// Write a function CountRepeats(s string) that counts how many characters
// appear more than once in the string.
// Example: "mississippi" → m=1, i=4, s=4, p=2 → 3 repeated characters (i, s, p)

func countRepeat(s string)int{

	if len(s) == 0{
		return 0
	}

	freq := make(map[rune]int)

	for _, ch := range s{
		freq[ch]++
	}

	// count how many runes where repeated
	repeated := 0;

	for _, count := range freq{
		if count > 1 {
			repeated++
		}
	}
return repeated
}


// // Let us try to search for how many times a specific character appears
// func CountRepeat(s string, r rune)int{
	 
// 	// freq := make(map[rune]int)

// 	count:= 0
// 	for _, ch := range s{
// 		if ch == r{
// 			count++
// 		}
// 	}
// 	return count
// }