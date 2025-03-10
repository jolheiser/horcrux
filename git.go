package main

import (
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func git(dir string, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (r RepoConfig) Sync() error {
	tmp, err := os.MkdirTemp(os.TempDir(), "horcrux-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)
	if err := git(os.TempDir(), "clone", "--mirror", r.Source, tmp); err != nil {
		return err
	}
	for _, dest := range r.Dest {
		token := dest.Forge.Token
		if dest.Forge.TokenFile != "" {
			tokenBytes, err := os.ReadFile(dest.Forge.TokenFile)
			if err != nil {
				slog.Error("could not read token file", slog.Any("err", err))
				continue
			}
			token = strings.TrimSpace(string(tokenBytes))
		}
		client, err := NewForgeClient(dest.Forge.Name, ForgeConfig{
			Username:    dest.Forge.Username,
			Token:       token,
			RepoName:    dest.Forge.RepoName,
			Description: fmt.Sprintf("MIRROR of %s", r.Source),
			APIURL:      dest.Forge.APIURL,
		})
		if err != nil {
			return err
		}
		ok, err := client.CheckRepoExists()
		if err != nil {
			return err
		}
		if !ok {
			if err := client.CreateRepo(); err != nil {
				return err
			}
		}
		u, err := url.Parse(dest.URL)
		if err != nil {
			return err
		}
		u.User = url.UserPassword(dest.Forge.Username, token)
		if err := git(tmp, "remote", "add", "--mirror=push", dest.Forge.Name, u.String()); err != nil {
			return err
		}
		if err := git(tmp, "push", "--mirror", "--force-with-lease", dest.Forge.Name); err != nil {
			return err
		}
	}
	return nil
}
