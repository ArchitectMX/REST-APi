package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	pathToDataBase string `env:"DATABASE_PATH"`
	id             int    `env:"ID"`
}

func ActivConfig() (*Config, error) {
	cfg := Config{}

	if err := cleanenv.ReadConfig("berendeev/restApi/config/config.env", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
