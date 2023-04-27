package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

import "fmt"

type ServerConfig struct {
	KeysDir string `yaml:"keys_dir" env:"KEYS_DIR"`
	GS      `yaml:"google_sheets"`
	TG      `yaml:"tg"`
	GRPC    `yaml:"grpc"`
}

type GS struct {
	ServiceKey string `yaml:"service_key" env:"GS_SERVICE_KEY"`
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

func NewServerConfig(filePath string, useEnv bool) (*ServerConfig, error) {
	cfg := &ServerConfig{}
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
