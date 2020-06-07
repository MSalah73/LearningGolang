package main

import "fmt"

// start from zero to what last value which in this case its 3
// the _ value mean we are discarding the value
const (
	_ = iota
	TraficLightStateRedLight
	TraficLightStateGreenLight
	TraficLightStateYellowLight
)

const (
	LosAngeles = 1984 + (iota * 4)
	Seoul
	Barcelona
	Atlanta
	Sydney
	Athens
	Beijing
	London
	Rio
	Tokyo
)

func main() {
	fmt.Println("Red Light State code: ", TraficLightStateRedLight)
	fmt.Println("Yellow Light State code: ", TraficLightStateYellowLight)
	fmt.Println("Green Light State code: ", TraficLightStateGreenLight)

	fmt.Println("\n######### A more realistic example #########\n")

	fmt.Println("These cities hosted or will host the summer Olympics in the provided year...")
	fmt.Printf("%-18s %-18s \n", "City", "Year")
	fmt.Printf("%-18s %-18v \n", "Los Angeles", LosAngeles)
	fmt.Printf("%-18s %-18v \n", "Atlanta", Atlanta)
	fmt.Printf("%-18s %-18v \n", "Tokyo", Tokyo)

}
