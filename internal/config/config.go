package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

import "fmt"

type Config struct {
	GS   `yaml:"google_sheets"`
	TG   `yaml:"tg"`
	GRPC `yaml:"grpc"`
}

type GS struct {
	ServiceKeyPath string `yaml:"service_key_path"`
	SpreadsheetID  string `yaml:"spreadsheet_id"`
}

type TG struct {
	IsEnabled bool   `yaml:"is_enabled" env:"TG_ENABLED"`
	Token     string `yaml:"token" env:"TG_TOKEN"`
	Chat      int64  `yaml:"chat" env:"TG_CHAT"`
}

type GRPC struct {
	IP   string `yaml:"ip" env:"GRPC_IP"`
	Port int    `yaml:"port" env:"GRPC_PORT"`
}

func NewConfig(filePath string, useEnv bool) (*Config, error) {
	cfg := &Config{}
	if useEnv {
		err := cleanenv.ReadEnv(cfg)
		if err != nil {
			return nil, err
		}
	} else {
		err := cleanenv.ReadConfig(filePath, cfg)
		if err != nil {
			return nil, fmt.Errorf("config error: %w", err)
		}
	}

	return cfg, nil
}
