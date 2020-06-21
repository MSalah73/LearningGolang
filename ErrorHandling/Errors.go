package main

import (
	"fmt"
	"time"
)

// Go programs express error state with error values.

// the error type is a built-in interface
// type error interface{
//   Error() string
// }

type MyError struct {
	When time.Time
	What string
}

func (err *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", err.When, err.What)
}

func run() error {
	return &MyError{time.Now(), "I did not work"}
}

func main() {
	// A nil error denotes success; a non-nil error denotes failure.
	if err := run(); err != nil {
		// we don't have to call Error()
		// The fmt package looks for the error interface when printing values.
		fmt.Println(err)
	}
}

// Functions often return an error value, and calling code should handle
// errors by testing whether the error equals nil.
// i, err := strconv.Atoi("42")
// if err != nil {
//     fmt.Printf("couldn't convert number: %v\n", err)
//     return
// }
// fmt.Println("Converted integer:", i)
