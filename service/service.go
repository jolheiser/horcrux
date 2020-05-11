package service

import "net/http"

type HorcruxPayload interface {
	GitRef() string
	Validate(r *http.Request, secret string) bool
}
