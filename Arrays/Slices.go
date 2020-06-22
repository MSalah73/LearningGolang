package main

import "fmt"

// Slices get passed by refrence into Functions, meaning if you make changes
// to a slice within a function, our changes will be reflected to the slice
// that was passed into the function.
func populateIntegerSlice(toPopulate []int) {

	// We set values in a slice, just like we do with arrays using the square
	// brackets notation [] with the elements index enclosed within the square
	// brackets.
	toPopulate[0] = 3
	toPopulate[1] = 6
	toPopulate[2] = 9
	toPopulate[3] = 12
	toPopulate[4] = 15
}

func changeLineupExample() {

	// Declaring and initializing a slice representing the original members of
	// rock band, INXS using a slice literal (notice it looks just like an array literal
	// except without the element count)

	fmt.Println("The original INXS lineup")
	inxs := []string{"Micheal", "Andrew", "Jon", "Tim", "Kirk", "Garry"}
	fmt.Println(inxs, "\n")

	// Micheal left the band in 1997
	fmt.Println("The lineup without Micheal:")
	inxs = append(inxs[:0], inxs[1:]...)
	fmt.Println("This is the inxs[:0]", inxs[:0], "\n")
	fmt.Println(inxs, "\n")

	// JD joins the band in 2005
	fmt.Println("The lineup with a new member, JD")
	inxs = append(inxs, "JD")
	fmt.Println(inxs, "\n")
}
func main() {

	// We use the make() built-in method yto create a new slice of length 5
	// Here we make a slice of integers pf length 5.
	var mySlice []int = make([]int, 5)
	fmt.Printf("Content of mySlice: %v\n", mySlice)

	populateIntegerSlice(mySlice)
	fmt.Printf("Content of mySlice: %v\n", mySlice)
	// We can use the len() function to return the length of the slice
	fmt.Println("The length of mySlice is: ", len(mySlice))
	// We can use the cap() function to return the capacity of the slice
	fmt.Println("The capacity of mySlice is: ", cap(mySlice))

	// Add a new element to mySlice, and notice the change to the length and
	// capacity of the slice.
	// https://stackoverflow.com/questions/32995623/why-does-slice-capacity-with-odd-numbers-differ-from-behavior-with-even-numbers
	// the capacity number is rounding up the slice capacity to fill the allocated memory blocks.
	fmt.Println("After adding a new element to mySlice...\n")
	mySlice = append(mySlice, 18)
	fmt.Printf("Content of mySlice: %v\n", mySlice)
	fmt.Println("The length of mySlice is: ", len(mySlice))
	fmt.Println("The capacity of mySlice is: ", cap(mySlice))

	// This short hand notation allows us to get the elements from index 1 up
	// to (but not including ) index 4.
	s := mySlice[1:4]
	fmt.Println("mySlice[1:4]", s, "\n")

	// When you slice a slice, you get a refrence back to that slice. Any changes
	// you make to the subslice will be reflected in the original slice.
	// This will cause mySlice[1] t be equal to 7
	s[0] = 7
	fmt.Println("mySlice", mySlice)

	// All elements in mySlice up to (not including) the element at index 4.
	t := mySlice[:4]
	fmt.Println("mySlice[:4]", t, "\n")

	// The element from (and including) the element at index 1
	u := mySlice[1:]
	fmt.Println("mySlice[1:]", u, "\n")

	changeLineupExample()
}
