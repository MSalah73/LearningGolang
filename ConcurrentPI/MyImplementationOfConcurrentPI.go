package main

import (
	"fmt"
)

// NilaKantha Series for calculating PI
// 3 + (4 / (2 * 3 * 4)) - (4 / (4 * 5 * 6)) + (4 / (6 * 7 * 8))....
// Let's use goroutines and channels to calculate PI

// This the NilaKantha algorithm trasualted into code
func Nilakantha(value chan float64, n int) {

	if n%2 != 0 {
		n *= 2
		fmt.Println("sending data NO ", n)
		value <- 4 / float64(n*(n+1)*(n+2))
	} else {
		n *= 2
		fmt.Println("sending data NO ", n)
		value <- -4 / float64(n*(n+1)*(n+2))
	}
}

func main() {

	PI := 3.0
	// running with larger channel size will elimeniate the error
	// off exceeding the goroutine limit
	value := make(chan float64, 2000)
	done := make(chan bool)

	// Using asynchronous function to run NilaKantha
	go func() {
		for i := 1; i <= 10000; i++ {
			go Nilakantha(value, i)
		}
	}()

	// This counter ensures the all the values from NilaKanta algorithm is
	// read
	counter := 0

	// Used a non blocking select
	for {
		select {
		case i := <-value:
			counter += 2
			fmt.Println("receiving data", counter)
			PI += i

			if counter == 20000 {
				go func() {
					done <- true
				}()
			}
		case <-done:
			fmt.Println(PI)
			return
		}

	}

}
