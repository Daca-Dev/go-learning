package main

import "fmt"

func Generator(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

// write channel chan<-
// read channel <-chan
func Double(chIn <-chan int, chOut chan<- int) {
	for value := range chIn {
		chOut <- 2 * value
	}
	close(chOut)
}

func Print(c <-chan int) {
	for value := range c {
		fmt.Printf("Value: %d\n", value)
	}
}

func main() {
	generator := make(chan int)
	doubles := make(chan int)

	go Generator(generator)
	go Double(generator, doubles)
	Print(doubles)
}
