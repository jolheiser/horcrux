package service

import "net/http"

type GiteaPayload struct {
	Secret string `json:"secret"`
	Ref    string `json:"ref"`
	Head   string `json:"after"`
	Repo   struct {
		CloneURL string `json:"clone_url"`
	} `json:"repository"`
}

func (g GiteaPayload) URL() string {
	return g.Repo.CloneURL
}

func (g GiteaPayload) GitRef() string {
	return g.Ref
}

func (g GiteaPayload) GitHead() string {
	return g.Head
}

func (g GiteaPayload) Validate(_ *http.Request, secret string) bool {
	return g.Secret == secret
}
