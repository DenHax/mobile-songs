package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server  ServerConfig `yaml:"server"`
	Storage Postgres     `yaml:"postgres"`
	Logger  Logger
}

type ServerConfig struct {
	Address      string        `yaml:"address" env-default:"localhost:8080"`
	SSLMode      string        `yaml:"ssl_mode" env-default:"disable"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type Postgres struct {
	URL           string `env:"POSTGRES_URL"`
	SSLMode       string `env:"SSL_MODE"`
	MigrationPath string `yaml:"migration_path"`
}

type Logger struct {
	Env   string `yaml:"env" env-default:"local"`
	Level string `yaml:"level"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	return &cfg
}
