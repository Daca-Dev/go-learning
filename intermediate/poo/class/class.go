package main

import "fmt"

/*
In Go the classes arent declarated with the keywork class, instead
we use the keyword struct that definen the structure or atributes of
the class that we want to create
*/
type Employee struct {
	id   int
	name string
}

/* The methods for our "classes" are created with functions, this type of
functions are called receiver function and the difference with a normal function
is that before the name of the function you define the Struct (class) which
you asign this method
take into account that you refer this Struct as value or pointer, is important
depending of the task that you do in the method or the object size*/
// receiver function
func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e Employee) GetId() int {
	return e.id
}

func (e Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}
	fmt.Println(e)
	e.id = 1
	e.name = "David"
	e.SetId(1953)
	e.SetName("David Alberto")
	fmt.Println(e)
}
