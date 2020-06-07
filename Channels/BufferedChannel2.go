package main

import (
	"fmt"
)

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		fmt.Println("receiving ", i+1)
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {

	c := make(chan int, 2)

	go squares(c)

	fmt.Println("sending 1")
	c <- 1
	fmt.Println("sending 2")
	c <- 2
	fmt.Println("sending 3")
	c <- 3
	fmt.Println("sending 4")
	c <- 4

}
