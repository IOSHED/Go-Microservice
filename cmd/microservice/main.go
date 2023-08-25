package main

import (
	"log/slog"
	"micriservice/internal/config"
	"micriservice/internal/logger"
	"micriservice/internal/storage/mysql"
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

	id, err := storage.SaveTask("todo2", "this do it v2")
	if err != nil {
		log.Error("Failed to save task in storage.", slog.String("err", err.Error()))
	} else {
		log.Info("Save Task id", slog.Int64("id", id))
	}

	// TODO: init router

	// TODO: run server
}
