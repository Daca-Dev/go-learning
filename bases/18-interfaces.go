package main

import (
	"fmt"
)

// GO DOC URL -> https://pkg.go.dev/?utm_source=godoc

// interfaces
// is like made one point of entry for many variables
type figure2D interface {
	area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base   float64
	height float64
}

func (s square) area() float64 {
	return s.base * s.base
}

func (r rectangle) area() float64 {
	return r.base * r.height
}

func calculate(f figure2D) {
	fmt.Println("Area:", f.area())
}

func main() {

	a := square{base: 2}
	fmt.Printf("the area is: %f\n", a.area())

	b := rectangle{base: 2, height: 3}
	fmt.Printf("the area is: %f\n", b.area())

	calculate(a)
	calculate(b)

	// interfaces list (this tip allow us to create a slice(list) of diferent types)
	myInterface := []interface{}{
		"Hola", 12, 3.12416,
	}
	fmt.Println(myInterface...)

}
