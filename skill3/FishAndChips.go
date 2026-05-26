package main
import "strings"

// QUESTION:
// Replace every occurrence of "fish" with "chips" inside a string.

func FishAndChips(s string)string{
		return strings.ReplaceAll(s,"fish", "chips")
}


// QUESTION:
// Write a function FishAndChips() that takes an int and returns a string.
// * If the number is divisible by 2, print fish.
// * If the number is divisible by 3, print chips.
// ^ If the number is divisibke by 2 and 3, print fish and chips.
// * If the number is negative return error : number is negative.
// * If the number is non divisble by 2 or 3 return error: non divisible.


func FishAndChips1(n int) string {
    if n < 0 {
        return "error : number is negative"
    }
    if n%2 == 0 && n%3 == 0 {
        return "fish and chips"
    }
    if n%2 == 0 {
        return "fish"
    }
    if n%3 == 0 {
        return "chips"
    }
    return "error: non divisible"
}