package router

import (
	"github.com/go-chi/chi"
	"go.jolheiser.com/horcrux/config"
)

var repoMap = make(map[string]config.RepoConfig)

func New(cfg *config.Config) *chi.Mux {
	// Init the repoMap
	for _, repo := range cfg.Repositories {
		repoMap[repo.Name] = repo
	}

	m := chi.NewMux()
	m.Route("/{name}", func(r chi.Router) {
		r.Get("/", HandleInfo)
		r.Post("/{service}", HandleService)
	})

	return m
}
