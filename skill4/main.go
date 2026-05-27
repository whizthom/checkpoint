package main
import "fmt"

func main(){
	// CAMEL to SNAKE
		fmt.Println(CamelToSnake("helloWorld")) // hello_world
		fmt.Println(CamelToSnake("HelloWorld"))
		fmt.Println(CamelToSnake("helloWorld"))
		fmt.Println(CamelToSnake("camelCase"))
		fmt.Println(CamelToSnake("CAMELtoSnackCASE"))
		fmt.Println(CamelToSnake("camelToSnakeCase"))
		fmt.Println(CamelToSnake("hey2"))
}