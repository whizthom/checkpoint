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
	// fmt.Println(firstWord("hello world"))     // "hello"
	// fmt.Println(firstWord("   golang rocks")) // "golang"

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

	fmt.Println(HashCode("ABC")) // 198
	fmt.Println(HashCode("hello"))


	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))

	fmt.Println(HashCode("A"))           // Output: B
	fmt.Println(HashCode("AB"))          // Output: CD
	fmt.Println(HashCode("BAC"))         // Output: EDF
	fmt.Println(HashCode("Hello World")) // Output: Spwwz+bz}wo

}