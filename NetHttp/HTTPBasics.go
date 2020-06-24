package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	response, err := http.Get("http://www.golang.org")
	if err != nil {
		log.Fatal("Encountered the following error: ", err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	// Convert byte array to a string
	fmt.Println(string(body))
}
