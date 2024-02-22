package main

import "fmt"

/*
We can say that an interface is like a contract that define some methods and properties
that a structure should has, this interfaces are asigned implicity by Go, depending on
the methods and properties that de struct has

this means that is more easy reuse code and other functions could work for many clases
without problem
*/

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

// Full time employee struc and methods
type FullTimeEmployee struct {
	Person
	Employee
}

func (fte FullTimeEmployee) GetName() string {
	return fte.name + " Full time"
}

func (fte FullTimeEmployee) SetName(name string) {
	fte.name = name
}

// Temporal employee struct and methods
type TemporalEmployee struct {
	Person
	Employee
}

func (te TemporalEmployee) GetName() string {
	return te.name
}

func (te TemporalEmployee) SetName(name string) {
	te.name = name
}

// interface
type PersonActions interface {
	GetName() string
	SetName(string)
}

// function base in interfaces
func PrintName(p PersonActions) {
	name := p.GetName()
	fmt.Printf("My name is %s\n", name)
}

func main() {
	fte := FullTimeEmployee{}
	fte.id = 1953
	fte.name = "David Casas"
	fte.age = 26

	fte.SetName("Master chief")

	te := TemporalEmployee{}
	te.id = 1953
	te.name = "David Casas"
	te.age = 26

	te.SetName("Cortana")

	PrintName(fte)
	PrintName(te)

}
