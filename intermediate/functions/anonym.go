package main

import "fmt"

/*
Anonym functions are functions that doesn't have name and their use
only apply for functions that you are sure that only going to use once

! is helpful in Go rutines
*/

func main() {
	x := 5
	y := func() int {
		return x * 2
	}()
	fmt.Println(y)
}
