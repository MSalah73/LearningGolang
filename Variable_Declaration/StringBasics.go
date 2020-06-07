package main

import "fmt"

/*
Strings
* Strings follow zero-based numbering and are represnted in UTF-8 format
* The zero value of string type is ""
*/
func main() {
	var subject string = "Gopher"

	fmt.Println("The first element of the 'Gopher' is: ", string("Gopher"[0]))
	fmt.Printf("The first vlaue of the subject string is: %v\n", string(subject[0]))
	fmt.Printf("The last vlaue of the subject string is: %v\n", string(subject[len(subject)-1]))
	fmt.Println("Hello " + subject + "!")
}
