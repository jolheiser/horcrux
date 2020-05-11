package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"go.jolheiser.com/beaver"
	"go.jolheiser.com/horcrux/service"
	"io/ioutil"
	"net/http"
	"strings"
)

func HandleService(w http.ResponseWriter, r *http.Request) {
	nameParam := strings.ToLower(chi.URLParam(r, "name"))
	serviceParam := strings.ToLower(chi.URLParam(r, "service"))
	payload, err := getPayload(r, serviceParam)
	if err != nil {
		beaver.Error(err)
		return
	}
	_, _ = w.Write([]byte(fmt.Sprintf("%s\n%s\n%s", nameParam, serviceParam, payload.GitRef())))
}

func getPayload(r *http.Request, serviceParam string) (service.HorcruxPayload, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read the request body: %s", err)
	}
	if err := r.Body.Close(); err != nil {
		return nil, err
	}

	// Replace the Body in case any service needs it
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	var payload service.HorcruxPayload
	switch serviceParam {
	case "gitea":
		payload = &service.GiteaPayload{}
	case "github":
		payload = &service.GitHubPayload{}
	case "gitlab":
		payload = &service.GitLabPayload{}
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, fmt.Errorf("cannot unmarshal the request body: %s", err)
	}
	return payload, nil
}
