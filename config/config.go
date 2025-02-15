package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
	}

	HTTP struct {
		Port string `yaml:"http"`
	}

	Log struct {
		Level string `yaml:"level"`
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

	return cfg, nil
}
