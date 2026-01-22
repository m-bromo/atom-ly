package config

import (
	"log"
	"log/slog"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

var Env Environment

func SetupEnvironment() {
	if err := godotenv.Load(".env"); err != nil {
		slog.Info("Loaded .env file")
		log.Fatal(err)
	}

	_, err := env.UnmarshalFromEnviron(&Env)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Environment loaded", "env", Env)
}
