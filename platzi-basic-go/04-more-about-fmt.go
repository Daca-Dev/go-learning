package main

import "fmt"

// FOR MORE INFORMATION ABOUT fmt -> https://pkg.go.dev/fmt

func main() {
	// var declaration
	helloMessage := "Hello"
	worldMessage := "World"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf: string with format
	name := "platzi"
	courses := 500

	fmt.Printf("%s have more than %d courses\n", name, courses)
	// if we dont know the type of data use %v
	//! a good practice is use the type of data if you know it

	// Sprintf
	message := fmt.Sprintf("%s have more than %d courses", name, courses)
	fmt.Println(message)

	// data type
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("worldMessage: %T\n", worldMessage)
}
