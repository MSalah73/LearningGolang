// Simple example of creating and using a templatenin Go.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Type representing a gopher
type Gopher struct {
	Name string
}

// Handler for the hello-gopher route
func helloGopherHandler(w http.ResponseWriter, r *http.Request) {

	var gopherName string
	gopherName = r.URL.Query().Get("gophername")

	if gopherName == "" {
		gopherName = "Gopher"
	}

	gopher := Gopher{Name: gopherName}
	renderTemplate(w, "./NetHttp/Templates/Greeting.html", gopher)

}

// Template rendering function
func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {

	// ParseFiles can process multiple files
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing the template:", err)
	}
	t.Execute(w, templateData)
}

func main() {

	http.HandleFunc("/hello-gopher", helloGopherHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<p>Hello From Server</p>")
	})

	fmt.Println("Listening on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
