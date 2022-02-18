package main

import "fmt"

// GO DOC URL -> https://pkg.go.dev/?utm_source=godoc

func main() {
	// DEFER
	// defer -> this keywork define functions or code that going
	// to be executed before the end or kill the function or block scope
	// an example of use is for example to close a data base connection
	fmt.Println("Hello")
	fmt.Println("World")

	defer fmt.Println("Hello")
	fmt.Println("World")

	// CONTINUE AND BREAK (Bboth of them are used in For )
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		} else if i == 7 {
			break
		}
			fmt.Println("Counter:", i)
	}
}
