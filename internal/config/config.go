package config

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`

	NodeUrl string `envconfig:"NODE_URL" required:"true"`

	Host           string        `envconfig:"HOST" default:"0.0.0.0"`
	Port           string        `envconfig:"PORT" default:"8000"`
	HttpTimeout    time.Duration `envconfig:"HTTP_TIMEOUT" default:"30s"`
	WorkerInterval time.Duration `envconfig:"WORKER_INTERVAL" default:"10s"`
}

func LoadConfig(prefix string) (*Config, error) {
	var cfg Config

	err := envconfig.Process(prefix, &cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, err
}
