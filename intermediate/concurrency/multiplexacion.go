package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go DoSomething(d1, ch1, 1)
	go DoSomething(d2, ch2, 2)

	for i := 0; i < 2; i++ {
		select {
		case channelMsg1 := <- ch1:
			fmt.Println(channelMsg1)
		case channelMsg2 := <- ch2:
			fmt.Println(channelMsg2)
		}
	}
}

func DoSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}