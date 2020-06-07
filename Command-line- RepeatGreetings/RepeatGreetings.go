package main

import (
	f "fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var repeatCount int
	var err error

	if len(os.Args) >= 2 {
		// convert cmd arg to integers with Atoi
		repeatCount, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < repeatCount; i++ {
			f.Println("Hello Gopher!")
			// use | wc -l to count printed lines
		}
	} else {
		f.Println("Usage:\nRepeatGreetings.exe <Number>")
	}
}
