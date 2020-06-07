package greetingspackage

import f "fmt"

// We indiacte to Gp that we want to export a function by upper casting the
// function's first letter
func PrintGreetings() {
	f.Println("I'm printing a message from the PrintGReetings() function")
}

// This function is a unexported (since the first letter is lower case)
// Since it's unexported it will only be visible to functions that are within
// the greeting package
func printGreetingsUnexported() {
	f.Println("I'm printing a message from the printGreetingsUnexported() function")
}
