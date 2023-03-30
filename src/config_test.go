package main

import (
	"os"
	"testing"
	// "time"

	"gopkg.in/yaml.v3"
)

const testConfigFileContent = `
environment: test
host: localhost
port: 8080
redis: redis://user:password@localhost:6379/0
cache:
  java_status_duration: 5m
  bedrock_status_duration: 10m
  icon_duration: 1h
`

func TestReadFile(t *testing.T) {
	// Create a temporary file to simulate a config file
	tempFile, err := os.CreateTemp("", "test_config_*.yaml")
	if err != nil {
		t.Fatalf("Failed to create a temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write the config content to the temporary file
	if _, err := tempFile.Write([]byte(testConfigFileContent)); err != nil {
		t.Fatalf("Failed to write content to the temporary file: %v", err)
	}
	tempFile.Close()

	// Load the config using ReadFile method
	var config Config
	if err := config.ReadFile(tempFile.Name()); err != nil {
		t.Fatalf("Failed to read config file: %v", err)
	}

	// Validate the config values
	if config.Environment != "test" || config.Host != "localhost" || config.Port != 8080 || *config.Redis != "redis://user:password@localhost:6379/0" {
		t.Errorf("Unexpected config values: %#v", config)
	}
}

func TestOverrideWithEnvVars(t *testing.T) {
	os.Setenv("ENVIRONMENT", "testenv")
	os.Setenv("HOST", "127.0.0.1")
	os.Setenv("PORT", "8081")
	os.Setenv("REDIS_URL", "redis://user:password@localhost:6379/1")
	defer os.Clearenv()

	config := &Config{}
	data := []byte(testConfigFileContent)
	err := yaml.Unmarshal(data, config)
	if err != nil {
		t.Fatalf("Failed to unmarshal config: %v", err)
	}

	err = config.overrideWithEnvVars()
	if err != nil {
		t.Fatalf("Failed to override config with environment variables: %v", err)
	}

	if config.Environment != "testenv" || config.Host != "127.0.0.1" || config.Port != 8081 || *config.Redis != "redis://user:password@localhost:6379/1" {
		t.Errorf("Environment variables not overridden correctly: %#v", config)
	}
}
