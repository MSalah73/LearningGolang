package main

import (
	"fmt"
)

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		fmt.Println("sending ", i+1)
		c <- i + 1
	}
}

func main() {

	c := make(chan int, 2)

	go squares(c)

	for i := 0; i <= 3; i++ {
		fmt.Println("receving ", i+1)
		num := <-c
		fmt.Println(num * num)
	}
}
