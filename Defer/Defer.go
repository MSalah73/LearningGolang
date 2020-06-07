package main

import (
	"fmt"
)

// defer statement will execute after the encapsulated function finish executing
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
func main() {
	defer fmt.Println("Defered")
	fmt.Println("Not defered")
}
