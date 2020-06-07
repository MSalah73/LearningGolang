package main

import (
	"fmt"
	"sync"
)

var greetings string
var howdyDone chan bool
var mutex = &sync.Mutex{}

func howdyGreetings() {

	mutex.Lock()
	greetings = "Howdy Gopher!"
	mutex.Unlock()
	howdyDone <- true

}

func main() {

	howdyDone = make(chan bool, 1)
	go howdyGreetings()
	mutex.Lock()
	greetings = "Hello Gopher!!!"
	// Note: we unlock after the print statment becouse it still considered
	// accessing the shared data which cause race condition.
	fmt.Println(greetings)
	mutex.Unlock()
	<-howdyDone

}
