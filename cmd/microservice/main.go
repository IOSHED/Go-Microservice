package main

import (
	"log/slog"
	"micriservice/internal/config"
	"micriservice/internal/http-server/handlers/url/save"
	"micriservice/internal/logger"
	"micriservice/internal/storage"
	"micriservice/internal/storage/mysql_db/table/task"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Init config.
	cfg := config.MustLoad()

	// Init logger
	log := logger.SetupLogger(cfg.Env)

	log.Info("Starting url-shortener.", slog.String("env", cfg.Env))

	// Init storage

	storage, err := storage.New("mysql", cfg.StoragePath)
	if err != nil {
		log.Error("Failed to init storage.", slog.String("err", err.Error()))
		return
	}

	// Init router

	e := echo.New()

	e.POST("/task/create", save.New(storage.Db, log, &task.Task{}))

	// Middleware

	e.Use(
		middleware.Recover(),
		middleware.RequestID(),
		middleware.Logger(),
	)

	// Run server

	e.Start(cfg.Address)
	if err != nil {
		log.Error("Failed to run server.", slog.String("err", err.Error()))
		return
	}
}
