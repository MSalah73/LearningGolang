package main

import (
	"fmt"
)

func writeMessageToChannel(message chan string) {

	message <- "Hello Gopher!"

}

func main() {

	fmt.Println("Channel Demo")

	message := make(chan string)
	go writeMessageToChannel(message)

	fmt.Println("Greating from the message channel: ", <-message)

	// This tells the profrom to wait until the message variable is empty
	close(message)
}
