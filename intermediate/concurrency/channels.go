package main

import "fmt"

/*
the channels is the elements that we use to share information between goroutines
a channel could have or not a buffer, that means that can host one or more values
is important never save more elements that the channel can host, because this
launch a panic and stop the program
*/
func main() {
	// create a channel
	// c := make(chan int) // with out buffer
	c := make(chan int, 3) // channel with buffer

	// send a value (1) to the channel
	// this channel doesn't has buffer so when you send a value ut block
	c <- 1
	c <- 2
	c <- 3
	c <- 4

	// read the value in the channel
	fmt.Println(<-c)
	fmt.Println(<-c)
}
