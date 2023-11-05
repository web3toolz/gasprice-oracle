package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"time"
)

var (
	exampleConfigPath = "config.example.yaml"
)

type ServerConfig struct {
	Host    string        `envconfig:"SERVER_HOST" yaml:"host" default:"0.0.0.0"`
	Port    string        `envconfig:"SERVER_PORT" yaml:"port" default:"8000"`
	Timeout time.Duration `envconfig:"SERVER_TIMEOUT" yaml:"timeout" default:"30s"`
	Path    string        `envconfig:"SERVER_PATH" yaml:"path" default:"/"`
}

type WorkerConfig struct {
	Interval time.Duration `envconfig:"WORKER_INTERVAL"  yaml:"interval" default:"10s"`
}

type NetworkConfig struct {
	NetworkName string `yaml:"name"`
	Url         string `yaml:"url"`
}

type Config struct {
	LogLevel string `envconfig:"LOG_LEVEL" yaml:"logLevel" default:"0"`
	Server   ServerConfig
	Worker   WorkerConfig
	Networks []NetworkConfig
}

func openExampleConfig() (*os.File, error) {
	fmt.Printf("Loading default config (path %s)\n", exampleConfigPath)

	f, err := os.Open(filepath.Clean(exampleConfigPath))

	if err != nil {
		return nil, fmt.Errorf("failed to open example config file: %w", err)
	}

	return f, nil

}

func LoadConfigFromEnv(prefix string) (*Config, error) {
	var cfg *Config

	err := envconfig.Process(prefix, &cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func LoadConfigFromFile(path string) (*Config, error) {
	var cfg *Config

	f, err := os.Open(filepath.Clean(path))

	if err != nil && os.IsNotExist(err) {
		f, err = openExampleConfig()
	}

	if err != nil {
		return nil, err
	}

	err = yaml.NewDecoder(f).Decode(&cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
