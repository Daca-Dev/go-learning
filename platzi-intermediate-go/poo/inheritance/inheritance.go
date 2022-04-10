package main

import "fmt"

/*
the inheritance in Go is apply using the concept of composition
composition isn't the same thing as inheritance, the concept of
composition is that some struct (class) copy all their properties
but it doesn't has a conecciont with their parent

this apply for methods to, a method that you declare for a especific
struct only works for that even if you make composition with that struct*/

// base
type Person struct {
	name string
	age  int
}

// base
type Employee struct {
	id int
}

// composition between Person and Employee
// we declare anonym structs
type FulltimeEmployee struct {
	Person
	Employee
}

// fucntion that only accepts Person struct
func SayHi(p Person) {
	fmt.Printf("%s has %d years old\nand say hi!\n", p.name, p.age)
}

func main() {
	ft := FulltimeEmployee{}
	ft.id = 1953
	ft.age = 26
	ft.name = "David Casas"

	fmt.Println(ft)
	// SayHi(ft) // <- error
}
