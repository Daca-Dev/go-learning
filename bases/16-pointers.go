package main

import (
	"fmt"
)

// GO DOC URL -> https://pkg.go.dev/?utm_source=godoc

// con los punteros podemos acceder directamente al valor en memoria
// de una variable ya que lo que almacena un puntero son valores de memoria

// son ampliamente usados para poder acceder y modificar los valores de variables
// en funciones (recuerda que las variables pasan por valor en una función)

func main() {

	a := 50
	b := &a

	fmt.Println(a) // valor
	fmt.Println(b) // puntero


	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

}
