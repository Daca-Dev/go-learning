package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan int, 2)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		ch <- 1
		wg.Add(1)
		go doSOmething(i, &wg, ch)
	}

	wg.Wait()
}

func doSOmething(i int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	fmt.Printf("Id: %d - Started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("Id: %d - Finished\n", i)
	<-ch
}
