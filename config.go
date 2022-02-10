package main

import (
	"encoding/json"
	"os"
)

const (
	configFileName = "hotpocket.json"
)

type Config struct {
	Directory  string   `json:"directory"`
	Command    string   `json:"command"`
	Arguments  []string `json:"arguments"`
	Exceptions []string `json:"ExceptionFiles"`
}

func LoadConfig() (*Config, error) {
	f, err := os.Open(configFileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cfg := &Config{}
	if err := json.NewDecoder(f).Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
