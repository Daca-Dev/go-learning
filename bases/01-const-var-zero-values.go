package main

import "fmt"

func main() {
	// consts declaration
	const pi float64 = 3.1416
	const pi2 = 3.14

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// variables declaration
	base := 12 // the := indicated that is the first time that i declare de variable
	var altura int = 14 // here declara a variable and indicate the value
	var area int // declarate but not initialice the value

	fmt.Println("base:", base)
	fmt.Println("altura:", altura)
	fmt.Println("area:", area) // how area is declared without value it use the respective zero value

	// Zero values
	// the zero values is the default values asign by go if a variable is not initilized
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a,b,c,d)

	// square area
	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("The area is ", squareArea)
}
