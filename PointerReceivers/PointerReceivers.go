package main

import "fmt"

// Note: The animal to human program is totally not accurate
type Animal struct {
	Age         float64
	Name        string
	Type        string
	AvgLifeSpan float64
}

func (animal Animal) AgeInHumanYears() {
	avgAge := 80.0 / animal.AvgLifeSpan
	fmt.Println(animal.Type, "\bs average lifespan in human year is", animal.AvgLifeSpan, "years.")
	fmt.Println(animal.Name, "\b's average lifespan in human year is", avgAge*animal.Age, "years.")

}

// Try removing the pointer and observe what happens to the output
// The pointer receiver copy the address of the object. If you remove
// the pointer it will be a value receiver which copies the whole object
// and pass in the AgeUp function.
func (animal *Animal) AgeUp() {
	animal.Age += 1
}

func main() {
	myAnimal := Animal{1, "Berserker", "Dinosaur", 300}
	myAnimal.AgeInHumanYears()
	myAnimal.AgeUp()
	myAnimal.AgeInHumanYears()

}
