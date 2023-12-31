package web

import (
	"fmt"
	"myapp/pkg/handlers"
	"net/http"
)

var port = ":5050"

func Web() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_, _ = fmt.Printf("Server started on port %s", port)

	http.ListenAndServe(port, nil)
}
