package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * 5,
}

// Configuration for forge clients
type ForgeConfig struct {
	Username    string
	Token       string
	RepoName    string
	Description string
	Private     bool
	APIURL      string
}

// Common interface for all forge clients
type ForgeClient interface {
	CheckRepoExists() (bool, error)
	CreateRepo() error
}

// GitHub implementation
type GitHubClient struct {
	config ForgeConfig
}

func NewGitHubClient(config ForgeConfig) *GitHubClient {
	if config.APIURL == "" {
		config.APIURL = "https://api.github.com"
	}
	return &GitHubClient{config: config}
}

func (c *GitHubClient) CheckRepoExists() (bool, error) {
	url := fmt.Sprintf("%s/repos/%s/%s", c.config.APIURL, c.config.Username, c.config.RepoName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", "token "+c.config.Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	}

	body, _ := io.ReadAll(resp.Body)
	return false, fmt.Errorf("GitHub API error: %d - %s", resp.StatusCode, string(body))
}

func (c *GitHubClient) CreateRepo() error {
	url := fmt.Sprintf("%s/user/repos", c.config.APIURL)

	payload := map[string]interface{}{
		"name":        c.config.RepoName,
		"description": c.config.Description,
		"private":     c.config.Private,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "token "+c.config.Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		return nil
	}

	body, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("GitHub API error: %d - %s", resp.StatusCode, string(body))
}

// GitLab implementation
type GitLabClient struct {
	config ForgeConfig
}

func NewGitLabClient(config ForgeConfig) *GitLabClient {
	if config.APIURL == "" {
		config.APIURL = "https://gitlab.com/api/v4"
	}
	return &GitLabClient{config: config}
}

func (c *GitLabClient) CheckRepoExists() (bool, error) {
	// URL encode the repo path for GitLab
	encodedPath := strings.Replace(c.config.Username+"/"+c.config.RepoName, "/", "%2F", -1)
	url := fmt.Sprintf("%s/projects/%s", c.config.APIURL, encodedPath)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("PRIVATE-TOKEN", c.config.Token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	}

	body, _ := io.ReadAll(resp.Body)
	return false, fmt.Errorf("GitLab API error: %d - %s", resp.StatusCode, string(body))
}

func (c *GitLabClient) CreateRepo() error {
	url := fmt.Sprintf("%s/projects", c.config.APIURL)

	visibility := "public"
	if c.config.Private {
		visibility = "private"
	}

	payload := map[string]interface{}{
		"name":        c.config.RepoName,
		"description": c.config.Description,
		"path":        c.config.RepoName,
		"visibility":  visibility,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("PRIVATE-TOKEN", c.config.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		return nil
	}

	body, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("GitLab API error: %d - %s", resp.StatusCode, string(body))
}

// Gitea/Codeberg/Forgejo implementation
type GiteaClient struct {
	config ForgeConfig
}

func NewGiteaClient(config ForgeConfig) *GiteaClient {
	// Default to Codeberg if no base URL provided
	if config.APIURL == "" {
		config.APIURL = "https://codeberg.org/api/v1"
	}
	return &GiteaClient{config: config}
}

func (c *GiteaClient) CheckRepoExists() (bool, error) {
	url := fmt.Sprintf("%s/repos/%s/%s", c.config.APIURL, c.config.Username, c.config.RepoName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", "token "+c.config.Token)
	req.Header.Set("Accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	}

	body, _ := io.ReadAll(resp.Body)
	return false, fmt.Errorf("Gitea API error: %d - %s", resp.StatusCode, string(body))
}

func (c *GiteaClient) CreateRepo() error {
	url := fmt.Sprintf("%s/user/repos", c.config.APIURL)

	payload := map[string]interface{}{
		"name":        c.config.RepoName,
		"description": c.config.Description,
		"private":     c.config.Private,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "token "+c.config.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		return nil
	}

	body, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("Gitea API error: %d - %s", resp.StatusCode, string(body))
}

// SourceHut implementation
type SourceHutClient struct {
	config ForgeConfig
}

func NewSourceHutClient(config ForgeConfig) *SourceHutClient {
	if config.APIURL == "" {
		config.APIURL = "https://git.sr.ht/api"
	}
	return &SourceHutClient{config: config}
}

func (c *SourceHutClient) CheckRepoExists() (bool, error) {
	url := fmt.Sprintf("%s/repos/%s/%s", c.config.APIURL, c.config.Username, c.config.RepoName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", "token "+c.config.Token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	}

	body, _ := io.ReadAll(resp.Body)
	return false, fmt.Errorf("SourceHut API error: %d - %s", resp.StatusCode, string(body))
}

func (c *SourceHutClient) CreateRepo() error {
	url := fmt.Sprintf("%s/repos", c.config.APIURL)

	visibility := "public"
	if c.config.Private {
		visibility = "private"
	}

	payload := map[string]interface{}{
		"name":        c.config.RepoName,
		"description": c.config.Description,
		"visibility":  visibility,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "token "+c.config.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		return nil
	}

	body, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("SourceHut API error: %d - %s", resp.StatusCode, string(body))
}

// Factory function to create the appropriate forge client
func NewForgeClient(forgeType string, config ForgeConfig) (ForgeClient, error) {
	switch strings.ToLower(forgeType) {
	case "github":
		return NewGitHubClient(config), nil
	case "gitlab":
		return NewGitLabClient(config), nil
	case "gitea", "codeberg", "forgejo":
		return NewGiteaClient(config), nil
	case "sourcehut", "srht":
		return NewSourceHutClient(config), nil
	default:
		return nil, fmt.Errorf("unsupported forge type: %s", forgeType)
	}
}
