package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World!")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf(fmt.Sprintf("Number of bytes is: %d", n))
	})

	http.ListenAndServe(":5050", nil)
}
