package main

import (
	"log/slog"
	"micriservice/internal/config"
	"micriservice/internal/http-server/handlers/task/save"
	"micriservice/internal/logger"
	http_server "micriservice/internal/pkg/http-server/echo"

	"micriservice/internal/storage"
	"micriservice/internal/storage/mysql_db/table/task"

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

	// Create http server

	e := http_server.New(cfg)

	// Init router

	e.Post("/task/create", save.New(log, task.New(storage)))

	// Middleware

	e.UseMiddleware(
		middleware.Recover(),
		middleware.RequestID(),
		middleware.Logger(),
	)

	// Run server

	e.Run()
	if err != nil {
		log.Error("Failed to run server.", slog.String("err", err.Error()))
		return
	}
}
