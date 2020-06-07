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

	// Evan though we close this non-empty channel, we can still receive the
	// remaining values but not add them -- you can try by commenting the "three" above
	// and uncomment the "three" below
	close(messageQueue)
	//messageQueue <- "three"

	// We use the range keyword to interate over each element as it gets
	// received from the messageQueue.
	for m := range messageQueue {

		// notice we didn't use the <- operator this because range extract from the buffered
		// channel.
		fmt.Println(m)

	}
}
