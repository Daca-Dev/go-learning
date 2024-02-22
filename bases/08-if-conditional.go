package main

import "fmt"

// GO DOC URL -> https://pkg.go.dev/?utm_source=godoc

func main() {
	value1 := 1
	value2 := 2

	if value1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// with logic operators
	if value1 == 1 && value2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("Es falso")
	}

	if value1 == 0 || value2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("No es verdad OR")
	}
}
