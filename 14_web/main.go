package main

import (
	"fmt"
	"net/http"
)

// index - Function to handle "/" route
func index(w http.ResponseWriter, r *http.Request) {
	// Writes Hello World to Response Writer
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

// index - Function to handle "/" route
func about(w http.ResponseWriter, r *http.Request) {
	// Writes Hello World to Response Writer
	fmt.Fprintf(w, "<h1>About Page</h1>")
}

func main() {
	// Think of this like a router
	// Takes in a route and handles it via function
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	// Think of this like app.serve()
	// Starts up local web server
	http.ListenAndServe(":3000", nil)
}
