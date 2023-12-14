package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	renderedTemplate, _ := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	err := renderedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("An error occured")
		return
	}
}
