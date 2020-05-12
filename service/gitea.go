package service

import "net/http"

type GiteaPayload struct {
	Secret string `json:"secret"`
	Ref    string `json:"ref"`
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

func (g GiteaPayload) Validate(_ *http.Request, secret string) bool {
	return g.Secret == secret
}
