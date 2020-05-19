package service

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"go.jolheiser.com/beaver"
	"net/http"
)

type HorcruxPayload interface {
	URL() string
	GitRef() string
	GitHead() string
	Validate(r *http.Request, secret string) bool
}

func compareHMAC(secret, payload, expectedHash string) bool {
	hash := hmac.New(sha1.New, []byte(secret))
	if _, err := hash.Write([]byte(payload)); err != nil {
		beaver.Errorf("Cannot compute the HMAC for request: %s", err)
		return false
	}

	hashSum := hex.EncodeToString(hash.Sum(nil))
	beaver.Debugf("Expected Hash: %s", expectedHash)
	beaver.Debugf("  Summed Hash: %s", hashSum)
	return expectedHash == hashSum
}
