package main

import "fmt"

/*
In go we have many ways to declare a struct (class) and the use of them
depends only on you, all are valid
*/
type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// 1: delcare an empty struct
	e1 := Employee{}
	// 2: declare a new struct but passing some or all properties
	e2 := Employee{
		id:       2,
		vacation: true,
	}
	// 3: using the built in function new, this functions return a pointer to the new struct
	// and the values are empty
	e3 := new(Employee)
	// 4: using a function that take all the properties and return a pointer to the struct
	e4 := NewEmployee(4, "David Casas", true)
	fmt.Println("E1", e1)
	fmt.Println("E2", e2)
	fmt.Println("E3", e3)
	fmt.Println("E4", e4)

	// notice that the the fmt.Println function add a & before the object, is because is a pointer
}
