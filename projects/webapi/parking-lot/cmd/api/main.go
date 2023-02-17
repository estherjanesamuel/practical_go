package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"sync"

	"parking-lot/internal/database"
	"parking-lot/internal/env"
	"parking-lot/internal/leveledlog"
	"parking-lot/internal/version"
)

func main() {
	logger := leveledlog.NewLogger(os.Stdout, leveledlog.LevelAll, true)

	err := run(logger)
	if err != nil {
		trace := debug.Stack()
		logger.Fatal(err, trace)
	}
}

type config struct {
	baseURL  string
	httpPort int
	db       struct {
		dsn string
	}
}

type application struct {
	config config
	db     *database.DB
	logger *leveledlog.Logger
	wg     sync.WaitGroup
}

func run(logger *leveledlog.Logger) error {
	var cfg config

	cfg.baseURL = env.GetString("BASE_URL", "http://localhost:4444")
	cfg.httpPort = env.GetInt("HTTP_PORT", 4444)
	cfg.db.dsn = env.GetString("DB_DSN", "user:pass@localhost:5432/db")

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	db, err := database.New(cfg.db.dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	app := &application{
		config: cfg,
		db:     db,
		logger: logger,
	}

	return app.serveHTTP()
}
