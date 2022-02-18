package main

import (
	"fmt"
)

// GO DOC URL -> https://pkg.go.dev/?utm_source=godoc


func main() {
	// maps: is sinonimus a Dictionary in Python
	// maps ar unordered
	m := make(map[string]int)

	m["David"] = 26
	m["Daniela"] = 25
	m["Andres"] = 25

	fmt.Println(m)

	// route a map
	for key, value := range(m) {
		fmt.Println(key, value)
	}

	// found a value
	value, ok := m["David"]
	fmt.Println(value, ok)

	value2, ok := m["Elon"]
	fmt.Println(value2, ok)
	// the ok value is to validated if a key exist or not in the map
	// you can create a set (set of python) using a map with where key is the value and
	// type is boolean with a value of true
}
