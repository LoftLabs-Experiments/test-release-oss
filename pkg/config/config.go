package config

import (
	"gopkg.in/yaml.v3"
)

// Config represents the application configuration
type Config struct {
	Name        string            `yaml:"name" json:"name"`
	Version     string            `yaml:"version" json:"version"`
	Port        int               `yaml:"port" json:"port"`
	Environment string            `yaml:"environment" json:"environment"`
	Features    map[string]string `yaml:"features" json:"features"`
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		Name:        "test-release-oss",
		Version:     "dev",
		Port:        8080,
		Environment: "development",
		Features: map[string]string{
			"test-feature": "enabled",
		},
	}
}

// ToYAML converts the config to YAML
func (c *Config) ToYAML() ([]byte, error) {
	return yaml.Marshal(c)
}
