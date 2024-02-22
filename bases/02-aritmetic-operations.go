package main

import "fmt"

func main() {
	// square area
	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("The area is ", squareArea)

	x := 10
	y := 50

	// suma
	result := x + y
	fmt.Println("add:", result)
	result = x - y
	fmt.Println("less:", result)
	result = x * y
	fmt.Println("product:", result)
	result = y / x
	fmt.Println("division:", result)
	result = x % y
	fmt.Println("module:", result)
	x++
	fmt.Println("incremental:", x)
	y--
	fmt.Println("decremental:", y)
}
