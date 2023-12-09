package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v7"
)

type Config struct {
	FilesStorageDirectory string        `env:"FILES_STORAGE_DIRECTORY"`
	NmapCacheTTl          time.Duration `env:"NMAP_CACHE_TTL" envDefault:"5m"`
}

func New() (*Config, error) {
	cfg := &Config{}

	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("env.Parse: %w", err)
	}

	return cfg, nil
}
