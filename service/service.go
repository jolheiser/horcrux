package service

import "net/http"

type HorcruxPayload interface {
	URL() string
	GitRef() string
	GitHead() string
	Validate(r *http.Request, secret string) bool
}
