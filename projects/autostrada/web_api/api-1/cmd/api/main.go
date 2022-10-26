package main

import (
	"api-1/internal/server"
	"api-1/internal/version"
	"flag"
	"fmt"
	"log"
	"os"
)

func main()  {
	var cfg config

	flag.StringVar(&cfg.addr, "addr", "localhost:4444", "server address to listen on")
	flag.StringVar(&cfg.baseURL, "base-ulr", "http:localhost:4444", "base URL for the application")
	flag.StringVar(&cfg.env, "env", "development", "operating environment: development, testing, staging or production")
	flag.BoolVar(&cfg.version, "version", false, "display version and exit")

	flag.Parse()

	if cfg.version {
		fmt.Printf("version: %s\n", version.Get())
	}

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)
	app := &application{
		config: cfg,
		logger: logger,
	}

	logger.Printf("starting server on %s (version %s)", cfg.addr, version.Get())
	err := server.Run(cfg.addr, app.routes())
	if err != nil {
		logger.Fatal(err)
	}
}

type config struct {
	addr string
	baseURL string
	env string
	version bool
}

type application struct {
	config config
	logger *log.Logger
}

