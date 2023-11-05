package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"time"
)

type ServerConfig struct {
	Host    string        `envconfig:"SERVER_HOST" yaml:"host" default:"0.0.0.0"`
	Port    string        `envconfig:"SERVER_PORT" yaml:"port" default:"8000"`
	Timeout time.Duration `envconfig:"SERVER_TIMEOUT" yaml:"timeout" default:"30s"`
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

	if err != nil {
		log.Fatalf("failed to unmarshal yaml file: %v", err)
	}

	err = yaml.NewDecoder(f).Decode(&cfg)

	if err != nil {
		log.Fatalf("failed to unmarshal yaml file: %v", err)
	}

	return cfg, nil
}
