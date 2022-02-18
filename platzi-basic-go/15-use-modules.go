package main

import (
	"fmt"
	pk "go-dev learning-go platzi-basic-go"
)

// GO DOC URL -> https://pkg.go.dev/?utm_source=godoc

// type struct can be used as a class

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Nissan"
	fmt.Println(myCar)
}
