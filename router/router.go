package router

import (
	"html/template"

	"go.jolheiser.com/horcrux/config"

	"github.com/go-chi/chi"
)

var repoMap = make(map[string]config.RepoConfig)

func New(cfg *config.Config) *chi.Mux {
	// Init the repoMap
	for _, repo := range cfg.Repositories {
		repoMap[repo.Name] = repo
	}

	m := chi.NewMux()
	m.Get("/", HandleIndex)
	m.Route("/{name}", func(r chi.Router) {
		r.Get("/", HandleInfo)
		r.Post("/{service}", HandleService)
	})

	return m
}

func funcMap() template.FuncMap {
	return template.FuncMap{
		"Version": func() string {
			return config.Version
		},
	}
}
