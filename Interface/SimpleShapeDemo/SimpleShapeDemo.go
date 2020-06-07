package main

import (
	"fmt"
	"gofullstack/Interface/SimpleShapeInterface"
)

func main() {
	r := SimpleShapeInterface.NewRectangle(9, 9)
	t := SimpleShapeInterface.NewTriangle(9, 9)

	fmt.Println("The area of myRectangle: ", SimpleShapeInterface.ShapeArea(r))
	fmt.Println("The area of myTriangle: ", SimpleShapeInterface.ShapeArea(t))
}
