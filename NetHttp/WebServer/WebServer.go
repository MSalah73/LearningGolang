// Using the username check from UnitTest ValidationKit
package main

import (
	"fmt"
	ValidationKit "gofullstack/UnitTest"
	"log"
	"net/http"
)

func helloGopherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gopher")
}

func checkUsernameSyntaxHanlder(w http.ResponseWriter, r *http.Request) {
	var usernameSynyaxResult bool
	username := r.URL.Query().Get("username")

	if username == "" {
		http.Error(w, "username not provided!", http.StatusInternalServerError)
	} else {
		usernameSynyaxResult = ValidationKit.CheckUsernameSyntax(username)
		fmt.Fprintf(w, "Synyax Check Result for %v is %v", username, usernameSynyaxResult)
	}
}
func main() {
	http.HandleFunc("/hello-gopher", helloGopherHandler)
	http.HandleFunc("/check-username-syntax", checkUsernameSyntaxHanlder)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<p>Hello From Server</p>")
	})

	fmt.Println("Listening on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
