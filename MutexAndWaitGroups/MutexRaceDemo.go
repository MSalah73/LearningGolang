package main

import (
	"fmt"
)

var greetings string
var howdyDone chan bool

func howdyGreetings() {

	greetings = "Howdy Gopher!"
	howdyDone <- true

}

// Note: run this program with the race flag -race
func main() {

	howdyDone = make(chan bool, 1)
	go howdyGreetings()
	greetings = "Hello Gopher!!!"
	fmt.Println(greetings)
	<-howdyDone

}
