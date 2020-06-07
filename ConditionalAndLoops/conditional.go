package main

import f "fmt"

func main() {
	x := 9

	// An if statement
	if x == 9 {
		f.Println("x is equal to 9")
	}

	//An if statement followed by else statement example
	// This a short hand if
	if y := 18; y == 18 {
		f.Println("y is equal to 18")
	} else {
		f.Println("y is not equal to 18")
	}

	z := 36
	// an if - else if - else example
	if z == 1 {
		f.Println("z is equal to 1")
	} else if z == 2 {
		f.Println("z is equal to 2")
	} else if z == 3 {
		f.Println("z is equal to 3")
	} else {
		f.Println("z is not equal to 1, 2, or 3")
	}
}
