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
		tests := []int{5, 4, 10, 17, 1, 2, 100, -5}

	for _, t := range tests {
		fmt.Printf("FindPrevPrime(%d) → %d\n", t, FindPrevPrime(t))
	}
}