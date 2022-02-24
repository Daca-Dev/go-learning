package main

import (
	"fmt"
)

// GO DOC URL -> https://pkg.go.dev/?utm_source=godoc

// con los punteros podemos acceder directamente al valor en memoria
// de una variable ya que lo que almacena un puntero son valores de memoria

// son ampliamente usados para poder acceder y modificar los valores de variables
// en funciones (recuerda que las variables pasan por valor en una función)

// METODOS DE ESTRUCTURAS
// no son métodos como tal pero son funciones asociadas a las structuras que creamos
// lo emula metodos de POO (pero no es POO)

type pc struct {
	ram int
	brand string
	disk int
}

func (mp pc) String() string {
	return fmt.Sprintf("I have %d GB RAM, %d Gb ROM and I'm %s PC", mp.ram, mp.disk, mp.brand)
}

func main() {

	myPc := pc{ram: 32, brand: "aorus", disk: 500}
	fmt.Println(myPc)


}
