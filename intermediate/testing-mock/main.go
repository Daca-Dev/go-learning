package main

import (
	"time"
)

type Person struct {
	DNI  string
	name string
	age  int
}

type Employee struct {
	id      int
	postion string
}

type FullTimeEmployee struct {
	Person
	Employee
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	// code to consult a data base
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	// code to consult another data base
	return Employee{}, nil
}

func GetFullTimeEmployeeId(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee = FullTimeEmployee{}

	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Employee = e
	ftEmployee.Person = p

	return ftEmployee, nil
}
