package config

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"go.jolheiser.com/beaver"
	"gopkg.in/yaml.v2"
)

var defaultConfig string

type Config struct {
	Port         string       `yaml:"port"`
	LogLevel     string       `yaml:"log_level"`
	Repositories []RepoConfig `yaml:"repositories"`
}

type RepoConfig struct {
	Name   string           `yaml:"name"`
	Gitea  []*ServiceConfig `yaml:"gitea"`
	GitHub []*ServiceConfig `yaml:"github"`
	GitLab []*ServiceConfig `yaml:"gitlab"`
}

func (rc RepoConfig) ServiceConfigs() []*ServiceConfig {
	return append(rc.Gitea, append(rc.GitHub, rc.GitLab...)...)
}

type ServiceConfig struct {
	RepoURL     string `yaml:"repo_url"`
	Secret      string `yaml:"secret"`
	Username    string `yaml:"username"`
	AccessToken string `yaml:"access_token"`
}

func (sc ServiceConfig) HumanURL() string {
	return sc.RepoURL[:len(sc.RepoURL)-4]
}

func Load() (*Config, error) {
	bin, err := os.Executable()
	if err != nil {
		return nil, err
	}

	binDir := filepath.Dir(bin)
	if os.Getenv("HORCRUX_PATH") != "" {
		binDir = os.Getenv("HORCRUX_PATH")
	}

	configPath := path.Join(binDir, "horcrux.yml")
	if os.Getenv("HORCRUX_CONFIG") != "" {
		configPath = os.Getenv("HORCRUX_CONFIG")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.MkdirAll(path.Dir(configPath), os.ModePerm); err != nil {
			return nil, err
		}
		fi, err := os.Create(configPath)
		if err != nil {
			return nil, err
		}
		defer fi.Close()
		if _, err := fi.WriteString(defaultConfig); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg *Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	beaver.Console.Level = beaver.ParseLevel(cfg.LogLevel)
	return cfg, nil
}
