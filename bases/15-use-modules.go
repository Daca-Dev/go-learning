package main

import (
	"fmt"
	pk "platzi-basic-go mypackage"
)

// GO DOC URL -> https://pkg.go.dev/?utm_source=godoc

// type struct can be used as a class

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Nissan"
	fmt.Println(myCar)
}
