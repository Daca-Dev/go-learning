package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup // variable that allow us to handler the gorutine

	for i := 0; i < 10; i++ {
		wg.Add(1) // add one to the counter of gorutines
		go doSomething(i, &wg)
	}

	wg.Wait() // this line block the run until all the gorutine are complete
}

func doSomething(i int, wg *sync.WaitGroup) {

	defer wg.Done() // reduce en 1 de counter

	fmt.Printf("Started - %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Finished %d\n", i)
}
