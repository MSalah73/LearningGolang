package main

import (
	"fmt"
)

var done chan bool = make(chan bool)

func printGreetings(source string) {

	for i := 0; i < 9; i++ {
		fmt.Println("Hello Gopher ", i, source)
	}

	if source == "goroutine" {
		done <- true
	}
}

func main() {

	go printGreetings("goroutine")
	printGreetings("main function")
	// Indiviual channels are sysnchronnous meaning that both sending and receiving
	// will wait until the other side is ready -- However, buffered channels are asynchronous
	// meaning the the both the sending and receiving messages through will not block
	// unless the channel buffer is full

	// this will wait until the bool value is returned
	<-done
	close(done)
	//Compare with goroutine.go in the goroutine folder
}
