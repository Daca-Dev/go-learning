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
	// ch := make(chan int) // with out buffer
	ch := make(chan int, 3) // channel with buffer

	// send a value (1) to the channel
	// this channel doesn't has buffer so when you send a value ut block
	ch <- 1
	ch <- 2
	ch <- 3
	// c <- 4 // error if we send this 4th value, we going to exced the buffer, Deadlock error raised

	// read the value in the channel
	for v := range ch {
		fmt.Printf("Value: %d\n", v)
	}
}
