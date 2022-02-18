package main

import "fmt"

// GO DOC URL -> https://pkg.go.dev/?utm_source=godoc

func main() {
	// module := 5 % 78
	// switch module {
	// case 0:
	// 	fmt.Println("Es par")
	// case 1:
	// 	fmt.Println("Es impar")
	// default:
	// 	fmt.Println("Default value")
	// }

	switch module := 5 % 78; module {
	case 0:
		fmt.Println("Es par")
	case 1:
		fmt.Println("Es impar")
	default:
		fmt.Println("Default value")
	}

	// switch without condition -> is used to write larger if else estatements
	value := 200
	switch {
		case value > 100:
			fmt.Println("Es mayor a 100")
		case value < 100:
			fmt.Println("Es menor a 100")
		default:
			fmt.Println("No condition")
	}
}
