package main

import (
	"fmt"
	"net/http"
)

var port = ":5050"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_, _ = fmt.Printf("Server started on port %s", port)

	http.ListenAndServe(port, nil)
}
