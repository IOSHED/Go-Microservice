package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	StoragPath string `yaml:"storag_path" env-required:"./storage"`
	HttpServer `yaml:"httpserver"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env-required:"./storage"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	// getting the --config flag when starting the program
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	if *configPath == "" {
		log.Fatalf("Flag --config undefined during program startup")
	}

	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", *configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(*configPath, &cfg); err != nil {
		log.Fatalf("Cannot read file config: %s", *configPath)
	}

	return &cfg
}
