package service

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"go.jolheiser.com/beaver"
	"io/ioutil"
	"net/http"
	"strings"
)

type GitHubPayload struct {
	Ref string `json:"ref"`
}

func (g GitHubPayload) GitRef() string {
	return g.Ref
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

	hash := hmac.New(sha1.New, []byte(secret))
	if _, err := hash.Write(b); err != nil {
		beaver.Errorf("Cannot compute the HMAC for request: %s", err)
		return false
	}

	expectedHash := hex.EncodeToString(hash.Sum(nil))
	beaver.Debugf("Expected Hash: %s", expectedHash)
	return gotHash[1] == expectedHash
}
