package config

import (
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()
	
	if config.Name != "test-release-oss" {
		t.Errorf("Expected name to be 'test-release-oss', got %s", config.Name)
	}
	
	if config.Version != "dev" {
		t.Errorf("Expected version to be 'dev', got %s", config.Version)
	}
	
	if config.Port != 8080 {
		t.Errorf("Expected port to be 8080, got %d", config.Port)
	}
	
	if config.Environment != "development" {
		t.Errorf("Expected environment to be 'development', got %s", config.Environment)
	}
}

func TestConfigToYAML(t *testing.T) {
	config := DefaultConfig()
	
	yaml, err := config.ToYAML()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	if len(yaml) == 0 {
		t.Error("Expected YAML output, got empty string")
	}
}
