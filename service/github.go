package service

import (
	"io/ioutil"
	"net/http"
	"strings"

	"go.jolheiser.com/beaver"
)

type GitHubPayload struct {
	Ref  string `json:"ref"`
	Head string `json:"after"`
	Repo struct {
		CloneURL string `json:"clone_url"`
	} `json:"repository"`
}

func (g GitHubPayload) URL() string {
	return g.Repo.CloneURL
}

func (g GitHubPayload) GitRef() string {
	return g.Ref
}

func (g GitHubPayload) GitHead() string {
	return g.Head
}

func (g GitHubPayload) Validate(r *http.Request, secret string) bool {
	gotHash := strings.SplitN(r.Header.Get("X-Hub-Signature"), "=", 2)
	if gotHash[0] != "sha1" {
		return false
	}

	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		beaver.Errorf("Cannot read the request body: %s", err)
		return false
	}

	return compareHMAC(secret, string(b), gotHash[1])
}
