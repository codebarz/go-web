package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var port = ":5050"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	renderedTemplate, _ := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	err := renderedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("An error occured")
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_, _ = fmt.Printf("Server started on port %s", port)

	http.ListenAndServe(port, nil)
}
