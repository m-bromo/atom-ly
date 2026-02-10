package config

import (
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type Config struct {
	Env Environment
}

func NewConfig() (*Config, error) {
	var config Config

	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	_, err := env.UnmarshalFromEnviron(&config.Env)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
