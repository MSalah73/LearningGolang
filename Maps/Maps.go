package main

import (
	"fmt"
	"sort"
)

func printNationalCaptialsMap(capitalsMap map[string]string) {

	for key, value := range capitalsMap {
		fmt.Println("The captial of", key, "is ", value)
	}
}

func printSortedNationalCaptialsMap(capitalsMap map[string]string) {

	keys := make([]string, len(capitalsMap))

	for key, _ := range capitalsMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, v := range keys {
		if v == "" {
			continue
		}
		fmt.Println("The captial of", v, "is ", capitalsMap[v])
	}

}

func nationCapitalsExample() {

	// Note: In Go, maps have no contract to maintain the order of the keys
	var nationCapitals map[string]string = make(map[string]string)
	nationCapitals["Afganistan"] = "Kabul"
	nationCapitals["Canada"] = "Ottawa"
	nationCapitals["Japan"] = "Tokyo"
	nationCapitals["Kenya"] = "Nairobi"
	nationCapitals["India"] = "New Dwlhi"
	nationCapitals["Mexico"] = "Mexico City"
	nationCapitals["South Korea"] = "Seoul"
	nationCapitals["United Kingdom"] = "London"
	nationCapitals["USA"] = "D.C."
	nationCapitals["Taiwan"] = "Taipei"

	fmt.Println("Print the map unsorted (random order): ")
	printNationalCaptialsMap(nationCapitals)
	fmt.Println("\n")

	fmt.Println("Print the map sorted by key (nation name): ")
	printSortedNationalCaptialsMap(nationCapitals)
	fmt.Println("\n")
}

func main() {
	nationCapitalsExample()
}
