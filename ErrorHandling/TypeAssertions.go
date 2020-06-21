package main

import (
	"fmt"
)

// You can do type check ith switch statement - Check out the conditionsAndLoops
// folder
func main() {
	var i interface{} = "Hello Gophers!"

	s := i.(string)
	fmt.Println("s contains", s)

	s, ok := i.(string)
	fmt.Println("s contains", s, "\b. Is it a string:", ok)

	f, ok := i.(float64)
	fmt.Println("f contains", f, "\b. Is it a string:", ok)

	f = i.(float64)
	fmt.Println("I hope one day I run and print", f)
}
