package main

import "fmt"

func main() {
	// This program will run fine since its syntactically correct and the
	// compiler will not catch any errors.
	// For error like these we use the vet command as shown below:
	// run: go vet WrongType.go
	fmt.Printf("String or number?: %s\n", 1)
}
