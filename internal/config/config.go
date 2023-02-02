package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

import "fmt"

type Config struct {
	App  `yaml:"app"`
	GS   `yaml:"google_sheets"`
	TG   `yaml:"tg"`
	GRPC `yaml:"grpc"`
}

type App struct {
	Name string `yaml:"name"`
}

type GS struct {
	ServiceKeyPath string `yaml:"service_key_path"`
	SpreadsheetID  string `yaml:"spreadsheet_id"`
}

type TG struct {
	Token string `yaml:"token"`
	Chat  int64  `yaml:"chat"`
}

type GRPC struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

func NewConfig(filePath string) (*Config, error) {
	cfg := &Config{}
	fmt.Println(filePath)
	err := cleanenv.ReadConfig(filePath, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
