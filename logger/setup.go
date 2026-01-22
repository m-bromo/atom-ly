package logger

import (
	"log/slog"
	"os"
)

func SetupLog(env string) {
	var logger *slog.Logger

	switch env {
	case "staging":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelWarn,
			AddSource: true,
		}))

	case "production":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: true,
		}))

	default:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: false,
		}))
	}

	slog.SetDefault(logger)
}
