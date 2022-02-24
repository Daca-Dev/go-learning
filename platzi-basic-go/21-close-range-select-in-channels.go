package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}


func main() {
	c := make(chan string, 2)
	c <- "Message 1"
	c <- "Message 2"
	
	fmt.Println(len(c), cap(c)) // len: items qty, cap: channel capacity

	// c <- "Message 3" // runtime error because we exced the capacity


	close(c) // close the channel: the channel can be used even if have capacity

	// c <- "Message 3" // runtime error
	// range
	for message := range(c) {
		fmt.Println(message)
	}
	// select
	email1 := make(chan string) // if we dont specify limit, is dymanic
	email2 := make(chan string) // if we dont specify limit, is dymanic

	go message("Email-1", email1)
	go message("Email-2", email2)

	for i := 0; i < 2; i++ {
		select {
		case ml := <-email1:
			fmt.Println("Email received from email-1:", ml)
		case ml := <-email2:
			fmt.Println("Email received from email-2:", ml)
		}
	}

}