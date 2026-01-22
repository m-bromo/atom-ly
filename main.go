package main

import (
	"log/slog"

	"github.com/m-bromo/atom-ly/config"
	"github.com/m-bromo/atom-ly/logger"
)

func main() {
	config.SetupEnvironment()
	logger.SetupLog(config.Env.Environment)

	slog.Info("Starting application")
}
