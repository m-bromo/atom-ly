package logger

import (
	"log/slog"
	"os"

	"github.com/m-bromo/atom-ly/config"
)

type Logger struct {
	Log *slog.Logger
}

func NewLogger(cfg *config.Config) *Logger {
	var logger Logger

	switch cfg.Env.Environment {
	case "staging":
		logger.Log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelWarn,
			AddSource: true,
		}))

	case "production":
		logger.Log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: true,
		}))

	default:
		logger.Log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: false,
		}))
	}

	return &logger
}
