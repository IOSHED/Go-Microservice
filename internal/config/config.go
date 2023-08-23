package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config holds the application configuration.
type Config struct {
	Env        string `yaml:"env" env-default:"local" env-description:"the way the microservice is launched: on local, in dev, in prod"`
	StoragPath string `yaml:"storag_path" env-required:"./storage"`
	HttpServer `yaml:"httpserver"`
}

// HttpServer holds HTTP server configuration.
type HttpServer struct {
	Address     string        `yaml:"address" env-required:"./storage"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s" env-description:"the maximum waiting time to complete a single operation (for example, to complete an HTTP request)"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s" env-description:"how long will the server wait for activity (for example, a new request) from the client"`
}

// MustLoad loads the configuration.
func MustLoad() *Config {

	// Getting the --config flag when starting the program.
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	if *configPath == "" {
		log.Fatalf("Flag --config undefined during program startup")
	}

	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", *configPath)
	}

	// Read config.
	var cfg Config
	if err := cleanenv.ReadConfig(*configPath, &cfg); err != nil {
		log.Fatalf("Cannot read file config: %s", *configPath)
	}

	return &cfg
}
