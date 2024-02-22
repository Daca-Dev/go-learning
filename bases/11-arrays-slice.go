package main

import "fmt"

// GO DOC URL -> https://pkg.go.dev/?utm_source=godoc

func main() {
	// ARRAY: an array is inmutable and is more efficient, it can not grow
	var array [4]int // [dimesion]data_type
	fmt.Println(array)

	array[0] = 1
	array[1] = 3

	fmt.Println(array, len(array), cap(array))
	// len: items number 
	// cap: capacity of array

	// SLICE: an slice is mutable it can grow if needed
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// slicing: spearate elements
	fmt.Println(slice[0])
	fmt.Println(slice[2:4])
	fmt.Println(slice[:4])
	fmt.Println(slice[4:])

	// append
	slice = append(slice, 7)
	fmt.Println(slice)

	newSlice := []int{8,9,10}
	slice = append(slice, newSlice...) // the three points meanign unzip the slice
	fmt.Println(slice)

	otherSlice := []int{}
	fmt.Println("Empty slice:", otherSlice)

}
