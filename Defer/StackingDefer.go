package main

import (
	"fmt"
)

// Defered fuctions calls can be stacked.
func main() {
	fmt.Println("For Loop Counting to 10")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	// The stacked defred print statement 10. Once the surrounding function retuns
	// the defered stack will pop from tail to head.
	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("\nDefred For Loop that start from 10 and decreases to 1")
}
