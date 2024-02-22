package main

import (
	"fmt"
)

// GO DOC URL -> https://pkg.go.dev/?utm_source=godoc

// type struct can be used as a class
type car struct {
	brand string
	model string
	year int
	seatings int
}


func main() {
	// CLASES
	//! Go does not have Class but it can make something similar
	myCar := car{
		brand: "Nissan",
		model: "Murano",
		year: 2022,
		seatings: 7,
	}

	fmt.Println(myCar)

	// other way we can instance
	var otherCar car
	otherCar.brand = "Mercedez Benz"
	fmt.Println(otherCar)
}
