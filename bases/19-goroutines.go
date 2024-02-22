package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

// GO DOC URL -> https://pkg.go.dev/?utm_source=godoc
// CONCURRENCY
// esta lidiando con multiples cosas mientras el paralelismo esta lidiando
// con multiples cosas al mismo tiempo

// the go routines normally use anonym functions

func main() {

	// packages that allow us to use go routines
	var wg sync.WaitGroup

	fmt.Println("Hello")
	wg.Add(1)

	go say("World", &wg)

	wg.Wait()

	go func(text string){
		fmt.Printf("I'm a anonimus function, Bye!\n%s", text)
	}("I'm David Casas")

	// say("hello 1")
	// go say("hello 2") // this going to be executy in concurrency
	// say("hello 3")
	// when main reach this point, the process kill even if concurrency is running

	time.Sleep(time.Second * 1) // add a time to wait (not recommend) for wait until gorouting finish
}
