package main

import f "fmt"

var y = 4

// An example of a function that takes one input parameter, but does not return
// anything back. Thus, It is considered a void function.
func oddOrEven(x int) {

	// Functions have a local scpoe. The variable x is the only accessible within this
	// function. For that reason we can consider x to be a local variable of the function.
	if x%2 == 0 {
		f.Println("The number ", x, " is even.")
	} else {
		f.Println("The number", x, " is odd.")
	}

	// The variable y is visible inside this function, because it exists in the global
	// scope. Therefore, y is a global variable.
	if y%2 == 0 {
		f.Println("The number ", y, " is even.")
	} else {
		f.Println("The number ", y, " is odd.")
	}

}

// An variable of a function that takes 2 input parameters and returns an output
// of type int
func sumTheNumbers(x int, y int) int {
	return x + y
}

// An example of a function returnings multiple named outputs parameters.
// name returns increase readablity
func sumAndDiffOprationsOnTwoNumbers(x, y int) (sum int, difference int) {
	return y + x, x - y
}

// Returnings name value acn also be returned this way.
func sumAndDiffOprationsOnTwoNumbersWithEmptyReturn(x, y int) (sum int, difference int) {
	sum = y + x
	difference = x - y
	return
}

// An example of a variadic function, where we can supply a varying number of
// parameters
func variadicFunction(parameters ...int) int {
	total := 0
	for _, value := range parameters {
		total += value
	}
	return total
}

// This can also be demonstrated with a recursion
func variadicFunctionRecursion(parameters ...int) int {
	if len(parameters) == 0 {
		return 0
	}
	return parameters[0] + variadicFunction(parameters[1:]...)
}

// main() and init() functions are example of niladic function because it doesn't accept any argeuments.
func main() {
	oddOrEven(10)
	oddOrEven(7)
	f.Println("Value from sumTheNumbers is ", sumTheNumbers(10, y))
	f.Println("Value from sumAndDiffOprationOnNumbers are:")
	f.Println(sumAndDiffOprationsOnTwoNumbers(3, y))
	// https://stackoverflow.com/questions/52653779/how-to-parse-multiple-returns-in-golang
	// Check the above to print as the below
	//f.Printf("Value from sumAndDiffOprationsOnNumbers are:\nSum = %v\nDiff = %v\n", sumAndDiffOprationsOnTwoNumbers(3, y))
	f.Println("Value from variadicFunction is ", variadicFunction(1, 2, 3))
	f.Println("Value from variadicFunctionRecursion is ", variadicFunctionRecursion(1, 2, 3))

	// Example of an anonymous function
	func() {
		f.Println("Hi I'm an anonmous function. They call me that, because I don't have a name -- like lambda")
	}()
}
