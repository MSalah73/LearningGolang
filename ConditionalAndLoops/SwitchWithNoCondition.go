package main

import (
	"fmt"
	"strconv"
)

// FizzBuzz is a game where numbers divisible by 3 would say Fizz, numbers
// divisible by 5 say Buzz and number divisible by 3 and 5 should say FizzBuzz.
// Otherwise, the program should print out the number
func main() {
	fmt.Print()
	for i := 0; i <= 100; i++ {
		toDisplay := ""
		// Switch without a condition is the same as writing switch true.
		// This lets you write a simpler and clearer long if-then-else statments
		switch {
		case i%15 == 0:
			toDisplay = "FizzBuzz"
		case i%3 == 0:
			toDisplay = "Fizz"
		case i%5 == 0:
			toDisplay = "Buzz"
		default:
			toDisplay = strconv.Itoa(i)

		}
		fmt.Println(toDisplay)
	}
}
