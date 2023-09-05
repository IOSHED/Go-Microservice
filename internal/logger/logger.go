package logger

import (
	"log"
	"log/slog"
	"micriservice/internal/config"
	"os"
)

func SetupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case config.EnvLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)

	case config.EnvDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)

	case config.EnvProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}),
		)

	default:
		log.Fatalf("the ENV value has an invalid value. It can have the values \"dev\" \"local\" \"prod\"")
	}

	return logger
}
