package main

import (
	"encoding/json"
	"time"
)

type Config struct {
	Key      string
	Interval Duration
	Storage  string
	Repos    []RepoConfig
}

type RepoConfig struct {
	Name   string
	Source string
	Dest   []string
}

type Duration time.Duration

func (d *Duration) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	dur, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	*d = Duration(dur)
	return nil
}
