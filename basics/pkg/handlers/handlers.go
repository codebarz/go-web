package handlers

import (
	"myapp/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about")
}
