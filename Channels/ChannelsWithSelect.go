package main

import "fmt"

// Simple implementation of fibonacci to show case non blocking and blocking
// channels via select
func fibonacci(c, quit chan int) {
	x, y, z := 0, 1, 0
	for {
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Print("Waited ", z, " loops to get ")
			z = 0
		case <-quit:
			fmt.Println("quit")
			return
			// Non Blocking: the default case will always run if non of the cases condition is met
			// Blocking: By remove the default case, the select will wait until one or more
			// cases conditions are met
			// Note: If multiple cases are ready, the select will choose one of the
			// cases at random
		default:
			z++
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
