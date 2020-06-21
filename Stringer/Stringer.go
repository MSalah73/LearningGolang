package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Age  int
}

func (animal Animal) String() string {
	return fmt.Sprintf("%v (%v years)", animal.Name, animal.Age)
}

func main() {
	Dog := Animal{"Biggs", 12}
	Cat := Animal{"Wedge", 10}
	// The fmt packageand many other look for this interface to print values
	// so you don't have write Dog.String()
	fmt.Println(Dog, "-", Cat)

}
