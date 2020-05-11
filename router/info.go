package router

import (
	"github.com/go-chi/chi"
	"go.jolheiser.com/beaver"
	"html/template"
	"net/http"
	"strings"
)

var infoTmpl *template.Template

func HandleInfo(w http.ResponseWriter, r *http.Request) {
	nameParam := strings.ToLower(chi.URLParam(r, "name"))
	repo := repoMap[nameParam]

	if err := infoTmpl.Execute(w, repo); err != nil {
		beaver.Errorf("could not execute template: %v", err)
	}
}
