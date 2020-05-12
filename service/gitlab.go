package service

import "net/http"

type GitLabPayload struct {
	Ref     string `json:"ref"`
	Project struct {
		HTTPURL string `json:"http_url"`
	} `json:"project"`
}

func (g GitLabPayload) URL() string {
	return g.Project.HTTPURL
}

func (g GitLabPayload) GitRef() string {
	return g.Ref
}

func (g GitLabPayload) Validate(r *http.Request, secret string) bool {
	return r.Header.Get("X-Gitlab-Token") == secret
}
