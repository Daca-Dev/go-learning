package main

import "fmt"

/*
the way that we can ask for dinamic parametersm this meaning that we can pass one, two, ten, any quantity
of params to process in the function

the sintax is how we can se here
*/
func sumDinamic(values ...int) (total int) {
	for _, value := range values {
		total += value
	}
	return
}

/*
Also you can return many values and declare their name in the function
Go interpret the name as values that you going to use in the function and crete them
and also you can avoid define wich value return because Go do that for you*/
func maniReturns(value int) (double int, triple int, quad int) {
	double = value * 2
	triple = value * 3
	quad = value * 4
	return
}

func main() {
	fmt.Println(sumDinamic(1, 2))
	fmt.Println(sumDinamic(1, 2, 3))
	fmt.Println(sumDinamic(1, 2, 4, 5))
	fmt.Println(sumDinamic(1, 2, -5, 23, 45, 0))

	fmt.Println(maniReturns(4))
}
