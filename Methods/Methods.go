package main

import (
	"fmt"
)

// This is very similar to stuct methods in the interface folder
// Methods withs with struct type and non-struct types
type MyInt int64

// The CallMe method has special receiver argument which is MyInt
// Methods are basicly a function with a receiver argument
func (i MyInt) CallMe() int64 {
	return int64(i)
}
func main() {
	i := MyInt(7)
	fmt.Println(i.CallMe())

}
