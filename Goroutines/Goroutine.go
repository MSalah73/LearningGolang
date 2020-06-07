package main

import (
	"fmt"
)

func printGreatings(source string) {

	for i := 0; i < 9; i++ {
		fmt.Println("Hello Gopher!", i, source)
	}
}

// You will notice that the goroutine labeled function will sometimes execute
// sometimes it does not at all - this is because after the main function terminates
// the program will exit even if there some goroutine labeled function did not
// finish all its work.
func main() {

	go printGreatings("goroutine")
	printGreatings("main function")

}
