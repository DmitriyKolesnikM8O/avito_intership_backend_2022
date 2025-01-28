package config

import (
	"fmt"
	"path"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `yaml:"app"`
		Http `yaml:"http"`
		Log  `yaml:"log"`
		Pg   `yaml:"postgres"`
		Jwt  `yaml:"jwt"`
	}

	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	Http struct {
		Port string `yaml:"port"`
	}

	Log struct {
		Level string `yaml:"level"`
	}

	Pg struct {
		URL         string `yaml: "PG_URL"`
		MaxPoolSize int    `yaml:"max_pool_size"`
	}

	Jwt struct {
		TokenTTL time.Duration `yaml:"token_ttl"`
	}
)

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path.Join("./", configPath), cfg)
	if err != nil {
		return &Config{}, fmt.Errorf("Error reading config file: %s", err)
	}

	return cfg, nil
}
