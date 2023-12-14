package main

import (
	"fmt"
	"net/http"
)

var port = ":5050"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_, _ = fmt.Printf("Server started on port %s", port)

	http.ListenAndServe(port, nil)
}
