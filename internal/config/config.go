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

func LoadConfig() (*Config, error) {
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

func SaveConfig(config *Config) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("YAML marshal error: %w", err)
	}

	err = os.WriteFile("config.yaml", data, 0644)
	if err != nil {
		return fmt.Errorf("file writing error: %w", err)
	}

	return nil
}
