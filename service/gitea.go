package service

import (
	"go.jolheiser.com/beaver"
	"io/ioutil"
	"net/http"
)

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

func (g GiteaPayload) Validate(r *http.Request, secret string) bool {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		beaver.Errorf("Cannot read the request body: %s", err)
		return false
	}

	return compareHMAC(secret, string(b), r.Header.Get("X-Gitea-Signature"))
}
