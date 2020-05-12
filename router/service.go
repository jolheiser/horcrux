package router

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	"go.jolheiser.com/horcrux/config"
	"go.jolheiser.com/horcrux/service"

	"github.com/go-chi/chi"
	"github.com/go-git/go-git/v5"
	gitConfig "github.com/go-git/go-git/v5/config"
	gitHttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"go.jolheiser.com/beaver"
)

func HandleService(w http.ResponseWriter, r *http.Request) {
	nameParam := strings.ToLower(chi.URLParam(r, "name"))
	repo, ok := repoMap[nameParam]
	if !ok {
		beaver.Errorf("repo config doesn't exist: %s", nameParam)
		return
	}
	serviceParam := strings.ToLower(chi.URLParam(r, "service"))
	payload, err := getPayload(r, serviceParam)
	if err != nil {
		beaver.Error(err)
		return
	}

	tmp, err := ioutil.TempDir(os.TempDir(), "horcrux")
	if err != nil {
		beaver.Errorf("could not create temp dir: %v", err)
		return
	}
	defer func() {
		if err := os.RemoveAll(tmp); err != nil {
			beaver.Errorf("could not remove temp dir: %v", err)
		}
	}()

	if err := gitSync(tmp, payload.URL(), repo); err != nil {
		beaver.Errorf("could not sync repository: %v", err)
		return
	}
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

func gitSync(tmp, repoURL string, cfg config.RepoConfig) error {
	repo, err := git.PlainClone(tmp, false, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		return fmt.Errorf("could not clone repository: %v", err)
	}

	var wg sync.WaitGroup
	for idx, svc := range cfg.ServiceConfigs() {
		if strings.EqualFold(svc.RepoURL, repoURL) {
			continue
		}
		svc := svc
		remoteName := fmt.Sprintf("horcrux-%d", idx)
		_, err := repo.CreateRemote(&gitConfig.RemoteConfig{
			Name: remoteName,
			URLs: []string{svc.RepoURL},
		})
		if err != nil {
			beaver.Errorf("could not create remote: %v", err)
			continue
		}

		auth := &gitHttp.BasicAuth{
			Username: svc.Username,
			Password: svc.AccessToken,
		}
		wg.Add(1)
		go func() {
			if err := repo.Push(&git.PushOptions{
				RemoteName: remoteName,
				Auth:       auth,
			}); err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
				beaver.Errorf("could not push to %s: %v", svc.RepoURL, err)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return nil
}
