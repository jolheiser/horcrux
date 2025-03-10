package main

import (
	"encoding/json"
	"time"
)

type Config struct {
	Interval Duration
	Storage  string
	Repos    []RepoConfig
}

type RepoConfig struct {
	Source string
	Dest   []DestConfig
}

type DestConfig struct {
	Forge DestForgeConfig
	URL   string
}

type DestForgeConfig struct {
	ForgeConfig
	Name      string
	TokenFile string
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
