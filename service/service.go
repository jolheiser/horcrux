package service

import "net/http"

type HorcruxPayload interface {
	URL() string
	GitRef() string
	Validate(r *http.Request, secret string) bool
}
