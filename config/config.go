package config

import (
	"fmt"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP HTTP `yaml:"http"`
		Log  Log  `yaml:"logger"`
		PG   PG   `yaml:"postgres"`
	}

	HTTP struct {
		Port string `yaml:"port"`
	}

	Log struct {
		Level       string `yaml:"level"`
		Destination string `yaml:"destination" env:"LOG_DESTINATION"`		
	}

	PG struct {
		PoolMax int    `yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `yaml:"pg_url" env:"PG_URL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("can't read yml config: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("can't read env config: %w", err)
	}

	if err := validateDestination(cfg.Log.Destination); err != nil {
		return nil, err
	}

	return cfg, nil
}

func validateDestination(destination string) error {
	destination = strings.ToLower(destination)
	if destination != "file" && destination != "console" {
		return fmt.Errorf("invalid log destination: %s. Use 'file' or 'console'", destination)
	}
	return nil
}
