package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server `yaml:"server"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func MustConfig() *Config {
	configPath := "./config/config.yaml"

	if _, err := os.Stat(configPath); os.IsExist(err) {
		log.Fatalf("config %v does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cant read config %v", err.Error())
	}

	return &cfg
}
