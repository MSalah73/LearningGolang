package main

import (
	f "fmt"
	"gofullstack/Packages/greetingspackage"
)

func main() {
	greetingspackage.PrintGreetings()
	greetingspackage.GopherGreetings()
	// greetingspackage.printGreetingsUnexported()

	f.Println("The value of the Magic Number is :", greetingspackage.MagicNumber)
}
