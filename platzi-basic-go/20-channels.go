package main

import "fmt"

// channels

// let you share data between go routines

func say(text string, c chan<- string) { // the chan,- indicates that only introduce data
	c <- text // the symbol <- at right indicates that we going to pass a date throught the channel
}

func main() {
	c := make(chan string, 1) // chan = channel and string is the type of data thats going to pass througt the channel, 1: limit
	fmt.Println("Hello")
	go say("World", c)

	fmt.Print(<-c) // the symbol at left indicates out the data
}