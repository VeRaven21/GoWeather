package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DefaultCity string `yaml:"city"`
	Language    string `yaml:"lang"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("file reading error: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("YAML unmarshal error: %w", err)
	}
	return &config, nil
}
