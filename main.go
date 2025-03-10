package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/google/go-jsonnet"
)

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

	vm := jsonnet.MakeVM()
	cfgJSON, err := vm.EvaluateAnonymousSnippet(*configFlag, string(cfg))
	if err != nil {
		return fmt.Errorf("could not evaluate jsonnet: %w", err)
	}

	var config Config
	if err := json.Unmarshal([]byte(cfgJSON), &config); err != nil {
		return fmt.Errorf("could not unmarshal JSON from config: %w", err)
	}

	ticker := time.NewTicker(time.Duration(config.Interval))
	go func() {
		for {
			slog.Debug("running sync...")
			for _, r := range config.Repos {
				if err := r.Sync(); err != nil {
					slog.Error("could not sync repo", slog.String("repo", r.Source), slog.Any("err", err))
				}
			}
			<-ticker.C
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch

	return nil
}

func main() {
	if err := maine(); err != nil {
		slog.Error("error running horcrux", slog.Any("err", err))
	}
}
