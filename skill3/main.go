package main
import "fmt"

func main() {

	// COUNT REPEAT
	// fmt.Println(countRepeat("mississippi")) // 3
	// fmt.Println(countRepeat("abcd"))        // 0

	// fmt.Println(CountRepeat("abaacd", 'a')) 

	// DIGIT LENGTH
	// fmt.Println(DigitLen(12345)) // 5
	// fmt.Println(DigitLen(-900))  // 3

	//Digitlen
	// fmt.Println(DigitLen2(100, 10)) 
	// fmt.Println(DigitLen2(100, 2)) 
	// fmt.Println(DigitLen2(-100, 16)) 
	// fmt.Println(DigitLen2(100, -1)) 

	//First Word Slice

	fmt.Println("\n FIRST WORD________________________________________\n")

	fmt.Println(firstWord("hello world"))     // "hello"
	fmt.Println(firstWord("   golang rocks")) // "golang"

	// Fish and Chips
	// fmt.Println(FishAndChips("I love fish and fish soup")) // Output: "I love chips and chips soup"

	// fmt.Println(FishAndChips1(4))
	// fmt.Println(FishAndChips1(9))
	// fmt.Println(FishAndChips1(6))


	// GCD
	// fmt.Println(GCD(48, 18)) // 6
	// fmt.Println(GCD(20, 8))  // 4

	// fmt.Println(Gcd(42, 10)) // 2
	// fmt.Println(Gcd(42, 12)) // 6
	// fmt.Println(Gcd(14, 77)) // 7
	// fmt.Println(Gcd(17, 3))   // 1


	// HASH CODE

	fmt.Println("\n HASH CODE________________________________________\n")

	fmt.Println(HashCode1("A"))           // Output: B
	fmt.Println(HashCode1("AB"))          // Output: CD
	fmt.Println(HashCode1("BAC"))         // Output: EDF
	fmt.Println(HashCode1("Hello World")) // Output: Spwwz+bz}wo

	// Last Word
	fmt.Println("\n LAST WORD________________________________________\n")

	fmt.Println(LastWord1("hello world")) // "world"
	fmt.Println(LastWord1("  learning go ")) // "go"

	// Longest Word
	// fmt.Println(LongestWord("go is expressive and powerful"))
	// fmt.Println(LongestWord("one two three"))

	fmt.Println("\n REPEAT ALPHA________________________________________\n")
	
	fmt.Println(RepeatAlpha("abc")) // a bb ccc
	fmt.Println(RepeatAlpha("A!c"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))

	// // WORD ANATOMY
	// fmt.Println(WordAnatomy("golang"))
	// fmt.Println(WordAnatomy("hello"))

}