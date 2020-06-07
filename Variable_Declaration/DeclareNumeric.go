package main

import "fmt"

/*
######## Numric Types #########
Integers:
int -- the bit value depends on your computer processor
uint -- same as above
uint8 .. unit64
int8 .. int64
Floating Point Numbers:
float32 and float64
Complex Numbers:
complex64 -- with float32 real and imaginary parts
complex128 -- with float64 real and imaginary parts
*/

func main() {
	// initial value is zero
	var myInteger int32

	fmt.Println("Value of myIngeger: ", myInteger)

	var myFloatingPointNumber float64 = 36.333
	var myOtherFloatingPointNumber float64 = 306.99996666

	fmt.Println("The sum of floating points numbers: ", myFloatingPointNumber+myOtherFloatingPointNumber)

	x, y, z := 0, 1, 2

	fmt.Printf("x: %d\ty: %d\tz: %d\n", x, y, z)

	//Example of Complex Number
	myComplexNumber := 5 + 24i

	fmt.Println("Value of ComplexNumber: ", myComplexNumber)

	// Example of grouping constant declaretion/initializations
	// NOTE: it figure out the type
	const (
		speedOfLight     = 299792458
		pi               = 3.14
		myFavoriteNumber = 7
	)

	// Same as above but with regular variable
	var (
		a int = 0 // type here is optional
		b     = 1.8 + 3i
		c     = 2.7
	)

	fmt.Printf("a: %v\tb: %v\tc: %v", a, b, c)
}
