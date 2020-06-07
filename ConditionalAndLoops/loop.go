package main

import (
	f "fmt"
	t "time"
)

func main() {

	// Classic for loop
	//NOTE: sadly for some reason prefix add in not impelmented in go
	for i := 0; i < 10; i++ {
		if i == 0 {
			continue
		}
		f.Println("Inside classic for loop, value of i is:", i)
	}

	f.Println("\n\n")

	//Single condition for loop
	j := -20
	for j != 0 {
		f.Println("Inside sinlge condition loop, value iof j is: ", j)
		j++
	}

	f.Println("\n\n")

	loopTimer := t.NewTimer(t.Second * 9)
	// An infinite loop, don't worry  we'll break out of it in 9 seconds
	for {
		f.Println("Inside the infinite loop! -- will exit after nine seconds")

		<-loopTimer.C
		break
	}
}
