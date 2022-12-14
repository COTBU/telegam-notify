package config

import (
	"bytes"
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Token   string `yaml:"token"`
	Channel int64  `yaml:"channel"`
	Port    string `yaml:"port"`
}

func (c Config) IsValid() bool {
	return c.Token != "" && c.Channel != 0 && c.Port != ""
}

func Get(path string) (*Config, error) {
	cfg := &Config{}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(file)

	decoder := yaml.NewDecoder(buf)
	if err := decoder.Decode(cfg); err != nil {
		return nil, err
	}

	if !cfg.IsValid() {
		return nil, errors.New("config is invalid")
	}

	return cfg, nil
}
