package main

import (
	"flag"
	"fmt"
)

func main() {

	var gopherName string
	// The first value is address of the gophername variable
	// The second value is the flag -- usage -gophername String
	// The third value is the default value
	// The last string is a describtion of the flag -- this appears when the user
	// run the program with -h ot -help
	flag.StringVar(&gopherName, "gophername", "Gopher", "The name of the Gopher")
	flag.Parse()
	fmt.Println("Hello " + gopherName + "!")
}
