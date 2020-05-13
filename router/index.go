package router

import (
	"html/template"
	"net/http"

	"go.jolheiser.com/beaver"
)

var indexTmpl *template.Template

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if err := indexTmpl.Execute(w, repoMap); err != nil {
		beaver.Errorf("could not execute template: %v", err)
	}
}
