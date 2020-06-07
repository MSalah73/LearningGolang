package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

// The back tick character tell go to use literal values
const UsernameRegex string = `^@?(\w){1,15}$`

func main() {
	var usernameInput string
	flag.StringVar(&usernameInput, "username", "Gopher", "The GpherFace Username To Check")
	flag.Parse()

	fmt.Println("GopherFace Username Validation Checker")
	fmt.Println("Checking Syntax for username, \"", usernameInput, "\", resulted in: ", CheckUSernameSyntax(usernameInput))
}

func CheckUSernameSyntax(username string) bool {
	validationResult := false
	r, err := regexp.Compile(UsernameRegex)
	if err != nil {
		log.Fatal(err)
	}

	validationResult = r.MatchString(username)
	return validationResult
}
