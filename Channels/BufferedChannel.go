package main

import (
	"fmt"
)

func main() {
	//Note: If make a channel of capacity of size 1 mean you creating a regular channel

	// We create a buffered channel of strings with a capacity of 3 meaning channel
	// can hold upto 3 strings
	messageQueue := make(chan string, 3)
	messageQueue <- "one"
	messageQueue <- "two"
	messageQueue <- "three"
	//messageQueue <- "four?"

	// We drain the messageQueue by receiving all the values from the buffered
	// channel.
	// Note: this will not wait for all the strings to pop out the buffered channel
	// Try commenting one the line below
	fmt.Println(<-messageQueue)
	fmt.Println(<-messageQueue)
	fmt.Println(<-messageQueue)
}
