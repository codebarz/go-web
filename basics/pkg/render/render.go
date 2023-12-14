package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	renderedTemplate, _ := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	err := renderedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("An error occured")
		return
	}
}
