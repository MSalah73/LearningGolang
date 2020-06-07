// By adding the line below, we say that this file is a part of greetingspackage
package greetingspackage

import f "fmt"

var MagicNumber int

// This is an exported function and can be called outside the package
func GopherGreetings() {
	f.Println("A very jolly hello my fellow gophers! I'm printing from the GopherGreetings() fucntion")
	// Now we're calling an unexported package, because this function
	// is within the same package, we have access to call it.
	printGreetingsUnexported()
}

// This is a special function that initialize the state of of a package
// Go will automatically call the init function defined here prior to calling
// a command line main programs function
func init() {
	MagicNumber = 108
}
