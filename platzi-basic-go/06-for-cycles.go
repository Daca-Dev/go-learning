package main

import "fmt"

// GO DOC URL -> https://pkg.go.dev/?utm_source=godoc

func main() {
	// en go solo tenemos un ciclo iterativo que cumple con

	// For conditional
	for i := 0; i < 10; i++ {
		fmt.Println("Counter for conditional:", i)
	}

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println("counter for while:", counter)
		counter++
	}

	// For forever
	counterForever := 0
	for {
		fmt.Println("counter for forever:", counterForever)
		counterForever++
	}
}
