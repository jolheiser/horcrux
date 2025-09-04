package main

import (
	"bytes"
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"time"

	"go.jolheiser.com/cuesonnet"
)

//go:embed schema.cue
var schema cuesonnet.Schema

func maine() error {
	fs := flag.NewFlagSet("horcrux", flag.ExitOnError)
	jsonFlag := fs.Bool("json", false, "Print logs in JSON format")
	debugFlag := fs.Bool("debug", false, "Debug logging")
	configFlag := fs.String("config", ".horcrux.jsonnet", "Path to config file")
	if err := fs.Parse(os.Args[1:]); err != nil {
		return err
	}
	level := slog.LevelInfo
	if *debugFlag {
		level = slog.LevelDebug
	}
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}
	var handler slog.Handler = slog.NewTextHandler(os.Stderr, opts)
	if *jsonFlag {
		handler = slog.NewJSONHandler(os.Stderr, opts)
	}
	slog.SetDefault(slog.New(handler))

	cfg, err := os.ReadFile(*configFlag)
	if err != nil {
		return fmt.Errorf("could not read config file: %w", err)
	}

	var config Config
	if err := schema.Decode(bytes.NewReader(cfg), &config); err != nil {
		return fmt.Errorf("could not decode config file: %w", err)
	}

	if err := os.MkdirAll(config.Storage, os.ModePerm); err != nil {
		return fmt.Errorf("could not create storage repo at %q: %w", config.Storage, err)
	}

	git := sshGit(config.Key)
	ticker := time.NewTicker(time.Duration(config.Interval))
	go func() {
		for {
			slog.Debug("running sync...")
			for _, r := range config.Repos {
				go func(r RepoConfig) {
					// Check if we need to clone first
					repoPath := filepath.Join(config.Storage, r.Name)
					_, err := os.Stat(repoPath)
					if err != nil {
						if errors.Is(err, os.ErrNotExist) {
							if err := git(config.Storage, "clone", "--mirror", r.Source, r.Name); err != nil {
								slog.Error("could not clone repo", slog.String("repo", r.Source), slog.Any("err", err))
							}
						} else {
							slog.Error("could not stat repo path", slog.Any("err", err))
						}
					}

					// Update from remote
					if err := git(repoPath, "remote", "update", "--prune"); err != nil {
						slog.Error("could not update repo", slog.String("repo", r.Source), slog.Any("err", err))
					}

					// Push
					for _, dest := range r.Dest {
						slog.Debug("syncing repo", slog.String("repo", r.Source), slog.String("dest", dest))
						if err := git(repoPath, "push", "--mirror", "--force", dest); err != nil {
							slog.Error("could not push repo", slog.String("repo", r.Source), slog.String("dest", dest), slog.Any("err", err))
						}
					}
				}(r)
			}
			<-ticker.C
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch

	return nil
}

func sshGit(key string) func(string, ...string) error {
	return func(dir string, args ...string) error {
		cmd := exec.Command("git", args...)
		cmd.Dir = dir
		cmd.Env = append(os.Environ(), fmt.Sprintf(`GIT_SSH_COMMAND=ssh -i %s`, key))
		// cmd.Stdout = os.Stdout
		// cmd.Stderr = os.Stderr
		return cmd.Run()
	}
}

func main() {
	if err := maine(); err != nil {
		slog.Error("error running horcrux", slog.Any("err", err))
	}
}
