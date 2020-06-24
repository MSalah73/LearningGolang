// Routes basics
package main

import (
	"fmt"
	"log"
	"net/http"
)

// Routes
func talkToServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Thank for contacting server....Plesae Hold\n\n\n\n\n\nServer says hi BYE\n\nServer connection lost")
}

func main() {

	// You also display a html file this way
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./NetHttp/HTML/StaticPage.html")
	})

	// Assigning paths to your routes
	http.HandleFunc("/talkToServer", talkToServer)

	fmt.Println("Listening on port 8000")
	// You replace the nil from the ListenAndServe function with a
	// HandleFunc to display a home page here.
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
