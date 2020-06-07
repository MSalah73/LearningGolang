// IMPORTANAT NOTE:
// All types in go impelments the empty interface{}
// A function that retrun a value of inteface can return any type
// We can stpre hetrogeneous values in homogeneous values like in an array, map, or slice using the the interface{} type

package main

import (
	"fmt"
	"gofullstack/Interface/SimpleShapeInterface"
	"math/rand"
	"time"
)

func giveMeARandomShape() interface{} {
	var shape interface{}
	var randomShapesSlice []interface{} = make([]interface{}, 2)

	// Seed the random generator
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	// Pick a random number, either 1 or 0
	randomNumber := r.Intn(2)
	fmt.Println("Random Number: ", randomNumber)

	// Let's make a new rectangle
	rectangle := SimpleShapeInterface.NewRectangle(3, 4)

	// Let's make a new Triangle
	triangle := SimpleShapeInterface.NewTriangle(4, 4)

	// Let's store the shapes into a slice data structure
	randomShapesSlice[0] = rectangle
	randomShapesSlice[1] = triangle
	shape = randomShapesSlice[randomNumber]

	return shape
}

func main() {
	myRectangle := SimpleShapeInterface.NewRectangle(3, 9)
	myTriangle := SimpleShapeInterface.NewTriangle(4, 9)
	shapesSlice := []interface{}{myRectangle, myTriangle}
	for index, shape := range shapesSlice {
		fmt.Printf("Shape in index, %d, of the shapesSlice is a %T\n", index, shape)
	}

	fmt.Println()

	aRandomShape := giveMeARandomShape()

	fmt.Printf("The type of aRandomShape is %T\n", aRandomShape)

	//Note: the .(type) method must be used with a switch command
	switch t := aRandomShape.(type) {
	case *SimpleShapeInterface.Rectangle:
		fmt.Println("I got a rectangle with an area equal to ", t.Area())
	case *SimpleShapeInterface.Triangle:
		fmt.Println("I got a trianlge with an area equal to ", t.Area())
	default:
		fmt.Println("I don't recognize what I got back!")
	}

}
