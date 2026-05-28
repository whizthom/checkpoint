package main
import "fmt"

func main(){
	// CAMEL to SNAKE
		// fmt.Println(CamelToSnake("helloWorld")) // hello_world
		// fmt.Println(CamelToSnake("HelloWorld"))
		// fmt.Println(CamelToSnake("helloWorld"))
		// fmt.Println(CamelToSnake("camelCase"))
		// fmt.Println(CamelToSnake("CAMELtoSnackCASE"))
		// fmt.Println(CamelToSnake("camelToSnakeCase"))
		// fmt.Println(CamelToSnake("hey2"))


		// FindPrevPrime
	// 	tests := []int{5, 4, 10, 17, 1, 2, 100, -5}

	// for _, t := range tests {
	// 	fmt.Printf("FindPrevPrime(%d) → %d\n", t, FindPrevPrime(t))
	// }

	// FROM TO
	// fmt.Print(FromTo(1, 10))
	// fmt.Print(FromTo(10, 1))
	// fmt.Print(FromTo(10, 10))
	// fmt.Print(FromTo(100, 10))


	// Capitalized
	// fmt.Println(IsCapitalized("Hello! How are you?"))
	// fmt.Println(IsCapitalized("Hello How Are You"))
	// fmt.Println(IsCapitalized("Whats 4this 100K?"))
	// fmt.Println(IsCapitalized("Whatsthis4"))
	// fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	// fmt.Println(IsCapitalized(""))

	//

	// IOTA
	// fmt.Println(Itoa(12345))
    // fmt.Println(Itoa(0))
    // fmt.Println(Itoa(-1234))
    // fmt.Println(Itoa(987654321))

	// // Print Memory 
	// PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})

	// ThirdTimeIsACharm
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))

}