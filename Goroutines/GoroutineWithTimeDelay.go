package main

import (
	"fmt"
	"time"
)

func printGreatings(source string) {

	for i := 0; i < 9; i++ {
		fmt.Println("Hello Gopher!", i, source)
	}
	time.Sleep(time.Millisecond * 1)
}

// You will notice that the goroutine labeled function will sometimes execute
// sometimes it does not at all - this is because after the main function terminates
// the program will exit even if there some goroutine labeled function did not
// finish all its work.
// Note: Time will show more of the goroutine labeled functions because the main
// funtion will have to wait 20 second and while it delayed the goroutine function
// can prefrom its tasks
// NotePlus: Time delay is not a sustainable solutions
func main() {

	go printGreatings("goroutine")
	printGreatings("main function")

}
