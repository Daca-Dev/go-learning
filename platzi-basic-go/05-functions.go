package main

import "fmt"

// basic function
func printMessage(value string) {
	fmt.Printf("hello world %s", value)
}

// function withm more arguments
func moreArguments(a, b int, c string) {
	fmt.Println(a, b, c)
}

// one return value
func returnValue(a int) int {
	return a * 2
}

// multiple return values
func returnMultiplesValues(a int) (c, d int) {
	return a * 2, a * -2
}

func main() {
	printMessage("2")
	moreArguments(1,2, "Amazing")
	
	value := returnValue(2)
	fmt.Println(value)

	value1, value2 := returnMultiplesValues(4)
	println("value1:", value1)
	println("value2:", value2)
}
