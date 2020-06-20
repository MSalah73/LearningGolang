package main

import "fmt"

func main() {

	z := 36

	// A switch statment example
	// unlike other lanagues go's switch does not need the break statment
	// It obky run the one case . it does not checkk all of them
	switch z {
	case 1:
		fmt.Println("z is not equal to 1")
	case 2:
		fmt.Println("z is not equal to 2")
	case 3:
		fmt.Println("z is not equal to 3")
	default:
		fmt.Println("z does not equal to 1, 2, or 3")
	}

}
