package main

import (
	"micriservice/internal/config"
	"micriservice/internal/logger"
)

func main() {

	// Init config.
	cfg := config.MustLoad()

	// Init logger
	log := logger.SetupLogger(cfg.Env)

	log.Info("Starting url-shortener.")
	log.Debug("Debug messages are enabled. Env: " + cfg.Env)

	// TODO: init storage

	// TODO: init router

	// TODO: run server
}
