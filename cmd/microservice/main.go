package main

import (
	"log/slog"
	"micriservice/internal/config"
	"micriservice/internal/logger"
	"micriservice/internal/storage/mysql"

	"github.com/labstack/echo"
)

func main() {

	// Init config.
	cfg := config.MustLoad()

	// Init logger
	log := logger.SetupLogger(cfg.Env)

	log.Info("Starting url-shortener.", slog.String("env", cfg.Env))

	// Init storage

	storage, err := mysql.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to init storage.", slog.String("err", err.Error()))
		return
	}

	_ = storage

	// Init router

	e := echo.New()

	// Run server

	e.Start(cfg.Address)
	if err != nil {
		log.Error("Failed to run server.", slog.String("err", err.Error()))
		return
	}
}
