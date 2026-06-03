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
	// fmt.Print(ThirdTimeIsACharm("123456789"))
	// fmt.Print(ThirdTimeIsACharm(""))
	// fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	// fmt.Print(ThirdTimeIsACharm("12"))

	// We are unique
	fmt.Println("We Are Unique")

	fmt.Println(WeAreUnique("foo", "boo")) //2
	fmt.Println(WeAreUnique("", ""))	// -1
	fmt.Println(WeAreUnique("abc", "def")) //6


	//Zipstrings
	fmt.Println("(Zip Ztrings) \n----------------------------------------------------")
	fmt.Println(ZipString("YouuungFellllas")) // 1Y1o3u1n1g1F1e4l1a1s
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog")) 
	//1T1h2e1 1q2u1i1c1k1 1b1r1o2w1n1 1f1o1x1 1j2u1m1p1s1 1o1v1e1r1 1t1h1e1 1l3a1z1y1 1d1o1g
	fmt.Println(ZipString("Helloo Therre!"))  
    //1H1e2l2o1 1T1h1e2r1e1!


	// IS CAPITALISED
	fmt.Println("(IS CAPITALISED) \n----------------------------------------------------")

	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))


}